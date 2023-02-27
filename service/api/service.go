package api

import (
	"gateway/models"
	"gateway/public"
	"gateway/schemas"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strconv"
)

type ServiceApi struct{}

func ServiceRegister(group *gin.RouterGroup) {
	service := &ServiceApi{}
	group.GET("/list", service.ServiceList)
	group.GET("/:id", service.ServiceDetail)
	group.DELETE("/:id", service.ServiceDelete)
	group.POST("/service_add_http", service.ServiceAddHTTP)
	group.POST("/service_update_http", service.ServiceUpdateHTTP)
}

// ServiceList godoc
// @Summary 获取服务列表
// @Description 获取服务列表
// @Tags 服务管理接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param info query string false "关键词"
// @Param page_no query int true "页数"
// @Param page_size query int true "每页条数"
// @Success 200 {object} public.Response{data=schemas.ServiceListOutput} "success"
// @Router /service/list [get]
func (s *ServiceApi) ServiceList(ctx *gin.Context) {
	params := &schemas.ServiceListInput{}
	if err := params.BindValidParam(ctx); err != nil {
		public.ResponseError(ctx, public.ServiceListParamInvalid, err)
		return
	}

	service := &models.ServiceInfo{}
	serviceList, total, err := service.PageList(ctx, public.GormDB, params)
	if err != nil {
		public.ResponseError(ctx, public.ServiceListPageList, err)
		return
	}

	itemList := make([]schemas.ServiceItemOutput, 0)
	for _, item := range serviceList {
		serviceDetail, err := item.ServiceDetail(ctx, public.GormDB, &item)
		if err != nil {
			public.ResponseError(ctx, public.ServiceListServiceDetail, err)
			return
		}

		itemList = append(itemList, schemas.ServiceItemOutput{
			ID:             uint64(item.ID),
			ServiceName:    item.ServiceName,
			ServiceDesc:    item.ServiceDesc,
			ServiceAddr:    serviceDetail.HTTPRule.Rule,
			ServiceRewrite: serviceDetail.HTTPRule.UrlRewrite,
		})
	}

	out := &schemas.ServiceListOutput{
		Total: total,
		List:  itemList,
	}
	public.ResponseSuccess(ctx, out)
}

// ServiceDetail godoc
// @Summary 获取服务信息
// @Description 获取服务信息
// @Tags 服务管理接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} public.Response{data=models.ServiceDetail} "success"
// @Router /service/{id} [get]
func (s *ServiceApi) ServiceDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	serviceId, err := strconv.Atoi(id)
	if err != nil {
		public.ResponseError(ctx, public.ServiceDetailParamInvalid, errors.New("ID为必填字段"))
		return
	}

	serviceInfo := &models.ServiceInfo{ID: int64(serviceId)}
	serviceInfo, err = serviceInfo.Find(ctx, public.GormDB, serviceInfo)
	if err != nil {
		public.ResponseError(ctx, public.ServiceDetailGetInfoError, err)
		return
	}
	serviceDetail, err := serviceInfo.ServiceDetail(ctx, public.GormDB, serviceInfo)
	if err != nil {
		public.ResponseError(ctx, public.ServiceDetailDataError, err)
		return
	}

	public.ResponseSuccess(ctx, serviceDetail)
}

// ServiceDelete godoc
// @Summary 删除服务
// @Description 删除服务
// @Tags 服务管理接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200 {object} public.Response{data=string} "success"
// @Router /service/{id} [delete]
func (s *ServiceApi) ServiceDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	serviceId, err := strconv.Atoi(id)
	if err != nil {
		public.ResponseError(ctx, public.ServiceDeleteParamInvalid, errors.New("ID为必填字段"))
		return
	}

	serviceInfo := &models.ServiceInfo{ID: int64(serviceId)}
	serviceInfo, err = serviceInfo.Find(ctx, public.GormDB, serviceInfo)
	if err != nil {
		public.ResponseError(ctx, public.ServiceDeleteGetInfoError, err)
		return
	}
	serviceInfo.IsDelete = 1
	if err = serviceInfo.Save(ctx, public.GormDB); err != nil {
		public.ResponseError(ctx, public.ServiceDeleteSaveError, err)
		return
	}

	// 添加服务到记录表中
	models.ServiceManagerHandler.DeleteService(serviceInfo.ServiceName)
	public.ResponseSuccess(ctx, "删除成功")
}

