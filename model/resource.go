package model

import "gorm.io/gorm"

type Resource struct {
	Model
	Name      string `gorm:"unique;type:varchar(50)" json:"name"`
	ParentId  int    `json:"parent_id"`
	Url       string `gorm:"type:varchar(255)" json:"url"`
	Method    string `gorm:"type:varchar(10)" json:"request_method"`
	Anonymous bool   `json:"is_anonymous"`

	Roles []*Role `json:"roles" gorm:"many2many:role_resource"`
}

func GetResource(db *gorm.DB, uri, method string) (resource Resource, err error) {
	result := db.Where("url = ? AND method = ?", uri, method).First(&resource)
	return resource, result.Error
}
