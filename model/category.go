package model

import "gorm.io/gorm"

// 分类
type Category struct {
	Model
	Name     string     `gorm:"unique;type:varchar(20);not null" json:"name"`
	Articles []*Article `gorm:"foreignKey:CategoryId"`
}

type CategoryVO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Total int    `json:"article_count"`
}

func GetCategoryList(db *gorm.DB) (data []CategoryVO, err error) {
	var categorylist []Category
	result := db.Model(&Category{}).Find(&categorylist)
	if result.Error != nil {
		return data, result.Error
	}
	for _, category := range categorylist {
		var total int64
		result = db.Model(&Article{}).Where("category_id = ?", category.ID).Count(&total)
		if result.Error != nil {
			return data, result.Error
		}
		var categoryvo = CategoryVO{
			ID:    category.ID,
			Name:  category.Name,
			Total: int(total),
		}
		data = append(data, categoryvo)
	}
	return data,nil
}
