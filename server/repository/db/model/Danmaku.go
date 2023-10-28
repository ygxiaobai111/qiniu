package model

import "gorm.io/gorm"

// Danmaku 弹幕模型
type Danmaku struct {
	gorm.Model
	VideoID   uint   `json:"video_id"`  // 弹幕所属视频的ID
	UserID    uint   `json:"user_id"`   // 发布弹幕的用户ID
	Content   string `json:"content"`   // 弹幕内容
	Color     string `json:"color"`     // 弹幕颜色
	Timestamp uint   `json:"timestamp"` // 弹幕出现的时间戳
}
