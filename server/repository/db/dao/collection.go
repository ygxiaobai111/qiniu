package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
	"www.github.com/ygxiaobai111/qiniu/server/repository/db/model"
)

type CollectionDao struct {
	*gorm.DB
}

func NewCollectionDao(ctx context.Context) *CollectionDao {
	return &CollectionDao{NewDBClient(ctx)}
}

// Update 更新收藏夹信息
func (dao *CollectionDao) Update(collection *model.Collection) (err error) {
	err = dao.Save(collection).Error

	return
}

// AddVideo 添加视频到收藏夹
func (dao *CollectionDao) AddVideo(collection *model.Collection, videoID uint) error {
	// 检查视频是否存在
	var video *model.Video
	err := db.First(&video, videoID).Error
	if err != nil {

		return err

	}

	// 将视频添加到收藏夹的 Videos 数组中
	collection.Videos = append(collection.Videos, video)

	// 保存收藏夹
	err = dao.Update(collection)
	if err != nil {

		return err
	}

	return nil
}

// RemoveVideo 从收藏夹中移除视频
func (dao *CollectionDao) RemoveVideo(collection *model.Collection, videoID uint) error {
	// 查找要移除的视频在 Videos 数组中的索引
	index := -1
	for i, video := range collection.Videos {
		if video.ID == videoID {
			index = i
			break
		}
	}
	if index == -1 {
		return fmt.Errorf("视频不存在于收藏夹中")
	}

	// 从 Videos 数组中移除视频
	collection.Videos = append(collection.Videos[:index], collection.Videos[index+1:]...)

	// 保存收藏夹
	err := dao.Save(collection).Error
	if err != nil {
		return err
	}

	return nil
}

func (dao *CollectionDao) Create(collection *model.Collection) (err error) {
	err = dao.Model(model.Collection{}).Create(collection).Error
	log.Println("err:", err)
	if err != nil {

		return
	}
	err = dao.UpdateUserCollectionCount(context.Background(), int64(collection.UserID), -1)
	return
}
func (dao *CollectionDao) GetCollections(userId int64) (collections []*model.Collection, err error) {
	err = dao.Model(model.Collection{}).Where("user_id=?", userId).Find(&collections).Error

	return
}
func (dao *CollectionDao) GetCollection(id int64) (collection *model.Collection, err error) {
	err = dao.Model(model.Collection{}).Where("id=?", id).First(&collection).Error

	return
}

// 删除视频收藏夹
func (dao *CollectionDao) DelCollection(id, userId int64) (err error) {
	err = dao.Model(model.Collection{}).Where("id=? and user_id=?", id, userId).Delete(model.Collection{}).Error
	if err != nil {
		return
	}
	err = dao.UpdateUserCollectionCount(context.Background(), userId, -1)
	return
}

// UpdateVideoCommentCount 更新用户的收藏计数
func (dao *CollectionDao) UpdateUserCollectionCount(ctx context.Context, videoID int64, increment int) error {
	result := dao.WithContext(ctx).Model(model.User{}).Where("id = ?", videoID).UpdateColumn("collection_count", gorm.Expr("collection_count + ?", increment))
	return result.Error
}
