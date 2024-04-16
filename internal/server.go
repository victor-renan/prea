package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"prea/internal/preas/auth"
	"prea/internal/preas/user"
)

type IServer interface {
	Run()
}

type MainServer struct {
	Port string
}

func (srv MainServer) Run() {
	router := gin.Default()

	user.Mount(router)
	auth.Mount(router)

	err := router.Run(srv.Port)

	if err != nil {
		log.Default().Fatal(err)
	}
}
