package models

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type HttpRule struct {
	ID            int64  `json:"id" gorm:"primary_key"`
	ServiceID     int64  `json:"service_id" gorm:"column:service_id" description:"服务id"`
	Rule          string `json:"rule" gorm:"column:rule" description:"type=domain表示域名，type=url_prefix时表示url前缀"`
	NeedWebsocket int    `json:"need_websocket" gorm:"column:need_websocket" description:"启用websocket 1=启用"`
	UrlRewrite    string `json:"url_rewrite" gorm:"column:url_rewrite" description:"url重写功能，每行一个	"`
}

func (t *HttpRule) TableName() string {
	return "gateway_service_http_rule"
}

func (t *HttpRule) Find(c *gin.Context, tx *gorm.DB, search *HttpRule) (*HttpRule, error) {
	model := &HttpRule{}
	result := tx.WithContext(c).Where(search).Find(model)
	if result.RowsAffected == 0 {
		return nil, errors.New("没有获取到Http规则")
	}
	return model, nil
}

func (t *HttpRule) Save(c *gin.Context, tx *gorm.DB) error {
	if err := tx.WithContext(c).Save(t).Error; err != nil {
		return err
	}
	return nil
}
