package service

import (
	"context"
	"github.com/h2non/filetype"
	"mime/multipart"
	"sync"
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
	return

}
func (s *VideoSrv) VideoUpdatePublish(ctx context.Context, req *types.VideoUpdatePublish) (resp interface{}, err error) {
	return

}
func (s *VideoSrv) VideoDelPublish(ctx context.Context, req *types.VideoDelPublish) (resp interface{}, err error) {
	return

}
func (s *VideoSrv) VideoBefore(ctx context.Context, req *types.VideoBefore) (resp interface{}, err error) {
	return

}
