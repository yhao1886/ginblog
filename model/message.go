package model

type Message struct {
	Model
	Nickname  string `gorm:"type:varchar(50);comment:昵称" json:"nickname"`
	Avatar    string `gorm:"type:varchar(255);comment:头像地址" json:"avatar"`
	Content   string `gorm:"type:varchar(255);comment:留言内容" json:"content"`
	IpAddress string `gorm:"type:varchar(50);comment:IP 地址" json:"ip_address"`
	IpSource  string `gorm:"type:varchar(255);comment:IP 来源" json:"ip_source"`
	Speed     int    `gorm:"type:tinyint(1);comment:弹幕速度" json:"speed"`
	IsReview  bool   `json:"is_review"`
}
