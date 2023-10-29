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

// 从收藏夹中删除
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

type GetFavResp struct {
	VideoId int64 `json:"video_id"`
	//视频创建时间
	CreateTime int64 `json:"create_time"`
	AuthorId   int64 `json:"author_id"`
	// 视频作者
	AuthorName string `json:"author_name"`
	//播放总数
	PlayCount int64 `json:"collection_count"`
	// 视频封面地址
	CoverURL string `json:"cover_url"`
	//播放地址
	PlayURL string `json:"play_url"`
	// 视频的点赞总数
	FavoriteCount int64 `json:"favorite_count"`
	// 视频总收藏数
	CollectionCount int64 `json:"collection_count"`
	// 视频标题
	Title string `json:"title"`
	// 视频所属领域
	Category string `json:"category_id"`
}
type GetFavlistResp struct {
	//收藏夹创建人
	UserName string `json:"user_name"`
	//收藏夹名
	CollectionName string        `json:"collection_name"`
	CreateTime     int64         `json:"create_time"`
	Favlist        []*GetFavResp `json:"favlist"`
	Total          int64         `json:"total"`
}
type GetFavoriteResp struct {
	// 视频作者
	AuthorName int64 `json:"author_name"`
	//播放总数
	PlayCount int64 `json:"collection_count"`
	// 视频封面地址
	CoverURL string `json:"cover_url"`
	// 视频标题
	Title string `json:"title"`
}
type GetCommentResp struct {
	UserId int64 `json:"user_id"`
	//用户名
	UserName string `json:"user_name"`
	//用户头像
	Avatar string `json:"avatar"`
	// 评论内容
	Content string `json:"content"`
	// 评论发布日期，格式 mm-dd
	CreateTime int64 `json:"create_time"`
}
type GetBarrageResp struct {
	Content   string `json:"content"`   // 弹幕内容
	Color     string `json:"color"`     // 弹幕颜色
	Timestamp uint   `json:"timestamp"` // 弹幕出现的时间戳
}
