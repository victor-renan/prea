package middlewares

import "github.com/gin-gonic/gin"

const (
	MtSucc = iota
	MtWarn
	MtDang
)

type TextRes struct {
	Msg string `json:"msg"`
	MsgT  int `json:"msgt"`
}

func Response(ctx *gin.Context, code int, data any) {
	ctx.IndentedJSON(code, data)
}

func DataRes(ctx *gin.Context, data any, code int) {
	Response(ctx, code, data)
}

func SuccRes(ctx *gin.Context, msg string, code int) {
	Response(ctx, code, TextRes{msg, MtSucc})
}

func WarnRes(ctx *gin.Context, msg string, code int) {
	Response(ctx, code, TextRes{msg, MtWarn})
}

func DangRes(ctx *gin.Context, msg string, code int) {
	Response(ctx, code, TextRes{msg, MtDang})
}