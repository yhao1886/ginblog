package model

import "gorm.io/gorm"

type Comment struct {
	Model
	UserId      int    `json:"user_id"`       // 评论者
	ReplyUserId int    `json:"reply_user_id"` // 被回复者
	TopicId     int    `json:"topic_id"`      // 评论的文章
	ParentId    int    `json:"parent_id"`     // 父评论 被回复的评论
	Content     string `gorm:"type:varchar(500);not null" json:"content"`
	Type        int    `gorm:"type:tinyint(1);not null;comment:评论类型(1.文章 2.友链 3.说说)" json:"type"` // 评论类型 1.文章 2.友链 3.说说
	IsReview    bool   `json:"is_review"`

	// Belongs To
	User      *UserAuth `gorm:"foreignKey:UserId" json:"user"`
	ReplyUser *UserAuth `gorm:"foreignKey:ReplyUserId" json:"reply_user"`
	Article   *Article  `gorm:"foreignKey:TopicId" json:"article"`
}

func AddComment(db *gorm.DB, userId, topicId, typ int, content string, isReview bool) (*Comment, error) {
	var comment = Comment{
		UserId:   userId,
		TopicId:  topicId,
		Content:  content,
		Type:     typ,
		IsReview: isReview,
	}
	result := db.Create(&comment)
	return &comment, result.Error
}

func ReplyComment(db *gorm.DB, userId, topicId, typ, replyId, parentId int, content string, isReview bool) (*Comment, error) {
	var parent Comment
	result := db.Where("id", parentId).First(&parent)
	if result.Error != nil {
		return nil, result.Error
	}
	var comment = Comment{
		UserId:      userId,
		TopicId:     topicId,
		Content:     content,
		Type:        typ,
		IsReview:    isReview,
		ReplyUserId: replyId,
		ParentId:    parentId,
	}
	result = db.Create(&comment)
	return &comment, result.Error
}
