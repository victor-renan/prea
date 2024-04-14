package app

import (
	"log"
	"prea/internal/preas/user"
	"github.com/gin-gonic/gin"
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

	err := router.Run(srv.Port)

	if err != nil {
		log.Default().Fatal(err)
	}
}
