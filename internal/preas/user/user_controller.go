package user

import (
	"net/http"
	"prea/internal/helpers"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service IUserService
}

func (uc UserController) RouteName() string {
	return "users"
}

func (uc UserController) ForEngine(router *gin.Engine) {
	gp := router.Group(uc.RouteName())
	{
		gp.GET("/", func(ctx *gin.Context) {
			objs, err := uc.Service.GetAll()
			if err != nil {
				helpers.Response(ctx, http.StatusInternalServerError, helpers.Message{
					Text: err.Error(),
					Code: helpers.DangerCode,
				})

				return
			}

			helpers.Response(ctx, http.StatusOK, objs)
		})

		gp.GET("/:id", func(ctx *gin.Context) {
			obj, err := uc.Service.GetById(ctx.Param("id"))
			if err != nil {
				helpers.Response(ctx, http.StatusInternalServerError, helpers.Message{
					Text: err.Error(),
					Code: helpers.DangerCode,
				})

				return
			}

			helpers.Response(ctx, http.StatusOK, obj)
		})

		gp.PUT("/", func(ctx *gin.Context) {
			var entity UserCreateDAO
			if err := ctx.ShouldBind(&entity); err != nil {
				helpers.Response(ctx, http.StatusInternalServerError, helpers.Message{
					Text: err.Error(),
					Code: helpers.WarningCode,
				})

				return
			}

			obj, err := uc.Service.Create(&entity)

			if err != nil {
				helpers.Response(ctx, http.StatusInternalServerError, helpers.Message{
					Text: err.Error(),
					Code: helpers.DangerCode,
				})

				return
			}

			helpers.Response(ctx, http.StatusOK, obj)
		})

		gp.PATCH("/:id", func(ctx *gin.Context) {
			var entity UserUpdateDAO
			if err := ctx.ShouldBind(&entity); err != nil {
				helpers.Response(ctx, http.StatusInternalServerError, helpers.Message{
					Text: err.Error(),
					Code: helpers.WarningCode,
				})

				return
			}

			obj, err := uc.Service.Update(ctx.Param("id"), &entity)
			if err != nil {
				helpers.Response(ctx, http.StatusNotFound, helpers.Message{
					Text: err.Error(),
					Code: helpers.WarningCode,
				})

				return
			}

			helpers.Response(ctx, http.StatusOK, obj)
		})

		gp.DELETE("/:id", func(ctx *gin.Context) {
			if err := uc.Service.Delete(ctx.Param("id")); err != nil {
				helpers.Response(ctx, http.StatusNotFound, helpers.Message{
					Text: err.Error(),
					Code: helpers.WarningCode,
				})
				return
			}

			helpers.Response(ctx, http.StatusOK, helpers.Message{
				Text: "Usuário deletado com sucesso!",
				Code: helpers.SuccessCode,
			})
		})
	}
}
