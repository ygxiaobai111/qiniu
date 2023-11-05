package model

import "gorm.io/gorm"

// Video
type Video struct {
	gorm.Model
	// 视频作者信息
	AuthorId uint `json:"author_id"`

	// 视频封面地址
	CoverURL string `json:"cover_url"`

	// 视频的评论总数
	CommentCount int64 `json:"comment_count"`
	// 视频的点赞总数
	FavoriteCount int64 `json:"favorite_count"`
	// 视频总收藏数
	CollectionCount int64 `json:"collection_count"`
	//弹幕总数
	DanmakuCount int64 `json:"danmaku_count"`
	// 视频播放地址
	PlayURL string `json:"play_url"`
	// 视频标题
	Title string `json:"title"`
	// 视频所属领域
	CategoryId uint `json:"category_id"`
}
