package schemas

import (
	"gateway/public"
	"github.com/gin-gonic/gin"
)

type AdminInfoOutput struct {
	ID       uint64 `json:"userId"`
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
}

type ChangPwdInput struct {
	Password string `json:"password" form:"password" comment:"旧密码" example:"123456" validate:"required"`
	NewPass  string `json:"new_pass" form:"new_pass" comment:"新密码" example:"123456" validate:"required"`
}

func (param *ChangPwdInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, param)
}
