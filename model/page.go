package model

import "gorm.io/gorm"

type Page struct {
	Model
	Name  string `gorm:"unique;type:varchar(20)" json:"name"`
	Label string `gorm:"unique;type:varchar(30)" json:"label"`
	Cover string `gorm:"type:varchar(255)" json:"cover"`
}

func GetPageList(db *gorm.DB) ([]Page, int64, error) {
	var pages = make([]Page, 0)
	var total int64

	result := db.Model(&Page{}).Count(&total).Find(&pages)
	return pages, total, result.Error
}