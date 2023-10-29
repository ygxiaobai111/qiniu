package types

type VideoCreateReq struct {

	// 视频标题
	Title string `json:"title"`
	// 视频所属领域
	CategoryId uint `json:"category_id"`
}
type VideoUpdateReq struct {
	// 视频标题
	Title string `json:"title"`
	// 视频所属领域
	CategoryId uint `json:"category_id"`
}

type VideoSearch struct {
	//视频关键字
	Text string `json:"text"`
}
type VideoChannel struct {
	//视频分类id
	ChannelId uint `json:"channel_id"`
}
type VideoGetPublish struct {
	//用户id
	UserId uint `json:"user_id"`
}
type VideoUpdatePublish struct {
	// 视频标题
	Title string `json:"title"`
	// 视频所属领域
	CategoryId uint `json:"category_id"`
	//视频id
	VideoId uint `json:"video_id"`
}
type VideoDelPublish struct {
	//视频id
	VideoId uint `json:"video_id"`
}
type VideoBefore struct {
	//用户id
	UserId uint `json:"user_id"`
}
