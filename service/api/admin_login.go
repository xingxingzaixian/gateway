package api

import (
	"gateway/models"
	"gateway/public"
	"gateway/schemas"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"
)

type AdminLoginApi struct{}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := AdminLoginApi{}
	group.POST("/login", adminLogin.AdminLogin)
	group.POST("/register", adminLogin.AdminRegister)
}

// AdminLogin godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @Accept application/json
// @Produce application/json
// @Param data body schemas.AdminLoginInput true "body"
// @Success 200 {object} public.Response{data=schemas.AdminLoginOutput} "success"
// @Router /api/admin_login/login [post]
func (a *AdminLoginApi) AdminLogin(ctx *gin.Context) {
	params := &schemas.AdminLoginInput{}
	if err := params.BindValidParam(ctx); err != nil {
		public.ResponseError(ctx, public.AdminLoginParamInvalid, err)
		return
	}

	// 1. 判断用户名、密码是否匹配
	admin := &models.Admin{}
	info, err := admin.LoginCheck(ctx, public.GormDB, params)
	if err != nil {
		public.ResponseError(ctx, public.AdminLoginUserOrPwdError, err)
		return
	}

	// 2. 根据用户信息创建jwt.token
	j := public.NewJWT()
	token, err := j.CreateToken(public.CustomClaims{
		ID:   uint64(info.ID),
		Name: info.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			Issuer:    "xingxingzaixian",
		},
	})
	if err != nil {
		public.ResponseError(ctx, public.AdminLoginCreateTokenError, err)
		return
	}

	out := &schemas.AdminLoginOutput{Token: token}
	public.ResponseSuccess(ctx, out)
}

// AdminRegister godoc
// @Summary 管理员注册
// @Description 管理员注册
// @Tags 管理员接口
// @Accept application/json
// @Produce application/json
// @Param data body schemas.AdminRegisterInput true "body"
// @Success 200 {object} public.Response{data=schemas.AdminRegisterInput} "success"
// @Router /api/admin_login/register [post]
func (a *AdminLoginApi) AdminRegister(ctx *gin.Context) {
	params := &schemas.AdminRegisterInput{}
	if err := params.BindValidParam(ctx); err != nil {
		public.ResponseError(ctx, public.AdminRegisterParamInvalid, err)
		return
	}

	// 0. 密码必须相同
	if params.Password != params.ConfirmPwd {
		public.ResponseError(ctx, public.AdminRegisterSamePassword, errors.New("两次密码必须一致"))
		return
	}

	admin := &models.Admin{}
	// 1. 判断用户是否已存在
	admin, err := admin.Find(ctx, public.GormDB, &models.Admin{UserName: params.UserName})
	if err == nil {
		public.ResponseError(ctx, public.AdminRegisterUserNotExist, errors.New("用户已存在"))
		return
	}

	newAdmin := models.Admin{}
	newAdmin.UserName = params.UserName
	newAdmin.NickName = params.NickName
	newAdmin.Password = public.GenSaltPassword(params.Password)
	if err = newAdmin.Save(ctx, public.GormDB); err != nil {
		public.ResponseError(ctx, public.AdminRegisterCreateUserError, err)
		return
	}

	public.ResponseSuccess(ctx, newAdmin)
}
