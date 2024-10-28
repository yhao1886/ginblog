package model

type Article struct {
	Model

	Title       string `gorm:"type:varchar(100);not null" json:"title"`
	Desc        string `json:"desc"`
	Content     string `json:"content"`
	Img         string `json:"img"`
	Type        int    `gorm:"type:tinyint;comment:类型(1-原创 2-转载 3-翻译)" json:"type"` // 1-原创 2-转载 3-翻译
	Status      int    `gorm:"type:tinyint;comment:状态(1-公开 2-私密)" json:"status"`    // 1-公开 2-私密
	IsTop       bool   `json:"is_top"`
	IsDelete    bool   `json:"is_delete"`
	OriginalUrl string `json:"original_url"`

	CategoryId int `json:"category_id"`
	UserId     int `json:"-"` // user_auth_id

	Tags     []*Tag    `gorm:"many2many:article_tag;joinForeignKey:article_id" json:"tags"`
	Category *Category `gorm:"foreignkey:CategoryId" json:"category"`
	User     *UserAuth `gorm:"foreignkey:UserId" json:"user"`
}
