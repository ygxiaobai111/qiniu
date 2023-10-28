package model

import "gorm.io/gorm"

// Comment
type Comment struct {
	gorm.Model
	//视频id
	VideoID int64 `json:"video_id"`
	//用户id
	UserId int64 `json:"user_id"`
	// 评论内容
	Content string `json:"content"`
	// 评论发布日期，格式 mm-dd
	CreateDate string `json:"create_date"`
}
