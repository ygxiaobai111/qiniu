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
	err := dao.DB.Model(&model.Danmaku{}).Create(dan).Error
	if err != nil {
		return err
	}
	err = dao.UpdateVideoDanCount(context.Background(), dan.VideoID, 1)
	return err
}
func (dao *DanDao) GetDanByVid(id uint) (dans []*model.Danmaku, err error) {
	// CreateUser 创建数据
	err = dao.DB.Model(&model.Danmaku{}).Where("video_id=?", id).Find(&dans).Error

	return
}

// UpdateVideoCommentCount 更新视频的评论计数
func (dao *DanDao) UpdateVideoDanCount(ctx context.Context, videoID uint, increment int) error {
	result := dao.WithContext(ctx).Model(model.Video{}).Where("id = ?", videoID).UpdateColumn("danmaku_count", gorm.Expr("danmaku_count + ?", increment))
	return result.Error
}
