package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"prea/internal/generics/repositories"
)

type UserModule struct {
	Controller UserController
	Service    UserService
	Repo       repositories.DBGeneric[User]
}

func Prepare(connStr string) UserModule {
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, connStr)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ub := UserModule{}

	ub.Repo = repositories.DBGeneric[User]{
		Conn: conn,
		Ctx:  ctx,
	}

	ub.Service = UserService{Repo: ub.Repo}
	ub.Controller = UserController{Service: ub.Service}

	return ub
}

func (ub UserModule) Mount(router *gin.Engine) {
	ub.Controller.ForEngine(router)
}
