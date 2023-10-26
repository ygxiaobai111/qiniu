package model

import "gorm.io/gorm"

// Category 区
type Category struct {
	gorm.Model
	CategoryName string
}
