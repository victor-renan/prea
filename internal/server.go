package app

import (
	"log"
	"prea/internal/controllers"
	"prea/internal/domain/models"
	"prea/internal/repositories/generic"
	"prea/internal/services"

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

	controllers.BookController{
		Service: services.BookService{
			Repo: generic.DBGeneric[models.Book]{},
		},
	}.ForEngine(router)

	err := router.Run(srv.Port)

	if err != nil {
		log.Default().Fatal(err)
	}
}
