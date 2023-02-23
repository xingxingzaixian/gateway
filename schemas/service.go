package schemas

import (
	"gateway/public"
	"github.com/gin-gonic/gin"
)

type ServiceListInput struct {
	Info     string `json:"info" form:"info" comment:"关键词" example:"" validate:""`                               //关键词
	PageNo   int    `json:"page_no" form:"page_no" comment:"page_size页数" example:"1" validate:"required"`        //页数
	PageSize int    `json:"page_size" form:"page_size" comment:"page_size每页条数" example:"20" validate:"required"` //每页条数
}

func (param *ServiceListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ServiceAddHTTPInput struct {
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名" validate:"required,valid_service_name"` //服务名
	ServiceDesc string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required,max=255,min=1"`     //服务描述

	Rule          string `json:"rule" form:"rule" comment:"接入路径：域名或者前缀" example:"类似/xxx/" validate:"required"`                    //域名或者前缀
	NeedWebsocket int    `json:"need_websocket" form:"need_websocket" comment:"是否支持websocket" example:"0" validate:"max=1,min=0"` //是否支持websocket
	UrlRewrite    string `json:"url_rewrite" form:"url_rewrite" comment:"url重写功能" example:"http://xx.xx.xx.xx:oo/"`               //url重写功能
}

func (param *ServiceAddHTTPInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ServiceItemOutput struct {
	ID             uint64 `json:"id" form:"id"`
	ServiceName    string `json:"service_name" form:"service_name"`
	ServiceDesc    string `json:"service_desc" form:"service_desc"`
	ServiceAddr    string `json:"service_addr" form:"service_addr"`
	ServiceRewrite string `json:"service_rewrite" form:"service_rewrite"`
}

type ServiceListOutput struct {
	Total int64               `json:"total" form:"total" comment:"总数"`
	List  []ServiceItemOutput `json:"list" form:"list" comment:"列表"`
}

type ServiceUpdateHTTPInput struct {
	ID          int64  `json:"id" form:"id" comment:"服务ID" example:"62" validate:"required,min=1"`                                                     //服务ID
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名" example:"test_http_service_indb" validate:"required,valid_service_name"` //服务名
	ServiceDesc string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"test_http_service_indb" validate:"required,max=255,min=1"`     //服务描述

	Rule          string `json:"rule" form:"rule" comment:"接入路径：域名或者前缀" example:"/test_http_service_indb" validate:"required"` //域名或者前缀 	//启用strip_uri
	NeedWebsocket int    `json:"need_websocket" form:"need_websocket" comment:"是否支持websocket" validate:"max=1,min=0"`          //是否支持websocket
	UrlRewrite    string `json:"url_rewrite" form:"url_rewrite" comment:"url重写功能"`                                             //header转换
}

func (param *ServiceUpdateHTTPInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
