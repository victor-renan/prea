package auth

import (
	"net/http"
	"prea/internal/helpers"
	"prea/internal/preas/user"
	"prea/internal/security"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	UserService user.IUserService
}

func (bc AuthController) RouteName() string {
	return "auth"
}

func (bc AuthController) ForEngine(router *gin.Engine) {
	gp := router.Group(bc.RouteName())
	{
		gp.POST("/login", func(ctx *gin.Context) {
			var body LoginBody
			if err := ctx.ShouldBind(&body); err != nil {
				helpers.Response(ctx, http.StatusInternalServerError, helpers.Message{
					Text: err.Error(),
					Code: helpers.WarningCode,
				})
				return
			}

			usr, _ := bc.UserService.GetByUsername(body.Username)
			match, err := security.ComparePassword(body.Password, usr.Password)

			if err != nil || !match {
				helpers.Response(ctx, http.StatusUnauthorized, helpers.Message{
					Text: "Credenciais incorretas",
					Code: helpers.WarningCode,
				})
				return
			}

			tm := time.Now()

			_, lastlogin, err := bc.UserService.AlterLastLogin(usr, tm)
			if err != nil {
				helpers.Response(ctx, http.StatusUnauthorized, helpers.Message{
					Text: err.Error(),
					Code: helpers.WarningCode,
				})
				return
			}

			tokenData := LoginPayload{
				Username:  usr.Username,
				Profile:   usr.Profile,
				LastLogin: lastlogin,
			}

			claims := jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(tm),
				ExpiresAt: jwt.NewNumericDate(tm.Add(time.Hour * 1)),
				Subject:   usr.Username,
			}

			token, err := security.Jwt[LoginPayload]{}.CreateToken(tokenData, claims)

			if err != nil || !match {
				helpers.Response(ctx, http.StatusUnauthorized, helpers.Message{
					Text: err.Error(),
					Code: helpers.WarningCode,
				})
				return
			}

			helpers.Response(ctx, http.StatusUnauthorized, helpers.DataMessage{
				Text: "Login realizado com sucesso!",
				Code: helpers.SuccessCode,
				Data: LoginResponse{
					Token:   token,
					Payload: tokenData,
				},
			})
		})
		gp.POST("/validate", func(ctx *gin.Context) {
			var body ValidateBody
			if err := ctx.ShouldBind(&body); err != nil {
				helpers.Response(ctx, http.StatusBadRequest, helpers.Message{
					Text: err.Error(),
					Code: helpers.WarningCode,
				})
				return
			}

			decode, err := security.Jwt[LoginPayload]{}.DecodeToken(body.Token)
			if err != nil {
				helpers.Response(ctx, http.StatusOK, helpers.Message{
					Text: "Token inválido!",
					Code: helpers.SuccessCode,
				})
				return
			}

			helpers.Response(ctx, http.StatusOK, helpers.DataMessage{
				Text: "Token válido!",
				Code: helpers.SuccessCode,
				Data: decode,
			})
		})
	}
}