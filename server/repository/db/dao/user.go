package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"www.github.com/ygxiaobai111/qiniu/server/repository/db/model"
)

type UserDao struct {
	*gorm.DB
}

// NewUserDao 获取 UserDao 的函数
func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// ExistOrNotByUserName 根据userName查询数据库中是否存在该名字
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).First(&user).Count(&count).Error
	//err == gorm.ErrRecordNotFound 代表未发现该条数据
	if count == 0 || err == gorm.ErrRecordNotFound {
		return nil, false, nil
	}
	return user, true, nil
}

// CreateUser 创建数据
func (dao *UserDao) CreateUser(user *model.User) (err error) {
	return dao.DB.Model(&model.User{}).Create(user).Error

}

// GetUserById 根据id获取user
func (dao *UserDao) GetUserById(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id = ?", id).First(&user).Error
	return
}

// UpdateUserById 根据id修改user信息
func (dao *UserDao) UpdateUserById(uId uint, user *model.User) (err error) {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).Updates(&user).Error
}

// Follow 关注
func (dao *UserDao) Follow(userId, toUserId int64) error {
	fmt.Printf("userid：%v，toUserId：%v", userId, toUserId)
	return dao.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Model(&model.User{Model: gorm.Model{ID: uint(userId)}}).
		Association("Follows").
		Append(&model.User{
			Model: gorm.Model{ID: uint(toUserId)},
		})
}

// Unfollow 取消关注
func (dao *UserDao) Unfollow(userId, toUserId int64) error {
	return dao.DB.
		Model(&model.User{
			Model: gorm.Model{ID: uint(userId)},
		}).
		Association("Follows").
		Delete(&model.User{
			Model: gorm.Model{ID: uint(toUserId)},
		})
}

// GetFollowList 获取关注列表
func (dao *UserDao) GetFollowList(userId int64) ([]*model.User, error) {
	var user *model.User
	if err := dao.
		Where("id = ?", userId).
		Preload("Follows").
		Find(&user).Error; err != nil {
		return nil, err
	}
	return user.Follows, nil
}

// GetFollowerList 获取粉丝列表
func (dao *UserDao) GetFollowerList(userId int64) ([]*model.User, error) {
	var user *model.User
	if err := dao.
		Where("id = ?", userId).
		Preload("Fans").
		Find(&user).Error; err != nil {
		return nil, err
	}
	return user.Fans, nil
}

// GetFriendList 获取好友列表
func (dao *UserDao) GetFriendList(userId int64) ([]*model.User, error) {

	// 获取粉丝的交集
	var friends []*model.User
	dao.
		Raw("SELECT * FROM user WHERE id IN (SELECT follow_id FROM follows WHERE user_id = ?) AND id IN (SELECT user_id FROM follows WHERE follow_id = ?)",
			userId, userId).
		Scan(&friends)
	return friends, nil
}
