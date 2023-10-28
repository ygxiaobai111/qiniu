package service

import (
	"context"
	"github.com/h2non/filetype"
	"mime/multipart"
	"sync"
	dao2 "www.github.com/ygxiaobai111/qiniu/server/repository/db/dao"
	"www.github.com/ygxiaobai111/qiniu/server/types"
)

/*
业务实现
*/
var VideoSrvIns *VideoSrv
var VideoSrvOnce sync.Once

type VideoSrv struct {
}

// GetVideoSrv 返回userSrv对象
func GetVideoSrv() *VideoSrv {
	VideoSrvOnce.Do(func() {
		VideoSrvIns = &VideoSrv{}
	})
	return VideoSrvIns
}

func (s *VideoSrv) VideoCreate(ctx context.Context, video *multipart.FileHeader, image *multipart.FileHeader) (resp interface{}, err error) {
	return
}
func (s *VideoSrv) IsVideoFile(file *multipart.FileHeader) (bool, error) {
	// 打开上传的文件
	uploadedFile, err := file.Open()
	if err != nil {
		return false, err
	}
	defer uploadedFile.Close()

	// 读取文件的前 261 字节
	buffer := make([]byte, 261)
	_, err = uploadedFile.Read(buffer)
	if err != nil {
		return false, err
	}

	// 使用文件类型检测库来判断文件类型
	kind, _ := filetype.Match(buffer)

	// 如果文件类型是图片类型，返回true
	if kind.MIME.Type == "video" {
		return true, nil
	}

	// 否则，不是图片文件
	return false, nil
}

func (s *VideoSrv) IsImageFile(file *multipart.FileHeader) (bool, error) {
	// 打开上传的文件
	uploadedFile, err := file.Open()
	if err != nil {
		return false, err
	}
	defer uploadedFile.Close()

	// 读取文件的前 261 字节
	buffer := make([]byte, 261)
	_, err = uploadedFile.Read(buffer)
	if err != nil {
		return false, err
	}

	// 使用文件类型检测库来判断文件类型
	kind, _ := filetype.Match(buffer)

	// 如果文件类型是图片类型，返回true
	if kind.MIME.Type == "image" {
		return true, nil
	}

	// 否则，不是图片文件
	return false, nil
}
func (s *VideoSrv) VideoSearch(ctx context.Context, req *types.VideoSearch) (resp interface{}, err error) {
	return

}
func (s *VideoSrv) VideoChannel(ctx context.Context, req *types.VideoChannel) (resp interface{}, err error) {
	return

}
func (s *VideoSrv) VideoGetPublish(ctx context.Context, req *types.VideoGetPublish) (resp interface{}, err error) {
	vdao := dao2.NewVideoDao(ctx)
	udao := dao2.NewUserDao(ctx)
	cdao := dao2.NewCateDao(ctx)
	videos, err := vdao.GetVideoByUId(req.UserId)
	if err != nil {
		return
	}
	user, _ := udao.GetUserById(uint(videos[0].AuthorId))
	var r []types.GetFavResp
	for _, video := range videos {
		c, _ := cdao.GetCateById(int64(video.CategoryId))
		data := types.GetFavResp{
			CreateTime:      video.CreatedAt.Unix(),
			AuthorName:      user.UserName,
			PlayCount:       0,
			CoverURL:        video.CoverURL,
			PlayURL:         video.PlayURL,
			FavoriteCount:   video.FavoriteCount,
			CollectionCount: video.CollectionCount,
			Title:           video.Title,
			Category:        c.CategoryName,
		}
		r = append(r, data)

	}
	resp = types.DataList{
		Item:  r,
		Total: uint(len(r)),
	}
	return

}
func (s *VideoSrv) VideoUpdatePublish(ctx context.Context, req *types.VideoUpdatePublish) (resp interface{}, err error) {
	dao := dao2.NewVideoDao(ctx)
	video, err := dao.GetVideoById(req.VideoId)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		video.Title = req.Title
	}
	if req.CategoryId != 0 {
		video.CategoryId = req.CategoryId
	}
	err = dao.UpdateVideo(video)
	return

}
func (s *VideoSrv) VideoDelPublish(ctx context.Context, req *types.VideoDelPublish) (resp interface{}, err error) {
	dao := dao2.NewVideoDao(ctx)
	_, err = dao.GetVideoById(req.VideoId)
	if err != nil {
		return nil, err
	}
	err = dao.DeleteVideoByID(req.VideoId)
	return

}

// 历史视频
func (s *VideoSrv) VideoBefore(ctx context.Context, req *types.VideoBefore) (resp interface{}, err error) {
	return

}
