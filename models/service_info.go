package models

import (
	"gateway/schemas"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type ServiceInfo struct {
	ID          int64     `json:"id" gorm:"primary_key"`
	ServiceName string    `json:"service_name" gorm:"column:service_name" description:"服务名称"`
	ServiceDesc string    `json:"service_desc" gorm:"column:service_desc" description:"服务描述"`
	UpdateAt    time.Time `json:"create_at" gorm:"column:create_at" description:"更新时间"`
	CreateAt    time.Time `json:"update_at" gorm:"column:update_at" description:"添加时间"`
	IsDelete    int8      `json:"is_delete" gorm:"column:is_delete" description:"是否已删除；0：否；1：是"`
}

func (s *ServiceInfo) TableName() string {
	return "gateway_service_info"
}

func (s *ServiceInfo) BeforeCreate(tx *gorm.DB) (err error) {
	s.CreateAt = time.Now()
	return
}

func (s *ServiceInfo) BeforeSave(tx *gorm.DB) (err error) {
	s.UpdateAt = time.Now()
	return
}

func (s *ServiceInfo) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdateAt = time.Now()
	return
}

func (s *ServiceInfo) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(s).Error
}

func (s *ServiceInfo) Find(c *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceInfo, error) {
	out := &ServiceInfo{}
	result := tx.WithContext(c).Where(search).Find(out)
	if result.RowsAffected == 0 {
		return nil, errors.New("服务信息不存在")
	}

	return out, nil
}

func (s *ServiceInfo) PageList(c *gin.Context, tx *gorm.DB, param *schemas.ServiceListInput) ([]ServiceInfo, int64, error) {
	var total int64
	list := []ServiceInfo{}
	query := tx.WithContext(c).Table(s.TableName()).Where("is_delete=0")

	if param.Info != "" {
		info := "%" + param.Info + "%"
		query = query.Where("(service_name like ? or service_desc like ?)", info, info)
	}

	offset := (param.PageNo - 1) * param.PageSize
	result := query.Limit(param.PageSize).Offset(offset).Order("id desc").Find(&list)
	if result.RowsAffected == 0 {
		return nil, 0, errors.New("没有查询到数据")
	}

	query.Count(&total)
	return list, total, nil
}

func (s *ServiceInfo) ServiceDetail(c *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceDetail, error) {
	if search.ServiceName == "" {
		info, err := s.Find(c, tx, search)
		if err != nil {
			return nil, err
		}
		search = info
	}

	detail := &ServiceDetail{}
	detail.Info = search

	httpRule := &HttpRule{ServiceID: search.ID}
	httpRule, err := httpRule.Find(c, tx, httpRule)
	if err == nil {
		detail.HTTPRule = httpRule
	}

	return detail, nil
}
