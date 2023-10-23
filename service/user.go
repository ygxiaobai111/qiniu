package service

import (
	"context"
	"sync"
	"www.github.com/ygxiaobai111/qiniu/pkg/util"
	"www.github.com/ygxiaobai111/qiniu/repository/db/dao"
	"www.github.com/ygxiaobai111/qiniu/repository/db/model"
	"www.github.com/ygxiaobai111/qiniu/types"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

// UserRegister 用户注册
func (s *UserSrv) UserRegister(ctx context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	if exist {
		//err = errors.New("用户已经存在了")
		return
	}
	user := &model.User{

		UserName: req.UserName,
	}
	// 加密密码
	if err = user.SetPassword(req.Password); err != nil {
		util.LogrusObj.Error(err)
		return
	}

	// 创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	return
}
