package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"prea/internal/generics/repositories"
)

type UserMod struct {
	Controller UserController
	Service    UserService
}

func Prepare(connStr string) UserMod {
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, connStr)

	if err != nil {
		log.Fatal(err)
	}

	ub := UserMod{}

	ub.Service = UserService{
		Repo: repositories.DBGeneric[User]{
			Conn: conn,
			Ctx:  ctx,
		},
	}

	ub.Controller = UserController{Service: ub.Service}

	return ub
}

func (ub UserMod) Mount(router *gin.Engine) {
	ub.Controller.ForEngine(router)
}
