package model

import "gorm.io/gorm"

// Collection 收藏夹模型
type Collection struct {
	gorm.Model
	Name      string   // 收藏夹名称
	IsPrivate int      // 是否私有
	UserID    uint     // 所属用户的ID
	Videos    []*Video `gorm:"many2many:collection_videos;"` // 包含的视频列表
}
