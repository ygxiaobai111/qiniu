package dao

import (
	"context"
	"gorm.io/gorm"
	"www.github.com/ygxiaobai111/qiniu/server/repository/db/model"
)

type FavDao struct {
	*gorm.DB
}

func NewFavDao(ctx context.Context) *FavDao {
	return &FavDao{NewDBClient(ctx)}
}

// CreateFav 点赞
func (dao *FavDao) CreateFav(ctx context.Context, videoId uint, userId uint) error {
	return dao.WithContext(ctx).Model(&model.Fav{}).Create(&model.Fav{UserId: userId, VideoId: videoId}).Error
}

// DeleteFav 取消点赞
func (dao *FavDao) DeleteFav(ctx context.Context, videoId uint, userId uint) error {
	return dao.WithContext(ctx).Model(&model.Fav{}).Where("user_id = ? and video_id = ?", userId, videoId).Delete(&model.Fav{}).Error
}

// IsFavorite 判断是否点赞
func (dao *FavDao) IsFavorite(ctx context.Context, videoId uint, userId uint) (bool, error) {
	var fav model.Fav
	result := dao.WithContext(ctx).Where("user_id = ? and video_id = ?", userId, videoId).First(&fav)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

// GetFavoriteCount 获取用户点赞数量(给别人的赞)
func (dao *FavDao) GetFavoriteCount(ctx context.Context, userId uint) (int64, error) {
	var favCount int64
	result := dao.WithContext(ctx).Model(&model.Fav{}).Where("user_id = ?", userId).Count(&favCount)
	if result.Error != nil {
		return 0, result.Error
	}
	return favCount, nil
}

// GetSingleVideoFavoriteCount 获取单个视频点赞数量
func (dao *FavDao) GetSingleVideoFavoriteCount(ctx context.Context, videoId uint) (int32, error) {
	var favCount int32
	result := dao.WithContext(ctx).Table("video").Where("id = ?", videoId).Select("favorite_count").Scan(&favCount)
	if result.Error != nil {
		return 0, result.Error
	}
	return favCount, nil
}

// UpdateFavoriteCountByVideoId 更新mysql中的值
func (dao *FavDao) UpdateFavoriteCountByVideoId(videoID uint, favoriteCount int64) error {
	return dao.Model(&model.Video{}).Where("id = ?", videoID).Update("favorite_count", favoriteCount).Error
}

// ListFav 获取用户喜欢列表
func (dao *FavDao) ListFav(ctx context.Context, userId uint) (favs []*model.Video) {

	dao.WithContext(ctx).Where("user_id = ? ", userId).Find(&favs)
	var videoIDs []uint
	for _, rel := range favs {
		videoIDs = append(videoIDs, rel.ID)
	}
	if len(videoIDs) == 0 {
		return []*model.Video{}
	}
	var videos []*model.Video
	dao.WithContext(ctx).Find(&videos, videoIDs)
	return videos
}
