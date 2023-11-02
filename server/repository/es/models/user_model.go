package models

type UserModel struct {
	UserId   int64  `json:"user_id"`
	NickName string `json:"nick_name"`
}

func (UserModel) UserIndex() string {
	return "user_index"
}

func (UserModel) UserMapping() string {
	return `
{
  "mappings": {
    "properties": {
      "nick_name": { 
        "type": "text"
      },
    
      "user_id": {
        "type": "integer"
      }
      }
    }
  }
}
`
}
