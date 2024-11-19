package model

import "gorm.io/gorm"

type Tag struct {
	Model
	Name string `gorm:"unique;type:varchar(20);not null" json:"name"`

	Articles []*Article `gorm:"many2many:article_tag;" json:"articles,omitempty"`
}

func GetTaglist(db *gorm.DB) (data []Tag, err error) {
	result := db.Model(&Tag{}).Find(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}
