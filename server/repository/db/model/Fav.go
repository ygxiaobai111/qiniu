package model

import "gorm.io/gorm"

// Fav 点赞模型
type Fav struct {
	gorm.Model
	UserId  int64 `gorm:"column:user_id;"`
	VideoId int64 `gorm:"column:video_id;"`

	// belongs to
	Video Video
}
