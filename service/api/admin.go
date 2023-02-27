package api

import (
	"gateway/models"
	"gateway/public"
	"gateway/schemas"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strconv"
)

type AdminApi struct{}

func AdminRegister(group *gin.RouterGroup) {
	admin := AdminApi{}
	group.GET("/info", admin.AdminInfo)
	group.POST("/change_pwd", admin.ChangePwd)
	group.DELETE("/:id", admin.DeleteUser)
}

// AdminInfo godoc
// @Summary 管理员信息
// @Description 管理员信息
// @Security ApiKeyAuth
// @Tags 管理员接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} public.Response{data=schemas.AdminInfoOutput} "success"
// @Router /admin/info [get]
func (a *AdminApi) AdminInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	out := &schemas.AdminInfoOutput{
		ID:        uint64(user.(*models.Admin).ID),
		Name:      user.(*models.Admin).UserName,
		LoginTime: user.(*models.Admin).UpdatedAt,
	}
	public.ResponseSuccess(ctx, out)
}

// ChangePwd godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags 管理员接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body schemas.ChangPwdInput true "body"
// @Success 200 {object} public.Response{data=string} "success"
// @Router /admin/change_pwd [post]
func (a *AdminApi) ChangePwd(ctx *gin.Context) {
	params := &schemas.ChangPwdInput{}
	if err := params.BindValidParam(ctx); err != nil {
		public.ResponseError(ctx, public.AdminChangePwdParamInvalid, err)
		return
	}

	// 1. 判断旧的密码是否正确
	auth, _ := ctx.Get("user")
	user := auth.(*models.Admin)
	if !user.CheckPassword(params.Password) {
		public.ResponseError(ctx, public.AdminPasswordError, errors.New("密码错误"))
		return
	}

	// 2. 设置新密码
	user.Password = public.GenSaltPassword(params.NewPass)
	if err := user.Save(ctx, public.GormDB); err != nil {
		public.ResponseError(ctx, public.AdminChangePwdError, err)
		return
	}

	public.ResponseSuccess(ctx, "密码修改成功")
}

// DeleteUser godoc
// @Summary 删除用户
// @Description 删除用户
// @Tags 管理员接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id  path int true "Account ID"
// @Success 200 {object} public.Response{data=string} "success"
// @Router /admin/{id} [delete]
func (a *AdminApi) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		public.ResponseError(ctx, public.AdminDeleteParamInvalid, err)
		return
	}

	admin := &models.Admin{}
	// 1. 判断用户是否已存在
	admin, err = admin.Find(ctx, public.GormDB, &models.Admin{ID: userId})
	if err != nil {
		public.ResponseError(ctx, public.AdminUserNotExist, errors.New("用户不存在"))
		return
	}

	err = admin.Delete(ctx, public.GormDB)
	if err != nil {
		public.ResponseError(ctx, public.AdminDeleteUserError, errors.New("删除用户失败"))
		return
	}
	public.ResponseSuccess(ctx, "用户已删除")
}
