package model

import "gorm.io/gorm"

// Video
type Video struct {
	// 视频作者信息
	AuthorId int64 `json:"author_id"`
	// 视频的评论总数
	CommentCount int64 `json:"comment_count"`
	// 视频封面地址
	CoverURL string `json:"cover_url"`
	// 视频的点赞总数
	FavoriteCount int64 `json:"favorite_count"`
	// 视频唯一标识
	ID int64 `json:"id"`

	// true-已点赞，false-未点赞
	IsFavorite bool `json:"is_favorite"`
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
	// 评论内容
	Content string `json:"content"`
	// 评论发布日期，格式 mm-dd
	CreateDate string `json:"create_date"`
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
