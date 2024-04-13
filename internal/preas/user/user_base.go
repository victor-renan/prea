package user

import (
	"github.com/gin-gonic/gin"
	"prea/internal/generics/repositories"
)

func Mount(router *gin.Engine) {
	UserController{
		Service: UserService{
			Repo: repositories.DBGeneric[User]{},
		},
	}.ForEngine(router)
}
