package app

import (
	"log"
	"prea/internal/generics/repositories"
_	"prea/internal/preas/book"
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

	user.UserController{
		Service: user.UserService{
			Repo: repositories.DBGeneric[user.User]{},
		},
	}.ForEngine(router)

	err := router.Run(srv.Port)

	if err != nil {
		log.Default().Fatal(err)
	}
}
