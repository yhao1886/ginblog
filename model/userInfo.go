package model

import "gorm.io/gorm"

type UserInfo struct {
	Model
	Email    string `json:"email" gorm:"type:varchar(30)"`
	Nickname string `json:"nickname" gorm:"unique;type:varchar(30);not null"`
	Avatar   string `json:"avatar" gorm:"type:varchar(1024);not null"`
	Intro    string `json:"intro" gorm:"type:varchar(255)"`
	Website  string `json:"website" gorm:"type:varchar(255)"`
}

type UserInfoVO struct {
	UserInfo
	ArticleLikeSet []string `json:"article_like_set"`
	CommentLikeSet []string `json:"comment_like_set"`
}

func GetUserInfoById(db *gorm.DB, id int) (data UserInfo, err error) {
	result := db.Model(&UserInfo{}).Where("id", id).First(&data)
	return data, result.Error
}
