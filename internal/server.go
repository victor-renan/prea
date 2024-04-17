package app

import (
	"prea/internal/common"
	"prea/internal/preas/auth"
	"prea/internal/preas/user"
	"github.com/gin-gonic/gin"
)

type IServer interface {
	Run()
}

type MainServer struct {
	Port string
}

func (srv MainServer) Run() error {
	router := gin.Default()
	dbStr := common.GetEnv("PGCONN")

	mUser := user.Prepare(dbStr)
	mUser.Mount(router)

	mAuth := auth.Prepare(mUser)
	mAuth.Mount(router)

	err := router.Run(srv.Port)

	return err
}