// ServiceAddHTTP godoc
// @Summary 添加HTTP服务
// @Description 添加HTTP服务
// @Tags 服务管理接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body schemas.ServiceAddHTTPInput true "body"
// @Success 200 {object} public.Response{data=string} "success"
// @Router /service/service_add_http [post]
func (s *ServiceApi) ServiceAddHTTP(ctx *gin.Context) {
	params := &schemas.ServiceAddHTTPInput{}
	if err := params.BindValidParam(ctx); err != nil {
		public.ResponseError(ctx, public.ServiceAddHTTPParamInvalid, err)
		return
	}

	tx := public.GormDB.Begin()

	serviceInfo := &models.ServiceInfo{ServiceName: params.ServiceName}
	serviceInfo, err := serviceInfo.Find(ctx, tx, serviceInfo)
	if err == nil {
		tx.Rollback()
		public.ResponseError(ctx, public.ServiceAddHTTPGetInfoError, errors.New("服务已存在"))
		return
	}

	httUrl := &models.HttpRule{Rule: params.Rule, NeedWebsocket: 0, UrlRewrite: params.UrlRewrite}
	if _, err = httUrl.Find(ctx, tx, httUrl); err == nil {
		tx.Rollback()
		public.ResponseError(ctx, public.ServiceAddHTTPHttpUrlError, errors.New("服务已存在"))
		return
	}

	serviceModel := &models.ServiceInfo{
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
	}
	if err := serviceModel.Save(ctx, tx); err != nil {
		tx.Rollback()
		public.ResponseError(ctx, public.ServiceAddHTTPSaveError, err)
		return
	}

	httpRule := &models.HttpRule{
		ServiceID:     serviceModel.ID,
		Rule:          params.Rule,
		NeedWebsocket: params.NeedWebsocket,
		UrlRewrite:    params.UrlRewrite,
	}

	if err := httpRule.Save(ctx, tx); err != nil {
		tx.Rollback()
		public.ResponseError(ctx, public.ServiceAddHTTPRuleSaveError, err)
		return
	}

	tx.Commit()

	// 添加服务到记录表中
	models.ServiceManagerHandler.UpdateServiceMap(&models.ServiceDetail{
		Info:     serviceInfo,
		HTTPRule: httpRule,
	})
	public.ResponseSuccess(ctx, "添加服务成功")
}

// ServiceUpdateHTTP godoc
// @Summary 修改HTTP服务
// @Description 修改HTTP服务
// @Tags 服务管理接口
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body schemas.ServiceUpdateHTTPInput true "body"
// @Success 200 {object} public.Response{data=string} "success"
// @Router /service/service_update_http [post]
func (s *ServiceApi) ServiceUpdateHTTP(ctx *gin.Context) {
	params := &schemas.ServiceUpdateHTTPInput{}
	if err := params.BindValidParam(ctx); err != nil {
		public.ResponseError(ctx, public.ServiceUpdateHTTPParamInvalid, err)
		return
	}

	tx := public.GormDB.Begin()

	serviceInfo := &models.ServiceInfo{ServiceName: params.ServiceName}
	serviceInfo, err := serviceInfo.Find(ctx, tx, serviceInfo)
	if err != nil {
		tx.Rollback()
		public.ResponseError(ctx, public.ServiceUpdateHTTPGetInfoError, errors.New("服务不存在"))
		return
	}

	serviceDetail, err := serviceInfo.ServiceDetail(ctx, tx, serviceInfo)
	if err != nil {
		tx.Rollback()
		public.ResponseError(ctx, public.ServiceUpdateHTTPNotExist, errors.New("服务不存在"))
		return
	}

	info := serviceDetail.Info
	info.ServiceDesc = params.ServiceDesc
	if err := info.Save(ctx, tx); err != nil {
		tx.Rollback()
		public.ResponseError(ctx, public.ServiceUpdateHTTPSaveError, err)
		return
	}

	httpRule := serviceDetail.HTTPRule
	httpRule.NeedWebsocket = params.NeedWebsocket
	httpRule.UrlRewrite = params.UrlRewrite
	if err := httpRule.Save(ctx, tx); err != nil {
		tx.Rollback()
		public.ResponseError(ctx, public.ServiceUpdateHTTPRuleSaveError, err)
		return
	}
	tx.Commit()

	// 添加服务到记录表中
	models.ServiceManagerHandler.UpdateServiceMap(&models.ServiceDetail{
		Info:     serviceInfo,
		HTTPRule: httpRule,
	})
	public.ResponseSuccess(ctx, "更新服务成功")
}
