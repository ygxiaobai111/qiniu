package dao

import (
	"context"
	"gorm.io/gorm"
	"www.github.com/ygxiaobai111/qiniu/server/repository/db/model"
)

type CateDao struct {
	*gorm.DB
}

// NewUserDao 获取 UserDao 的函数
func NewCateDao(ctx context.Context) *CateDao {
	return &CateDao{NewDBClient(ctx)}
}

// 通过id获取分类
func (dao *CateDao) GetCateById(id uint) (cate *model.Category, err error) {
	err = dao.Model(model.Video{}).Where("id=?", id).Find(cate).Error
	return
}

// 获取分类
func (dao *CateDao) GetCates() (cate *[]model.Category, err error) {
	err = dao.Model(model.Video{}).Find(cate).Error
	return
}
