package models

type VideoModel struct {
	UserId         int64  `json:"user_id"`
	VideoId        int64  `json:"video_id"`
	TagId          int64  `json:"tag_id"`
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
