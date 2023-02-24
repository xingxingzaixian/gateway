package models

import (
	"fmt"
	"gateway/public"
	"gateway/schemas"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http/httptest"
	"strings"
	"sync"
)

type ServiceDetail struct {
	Info     *ServiceInfo `json:"info" description:"基本信息"`
	HTTPRule *HttpRule    `json:"http_rule" description:"http_rule"`
}

var ServiceManagerHandler *ServiceManager

func init() {
	ServiceManagerHandler = NewServiceManager()
}

type ServiceManager struct {
	ServiceMap sync.Map
	init       sync.Once
	err        error
}

func NewServiceManager() *ServiceManager {
	return &ServiceManager{
		ServiceMap: sync.Map{},
		init:       sync.Once{},
	}
}

func (s *ServiceManager) HTTPAccessMode(c *gin.Context) (*ServiceDetail, error) {
	//1、前缀匹配 /abc ==> serviceSlice.rule
	//2、域名匹配 www.test.com ==> serviceSlice.rule
	//host c.Request.Host
	//path c.Request.URL.Path
	host := c.Request.Host
	host = host[0:strings.Index(host, ":")]
	path := c.Request.URL.Path
	var serviceDetail *ServiceDetail
	s.ServiceMap.Range(func(key, serviceItem any) bool {
		service := serviceItem.(*ServiceDetail)
		fmt.Println("222: ", path, service.HTTPRule.Rule)
		if strings.HasPrefix(path, service.HTTPRule.Rule) {
			serviceDetail = service
			return false
		}
		return true
	})
	if serviceDetail != nil {
		return serviceDetail, nil
	}
	return nil, errors.New("not matched service")
}

func (s *ServiceManager) LoadOnce() error {
	s.init.Do(func() {
		serviceInfo := &ServiceInfo{}
		c, _ := gin.CreateTestContext(httptest.NewRecorder())

		params := &schemas.ServiceListInput{PageNo: 1, PageSize: 99999}
		list, _, err := serviceInfo.PageList(c, public.GormDB, params)
		if err != nil {
			s.err = err
			return
		}
		for _, listItem := range list {
			tmpItem := listItem
			serviceDetail, err := tmpItem.ServiceDetail(c, public.GormDB, &tmpItem)
			if err != nil {
				s.err = err
				return
			}
			s.ServiceMap.Store(listItem.ServiceName, serviceDetail)
		}
	})

	return s.err
}

func (s *ServiceManager) UpdateServiceMap(serviceDetail *ServiceDetail) {
	s.ServiceMap.Store(serviceDetail.Info.ServiceName, serviceDetail)
}

func (s *ServiceManager) DeleteService(serviceName string) {
	s.ServiceMap.Delete(serviceName)
}
