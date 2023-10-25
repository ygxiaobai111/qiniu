package model

import "gorm.io/gorm"

// Category 收藏夹
type Category struct {
	gorm.Model
	CategoryName string
}
