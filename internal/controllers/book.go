package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prea/internal/common"
	"prea/internal/domain/models"
	"prea/internal/middlewares"
	"prea/internal/services"
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

		middlewares.DataRes(ctx, objs, http.StatusOK)
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

		middlewares.DataRes(ctx, obj, http.StatusOK)
	})

	books.PUT("/", func(ctx *gin.Context) {
		var entity models.Book

		if err := ctx.Bind(&entity); err != nil {
			middlewares.WarnRes(ctx,
				err.Error(),
				http.StatusNotFound,
			)
			return
		}

		obj, err := bc.Service.Create(entity)
		if err != nil {
			middlewares.WarnRes(ctx,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		middlewares.DataRes(ctx, obj, http.StatusOK)
	})

	books.PATCH("/:id", func(ctx *gin.Context) {
		bookId := common.Stoi64(ctx.Param("id"))

		var entity models.Book

		if err := ctx.Bind(&entity); err != nil {
			middlewares.WarnRes(ctx,
				err.Error(),
				http.StatusNotFound,
			)
			return
		}

		obj, err := bc.Service.Update(bookId, entity)
		if err != nil {
			middlewares.WarnRes(ctx,
				err.Error(),
				http.StatusNotFound,
			)
			return
		}

		middlewares.DataRes(ctx, obj, http.StatusOK)
	})

	books.DELETE("/:id", func(ctx *gin.Context) {
		bookId := common.Stoi64(ctx.Param("id"))

		if err := bc.Service.Delete(bookId); err != nil {
			middlewares.WarnRes(ctx,
				err.Error(),
				http.StatusNotFound,
			)
			return
		}

		middlewares.SuccRes(ctx,
			"Item deletado com sucesso!",
			http.StatusOK,
		)
	})
}
