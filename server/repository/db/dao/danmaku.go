package dao

import (
	"context"
	"gorm.io/gorm"
	"www.github.com/ygxiaobai111/qiniu/server/repository/db/model"
)

type DanDao struct {
	*gorm.DB
}

// NewUserDao 获取 UserDao 的函数
func NewDanDao(ctx context.Context) *DanDao {
	return &DanDao{NewDBClient(ctx)}
}

func (dao *DanDao) Create(dan model.Danmaku) error {
	// CreateUser 创建数据
	return dao.DB.Model(&model.Danmaku{}).Create(dan).Error
}
func (dao *DanDao) GetDanByVid(id int64) (Dans []*model.Danmaku, err error) {
	// CreateUser 创建数据
	err = dao.DB.Model(&model.Danmaku{}).Where("video_id=?", id).Find(&Dans).Error
	return
}
