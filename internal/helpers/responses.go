package helpers

import "github.com/gin-gonic/gin"

const (
	SuccessCode = iota + 1
	WarningCode
	DangerCode
)

type Message struct {
	Text string `json:"text"`
	Code int    `json:"code"`
}

type DataMessage struct {
	Text string `json:"text"`
	Code int    `json:"code"`
	Data any    `json:"data"`
}

func Response(ctx *gin.Context, code int, data any) {
	ctx.IndentedJSON(code, data)
}