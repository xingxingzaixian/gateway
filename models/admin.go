package models

import (
	"gateway/public"
	"gateway/schemas"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	ID        int       `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName  string    `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Password  string    `json:"password" gorm:"column:password" description:"密码"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete  int       `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (a *Admin) TableName() string {
	return "gateway_admin"
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	a.CreatedAt = time.Now()
	return
}

func (a *Admin) BeforeSave(tx *gorm.DB) (err error) {
	a.UpdatedAt = time.Now()
	return
}

func (a *Admin) BeforeUpdate(tx *gorm.DB) (err error) {
	a.UpdatedAt = time.Now()
	return
}

func (a *Admin) LoginCheck(c *gin.Context, tx *gorm.DB, param *schemas.AdminLoginInput) (*Admin, error) {
	adminInfo, err := a.Find(c, tx, &Admin{UserName: param.UserName, IsDelete: 0})
	if err != nil {
		return nil, err
	}
	saltPassword := public.GenSaltPassword(param.Password)
	if adminInfo.Password != saltPassword {
		return nil, errors.New("密码错误，请重新输入")
	}
	return adminInfo, nil
}

func (a *Admin) CheckPassword(password string) bool {
	saltPassword := public.GenSaltPassword(password)
	if saltPassword != a.Password {
		return false
	}
	return true
}

func (a *Admin) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(a).Error
}

func (a *Admin) Find(c *gin.Context, tx *gorm.DB, search *Admin) (*Admin, error) {
	out := &Admin{}
	result := tx.WithContext(c).Where(search).Find(out)
	if result.RowsAffected == 0 {
		return nil, errors.New("用户不存在")
	}

	return out, nil
}

func (a *Admin) Delete(c *gin.Context, tx *gorm.DB) error {
	result := tx.WithContext(c).Delete(a)
	if result.RowsAffected == 0 {
		return errors.New("删除数据失败")
	}
	return nil
}
