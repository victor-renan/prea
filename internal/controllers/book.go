package controllers

import (
	"net/http"
	"prea/internal/common"
	"prea/internal/middlewares"
	"prea/internal/services"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	Service services.IBookService
}

func (bc BookController) RouteName() string {
	return "books"
}

func (bc BookController) ForEngine(router *gin.Engine) {
	books := router.Group(bc.RouteName())

	books.GET("/", func(ctx *gin.Context) {
		objs, err := bc.Service.GetAll()
		if err != nil {
			middlewares.WarnRes(ctx,
				"Erro ao encontrar itens!",
				http.StatusInternalServerError,
			)
			return
		}

		ctx.JSON(200, objs)
	})

	books.GET("/:id", func(ctx *gin.Context) {
		bookId := common.Stoi64(ctx.Param("id"))
		obj, err := bc.Service.GetById(bookId)
		if err != nil {
			middlewares.WarnRes(ctx,
				"Item não existe",
				http.StatusNotFound,
			)
			return
		}

		middlewares.DataRes(ctx, obj, 200)
	})
}
