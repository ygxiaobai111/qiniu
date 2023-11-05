package model

import "gorm.io/gorm"

// Fav 点赞模型
type Fav struct {
	gorm.Model
	UserId  uint `gorm:"column:user_id;"`
	VideoId uint `gorm:"column:video_id;"`

	// belongs to
	Video Video
}
