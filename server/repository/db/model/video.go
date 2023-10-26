package model

import "gorm.io/gorm"

// Video
type Video struct {
	gorm.Model
	// 视频作者信息
	AuthorId int64 `json:"author_id"`

	// 视频封面地址
	CoverURL string `json:"cover_url"`

	// 视频的评论总数
	CommentCount int64 `json:"comment_count"`
	// 视频的点赞总数
	FavoriteCount int64 `json:"favorite_count"`
	// 视频总收藏数
	CollectionCount int64 `json:"collection_count"`
	//弹幕总数
	DanmakuCount int64
	// 视频播放地址
	PlayURL string `json:"play_url"`
	// 视频标题
	Title string `json:"title"`
	// 视频所属领域
	CategoryId string `json:"category_id"`
}

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

// Collection 收藏夹模型
type Collection struct {
	gorm.Model
	Name      string  // 收藏夹名称
	IsPrivate bool    // 是否私有
	UserID    uint    // 所属用户的ID
	Videos    []Video `gorm:"many2many:collection_videos;"` // 包含的视频列表
}

// Danmaku 弹幕模型
type Danmaku struct {
	gorm.Model
	VideoID   uint   `json:"video_id"`  // 弹幕所属视频的ID
	UserID    uint   `json:"user_id"`   // 发布弹幕的用户ID
	Content   string `json:"content"`   // 弹幕内容
	Color     string `json:"color"`     // 弹幕颜色
	Timestamp uint   `json:"timestamp"` // 弹幕出现的时间戳
}
