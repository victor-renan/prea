package controllers

import (
	"log"
	"prea/internal/common"

	"github.com/gin-gonic/gin"
)

func GetLogger() *log.Logger {
	return common.MakeLogger("controllers")
}

type IController interface {
	RouteName() string
	AppendTo(router *gin.Engine)
}
