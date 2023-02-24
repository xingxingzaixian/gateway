package public

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, &Response{
		Code: http.StatusOK,
		Msg:  "",
		Data: data,
	})
}

func ResponseError(ctx *gin.Context, code int, err error) {
	ctx.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  err.Error(),
		Data: "",
	})
}

const (
	ServiceListParamInvalid = 1001 + iota
	ServiceListPageList
	ServiceListServiceDetail
)

const (
	MiddleAccessMode = 601 + iota
	MiddleReverseProxy
)
