package types

// 收藏夹
type GetFavlistReq struct {
	UserId    uint `json:"user_id" form:"user_id"`
	FavlistId uint `json:"favlist_id" form:"favlist_id"`
}

// 用户喜欢列表
type GetFavoriteReq struct {
	UserId uint `json:"user_id" form:"user_id"`
}

// 评论列表
type GetCommentReq struct {
	VideoId uint `json:"video_id" form:"video_id"`
}

// 弹幕获取
type GetBarrageReq struct {
	VideoId uint `json:"video_id" form:"video_id"`
}

// 创建收藏夹
type FavlisCreatetReq struct {
	FavlistName string `json:"favlist_name" form:"favlist_name"`
	Type        int    `json:"type" form:"type"`
}

// 加入收藏夹
type FavlistAddReq struct {
	FavlistId uint `json:"favlist_id" form:"favlist_id"`
	VideoId   uint `json:"video_id" form:"video_id"`
}

// 从收藏夹中删除
type FavlistDelReq struct {
	FavlistId uint `json:"favlist_id" form:"favlist_id"`
	VideoId   uint `json:"video_id" form:"video_id"`
}

// 删除收藏夹
type DelFavlistReq struct {
	FavlistId uint `json:"favlist_id" form:"favlist_id"`
}

// 评论
type CommentCreateReq struct {
	VideoId uint   `json:"video_id" form:"video_id"`
	Content string `json:"content" form:"content"` // 内容
}

// 点赞/取消点赞
type FavoriteReq struct {
	VideoId uint `json:"user_id" form:"user_id"`
	Type    int  `json:"type" form:"type"`
}

// 弹幕发送
type BarrageReq struct {
	VideoID   uint   `json:"video_id" form:"video_id"`   // 弹幕所属视频的 ID
	Content   string `json:"content" form:"content"`     // 弹幕内容
	Color     string `json:"color" form:"color"`         // 弹幕颜色
	Timestamp uint   `json:"timestamp" form:"timestamp"` // 弹幕出现的时间戳
}

// 用户画像构建请求
type PersonasReq struct {
	CategoryId uint `json:"category_id" form:"category_id"`
}
type GetFavResp struct {
	VideoId         uint   `json:"video_id" form:"video_id"`
	CreateTime      int64  `json:"create_time" form:"create_time"`
	AuthorId        uint   `json:"author_id" form:"author_id"`
	AuthorName      string `json:"author_name" form:"author_name"`
	PlayCount       int64  `json:"collection_count" form:"collection_count"`
	CoverURL        string `json:"cover_url" form:"cover_url"`
	PlayURL         string `json:"play_url" form:"play_url"`
	FavoriteCount   int64  `json:"favorite_count" form:"favorite_count"`
	CollectionCount int64  `json:"collection_count" form:"collection_count"`
	Title           string `json:"title" form:"title"`
	Category        string `json:"category" form:"category"`
	IsFav           bool   `json:"is_fav" form:"is_fav"`
}

type GetFavlistResp struct {
	UserName       string        `json:"user_name" form:"user_name"`
	CollectionName string        `json:"collection_name" form:"collection_name"`
	CreateTime     int64         `json:"create_time" form:"create_time"`
	Favlist        []*GetFavResp `json:"favlist" form:"favlist"`
	Total          int64         `json:"total" form:"total"`
}

type GetCommentResp struct {
	UserId     uint   `json:"user_id" form:"user_id"`
	UserName   string `json:"user_name" form:"user_name"`
	Avatar     string `json:"avatar" form:"avatar"`
	Content    string `json:"content" form:"content"`
	CreateTime int64  `json:"create_time" form:"create_time"`
}

type GetBarrageResp struct {
	Content   string `json:"content" form:"content"`
	Color     string `json:"color" form:"color"`
	Timestamp uint   `json:"timestamp" form:"timestamp"`
}
