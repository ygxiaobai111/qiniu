package types

/*
存放请求结构体与返回结构体
*/
type UserServiceReq struct {
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
}
type UserLoginResp struct {
}
type UserRegisterReq struct {
	UserName string `form:"user_name" json:"username"`
	Password string `form:"password" json:"password"`
}

type UserTokenData struct {
	User         interface{} `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}

type UserLoginReq struct {
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
}

type UserInfoUpdateReq struct {
	UserName string `form:"user_name" json:"user_name"`
}

type UserInfoShowReq struct {
}

type UserFollowingReq struct {
	Id uint `json:"id" form:"id"`
}

type UserUnFollowingReq struct {
	Id uint `json:"id" form:"id"`
}

type UserInfoResp struct {
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	ID              int64  `json:"id"`               // 用户id
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Name            string `json:"name"`             // 用户名称
	TotalFavorited  int64  `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
}
