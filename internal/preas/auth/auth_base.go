package auth

import (
	"github.com/gin-gonic/gin"
	"prea/internal/generics/repositories"
	"prea/internal/preas/user"
)

func Mount(router *gin.Engine) {
	AuthController{
		UserService: user.UserService{
			Repo: repositories.DBGeneric[user.User]{},
		},
	}.ForEngine(router)
}
