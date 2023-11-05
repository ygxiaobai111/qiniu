package models

type VideoModel struct {
	UserId         uint   `json:"user_id"`
	VideoId        uint   `json:"video_id"`
	TagId          uint   `json:"tag_id"`
	NickVideoTitle string `json:"nick_video_title"`
}

func (VideoModel) VideoIndex() string {
	return "video_index"
}

func (VideoModel) VideoMapping() string {
	return `{
  "mappings": {
    "properties": {
      "user_id": { 
        "type": "long"
      },
      "video_id": { 
        "type": "long" 
      },
      "tag_id": {
        "type": "long"
      },
      "nick_video_title":{
        "type": "text"
      }
    }
  }
}`
}
