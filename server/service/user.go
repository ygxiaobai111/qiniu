package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"sync"
	e2 "www.github.com/ygxiaobai111/qiniu/server/pkg/e"
	"www.github.com/ygxiaobai111/qiniu/server/pkg/util"
	"www.github.com/ygxiaobai111/qiniu/server/repository/db/dao"
	"www.github.com/ygxiaobai111/qiniu/server/repository/db/model"
	"www.github.com/ygxiaobai111/qiniu/server/repository/es"
	"www.github.com/ygxiaobai111/qiniu/server/types"
)

/*
业务实现
*/
var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

// GetUserSrv 返回userSrv对象
func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

// UserRegister 用户注册方法 返回是 给用户的数据 与 错误信息
func (s *UserSrv) UserRegister(ctx context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {
	//获取user的数据库连接对象
	userDao := dao.NewUserDao(ctx)

	//查询该name是否存在于数据库
	_, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if err != nil {
		//打日志
		util.LogrusObj.Error(err)
		return
	}
	if exist {
		err = errors.New(e2.GetMsg(e2.ErrorAdminFindUser))
		return
	}
	user := &model.User{

		UserName:       req.UserName,
		PasswordDigest: req.Password,
		FollowCount:    0,
		FanCount:       0,
		Avatar:         model.Avatar,
	}
	// 加密密码
	if err = user.SetPassword(req.Password); err != nil {
		util.LogrusObj.Error(err)
		return
	}

	// 在数据库创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	es.UserCreate(user.ID, user.UserName)

	return
}
func (s *UserSrv) UserLogin(ctx context.Context, req *types.UserLoginReq) (resp interface{}, err error) {

	defer func() {
		// 返回时若err!=nil则写入日志
		if err != nil {
			util.LogrusObj.Error("<login> ", err, " [be from req]:", req)

		}
	}()
	// 数据验证
	if err = util.ValidateUser(req.UserName, req.Password); err != nil {
		return nil, errors.New(e2.GetMsg(e2.InvalidParams))
	}

	// 查询用户是否存在
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if err != nil {
		return nil, errors.New(e2.GetMsg(e2.ERROR))
	}
	if exist == false {
		return nil, errors.New(e2.GetMsg(e2.ErrorUserNotFound))
	}

	// 比较密码是否匹配
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(req.Password)); err != nil {
		return nil, errors.New(e2.GetMsg(e2.ErrorNotCompare))
	}

	// 签发token
	token, err := util.GenerateToken(user.ID, req.UserName, 0)
	if err != nil {
		return nil, errors.New(e2.GetMsg(e2.ERROR))
	}
	userResp := &types.UserInfoResp{
		ID:            user.ID,
		Name:          user.UserName,
		Avatar:        user.Avatar,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FanCount,
	}
	return &types.TokenData{
		User:  userResp,
		Token: token,
	}, nil
}
func (s *UserSrv) UserInfo(ctx context.Context, req *types.UserInfoShowReq, uid uint) (resp interface{}, err error) {
	udao := dao.NewUserDao(ctx)

	data, err := udao.GetUserById(req.UserId)
	if err != nil {
		return
	}

	user := BuildUser(ctx, data, uid)
	resp = user
	return
}

func (s *UserSrv) UserAction(ctx context.Context, req *types.UserFollowingReq, userId uint) (resp interface{}, err error) {

	dao := dao.NewUserDao(ctx)
	switch req.Type {
	case 1:
		err = dao.Follow(userId, req.Id)
	case 2:
		err = dao.Unfollow(userId, req.Id)
	default:
		err = errors.New(e2.GetMsg(e2.InvalidParams))
	}
	return
}
func (s *UserSrv) UserFollow(ctx context.Context, req *types.UserFollowReq, uid uint) (resp interface{}, err error) {
	udao := dao.NewUserDao(ctx)

	users, err := udao.GetFollowList(req.UserId)
	if err != nil {
		return
	}
	r := BuildUsers(ctx, users, uid)
	resp = types.DataList{
		Item:  r,
		Total: uint(len(r)),
	}
	return

}
func (s *UserSrv) UserFollower(ctx context.Context, req *types.UserFollowerReq, uid uint) (resp interface{}, err error) {
	udao := dao.NewUserDao(ctx)
	users, err := udao.GetFollowerList(int64(req.UserId))
	if err != nil {
		return
	}
	r := BuildUsers(ctx, users, uid)
	resp = types.DataList{
		Item:  r,
		Total: uint(len(r)),
	}
	return
}
func (s *UserSrv) UserFriend(ctx context.Context, req *types.UserFriendReq, uid uint) (resp interface{}, err error) {
	udao := dao.NewUserDao(ctx)
	users, err := udao.GetFriendList(req.UserId)
	if err != nil {
		return
	}
	r := BuildUsers(ctx, users, uid)
	resp = types.DataList{
		Item:  r,
		Total: uint(len(r)),
	}
	return
}

func BuildUser(ctx context.Context, data *model.User, myUId uint) (user *types.UserInfoResp) {

	vdao := dao.NewVideoDao(ctx)
	fdao := dao.NewFavDao(ctx)
	udao := dao.NewUserDao(ctx)
	workcount, _ := vdao.GetVideoCountByUId(data.ID)
	favcount, _ := fdao.GetFavoriteCount(ctx, data.ID)
	isFollow, _ := udao.IsFollow(myUId, data.ID)
	user = &types.UserInfoResp{
		Avatar:          data.Avatar,
		BackgroundImage: "",
		FavoriteCount:   favcount,
		FollowCount:     data.FollowCount,
		FollowerCount:   data.FanCount,
		ID:              data.ID,
		IsFollow:        isFollow,
		Name:            data.UserName,
		WorkCount:       workcount,
	}
	return
}
func BuildUsers(ctx context.Context, datas []*model.User, uid uint) (users []*types.UserInfoResp) {

	for _, data := range datas {
		user := BuildUser(ctx, data, uid)
		users = append(users, user)
	}
	return
}
