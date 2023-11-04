package types

type VideoCreateReq struct {
	Title      string `json:"title" form:"title"`
	CategoryId uint   `json:"category_id" form:"category_id"`
}

type VideoUpdateReq struct {
	Title      string `json:"title" form:"title"`
	CategoryId uint   `json:"category_id" form:"category_id"`
	Page
}

type VideoSearch struct {
	Text string `json:"text" form:"text"`
	Type int    `json:"type" form:"type"`
	Page
}
type VideoChannel struct {
	ChannelId uint `json:"channel_id" form:"channel_id"`
	Page
}

type VideoGetPublish struct {
	UserId uint `json:"user_id" form:"user_id"`
}

type VideoUpdatePublish struct {
	Title      string `json:"title" form:"title"`
	CategoryId uint   `json:"category_id" form:"category_id"`
	VideoId    uint   `json:"video_id" form:"video_id"`
}

type VideoDelPublish struct {
	VideoId uint `json:"video_id" form:"video_id"`
}

type VideoBefore struct {
	UserId uint `json:"user_id" form:"user_id"`
}
