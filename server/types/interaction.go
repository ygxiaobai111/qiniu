package types

// 收藏夹
type GetFavlistReq struct {
	UserId    int64 `json:"user_id"`
	FavlistId int64 `json:"favlist_id"`
}

// 用户喜欢列表
type GetFavoriteReq struct {
	UserId int64 `json:"user_id"`
}

// 评论列表
type GetCommentReq struct {
	VideoId int64 `json:"video_id"`
}

// 弹幕获取
type GetBarrageReq struct {
	VideoId int64 `json:"video_id"`
}

// 创建收藏夹
type FavlisCreatetReq struct {
	FavlistName string `json:"favlist_name"`
	Type        int    `json:"type"`
}

// 加入收藏夹
type FavlistAddReq struct {
	FavlistId int64 `json:"favlist_id"`
	VideoId   int64 `json:"video_id"`
}
type FavlistDelReq struct {
	FavlistId int64 `json:"favlist_id"`
	VideoId   int64 `json:"video_id"`
}

// 删除收藏夹
type DelFavlistReq struct {
	FavlistId int64 `json:"favlist_id"`
}

// 评论
type CommentCreateReq struct {
	VideoId int64  `json:"video_id"`
	Content string `json:"content"` // 内容
}

// 点赞/取消点赞
type FavoriteReq struct {
	VideoId int64 `json:"user_id"`
	Type    int   `json:"type"`
}

// 弹幕发送
type BarrageReq struct {
	VideoID uint `json:"video_id"` // 弹幕所属视频的ID

	Content   string `json:"content"`   // 弹幕内容
	Color     string `json:"color"`     // 弹幕颜色
	Timestamp uint   `json:"timestamp"` // 弹幕出现的时间戳

}
