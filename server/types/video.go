package types

type VideoCreateReq struct {

	// 视频标题
	Title string `json:"title"`
	// 视频所属领域
	CategoryId string `json:"category_id"`
}
type VideoUpdateReq struct {
	// 视频标题
	Title string `json:"title"`
	// 视频所属领域
	CategoryId string `json:"category_id"`
}

type VideoSearch struct {
	Text string `json:"text"`
}
