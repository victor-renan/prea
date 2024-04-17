package auth

import (
	"github.com/gin-gonic/gin"
	"prea/internal/preas/user"
)

type AuthModule struct {
	Controller  AuthController
	UserService user.UserService
}

func Prepare(mUser user.UserMod) AuthModule {
	am := AuthModule{}

	am.UserService = mUser.Service
	am.Controller = AuthController{UserService: am.UserService}

	return am
}

func (am AuthModule) Mount(router *gin.Engine) {
	am.Controller.ForEngine(router)
}
