package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
	"www.github.com/ygxiaobai111/qiniu/server/repository/db/model"
)

type VideoDao struct {
	*gorm.DB
}

// NewUserDao 获取 UserDao 的函数
func NewVideoDao(ctx context.Context) *VideoDao {
	return &VideoDao{NewDBClient(ctx)}
}

// Create a new video
func (dao *VideoDao) CreateVideo(video *model.Video) error {
	return dao.Model(model.Video{}).Create(video).Error
}

func (dao *VideoDao) VideoFeed(cate int64) (videos []*model.Video, err error) {
	err = dao.Order("RAND()").Limit(10).Find(&videos).Error
	return
}

// Update video
func (dao *VideoDao) UpdateVideo(video *model.Video) error {
	return dao.Model(model.Video{}).Save(video).Error
}

// Delete video by ID
func (dao *VideoDao) DeleteVideoByID(id uint) error {
	return dao.Delete(&model.Video{}, id).Error
}

// GetVideoByUId 根据uid获取video
func (dao *VideoDao) GetVideoByUId(id uint) (videos []*model.Video, err error) {
	err = dao.DB.Model(&model.Video{}).Where("user_id = ?", id).Find(&videos).Error
	return
}

// GetVideoCountByUId 根据uid获取video数量
func (dao *VideoDao) GetVideoCountByUId(id uint) (count int64, err error) {
	err = dao.DB.Model(&model.Video{}).Where("user_id = ?", id).Count(&count).Error
	return
}

// GetHotVideo 获取7天内的点赞最高的视频
func (dao *VideoDao) GetHotVideo() (videos []*model.Video, err error) {
	// 计算7天前的时间
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	// 查询视频记录，按点赞量倒序排序，限制返回前20条记录
	err = dao.Where("created_at >= ?", sevenDaysAgo).Order("favorite_count DESC").Limit(20).Find(&videos).Error
	return
}

// GetVideoById 根据id获取video
func (dao *VideoDao) GetVideoById(id uint) (video *model.Video, err error) {
	err = dao.DB.Model(&model.Video{}).Where("id = ?", id).First(&video).Error
	return
}
