package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prea/internal/middlewares"
)

type UserController struct {
	Service IUserService
}

func (bc UserController) RouteName() string {
	return "users"
}

func (bc UserController) ForEngine(router *gin.Engine) {
	users := router.Group(bc.RouteName())
	{
		users.GET("/", func(ctx *gin.Context) {
			objs, err := bc.Service.GetAll()
			if err != nil {
				middlewares.WarnRes(ctx, err.Error(), http.StatusInternalServerError)
				return
			}

			middlewares.DataRes(ctx, objs, http.StatusOK)
		})

		users.GET("/:id", func(ctx *gin.Context) {
			obj, err := bc.Service.GetById(ctx.Param("id"))
			if err != nil {
				middlewares.WarnRes(ctx, err.Error(), http.StatusNotFound)
				return
			}

			middlewares.DataRes(ctx, obj, http.StatusOK)
		})

		users.PUT("/", func(ctx *gin.Context) {
			var entity UserCreateDAO
			if err := ctx.ShouldBind(&entity); err != nil {
				middlewares.WarnRes(ctx, err.Error(), http.StatusNotFound)
				return
			}

			obj, err := bc.Service.Create(entity)
			if err != nil {
				middlewares.WarnRes(ctx, err.Error(), http.StatusInternalServerError)
				return
			}

			middlewares.DataRes(ctx, obj, http.StatusOK)
		})

		users.PATCH("/:id", func(ctx *gin.Context) {
			var entity UserUpdateDAO
			if err := ctx.ShouldBind(&entity); err != nil {
				middlewares.WarnRes(ctx, err.Error(), http.StatusNotFound)
				return
			}

			obj, err := bc.Service.Update(ctx.Param("id"), entity)
			if err != nil {
				middlewares.WarnRes(ctx, err.Error(), http.StatusNotFound)
				return
			}

			middlewares.DataRes(ctx, obj, http.StatusOK)
		})

		users.DELETE("/:id", func(ctx *gin.Context) {
			if err := bc.Service.Delete(ctx.Param("id")); err != nil {
				middlewares.WarnRes(ctx, err.Error(), http.StatusNotFound)
				return
			}

			middlewares.SuccRes(ctx, "Item deletado com sucesso!", http.StatusOK)
		})
	}
}
