package schemas

import (
	"gateway/public"
	"github.com/gin-gonic/gin"
)

type AdminLoginInput struct {
	UserName string `json:"username" form:"username" comment:"姓名" example:"admin" validate:"required"`
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`
}

func (param *AdminLoginInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, param)
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"Token" example:"token"`
}

type AdminRegisterInput struct {
	NickName   string `json:"nickname" form:"nickname" comment:"昵称" example:"管理员" validate:"required"`
	UserName   string `json:"username" form:"username" comment:"姓名" example:"admin" validate:"required"`
	Password   string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`
	ConfirmPwd string `json:"confirmPwd" form:"confirmPwd" comment:"确认密码" example:"123456" validate:"required"`
}

func (param *AdminRegisterInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, param)
}
