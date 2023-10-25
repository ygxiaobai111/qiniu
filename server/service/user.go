package service

import (
	"context"
	"errors"
	"sync"
	"www.github.com/ygxiaobai111/qiniu/server/pkg/util"
	"www.github.com/ygxiaobai111/qiniu/server/repository/db/dao"
	"www.github.com/ygxiaobai111/qiniu/server/repository/db/model"
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
		err = errors.New("用户已经存在了")
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

	// 在数据库创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}

	return
}
func (s *UserSrv) UserLogin(ctx context.Context, req *types.UserLoginReq) (resp interface{}, err error) {
	return
}
