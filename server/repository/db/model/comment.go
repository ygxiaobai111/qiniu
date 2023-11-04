package model

import "gorm.io/gorm"

// Comment
type Comment struct {
	gorm.Model
	//视频id
	VideoID uint `json:"video_id"`
	//用户id
	UserId uint `json:"user_id"`
	// 评论内容
	Content string `json:"content"`
}
