package app

import (
	"log"
	"prea/internal/preas/book"
	"prea/internal/generics/repositories"
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

	book.BookController{
		Service: book.BookService{
			Repo: repositories.DBGeneric[book.Book]{},
		},
	}.ForEngine(router)

	err := router.Run(srv.Port)

	if err != nil {
		log.Default().Fatal(err)
	}
}
