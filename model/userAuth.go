package model

import (
	"time"

	"gorm.io/gorm"
)

type UserAuth struct {
	Model
	Username      string     `gorm:"unique;type:varchar(50)" json:"username"`
	Password      string     `gorm:"type:varchar(100)" json:"-"`
	LoginType     int        `gorm:"type:tinyint(1);comment:登录类型" json:"login_type"`
	IpAddress     string     `gorm:"type:varchar(20);comment:登录IP地址" json:"ip_address"`
	IpSource      string     `gorm:"type:varchar(50);comment:IP来源" json:"ip_source"`
	LastLoginTime *time.Time `json:"last_login_time"`
	IsDisable     bool       `json:"is_disable"`
	IsSuper       bool       `json:"is_super"` // 超级管理员只能后台设置

	UserInfoId int       `json:"user_info_id"`
	UserInfo   *UserInfo `json:"info"`
	Roles      []*Role   `json:"roles" gorm:"many2many:user_auth_role"`
}

func GetUserAuthInfoByName(db *gorm.DB, username string) (data UserAuth, err error) {
	var userAuth UserAuth
	result := db.Model(&UserAuth{}).Where("username = ?", username).First(&userAuth)
	if result.Error != nil {
		return data, result.Error
	}
	return userAuth, nil
}

func GetUserAuthInfoById(db *gorm.DB, id int) (data UserAuth, err error) {
	var userAuth UserAuth
	result := db.Model(&UserAuth{}).Where("id", id).First(&userAuth)
	if result.Error != nil {
		return data, result.Error
	}
	result = db.Model(&UserInfo{}).Where("id", userAuth.UserInfoId).First(&userAuth.UserInfo)
	if result.Error != nil {
		return data, result.Error
	}
	return userAuth, nil
}

func UpdateUserLoginInfo(db *gorm.DB, id int, ipAddress, ipSource string) error {
	now := time.Now()
	userAuth := UserAuth{
		IpAddress:     ipAddress,
		IpSource:      ipSource,
		LastLoginTime: &now,
	}
	result := db.Where("id = ?", id).Updates(userAuth)
	return result.Error
}
