package service

import (
	"context"
	"errors"
	"github.com/h2non/filetype"
	"math"
	"mime/multipart"
	"sync"
	"time"
	e2 "www.github.com/ygxiaobai111/qiniu/server/pkg/e"
	"www.github.com/ygxiaobai111/qiniu/server/repository/cache"
	dao2 "www.github.com/ygxiaobai111/qiniu/server/repository/db/dao"
	"www.github.com/ygxiaobai111/qiniu/server/repository/db/model"
	"www.github.com/ygxiaobai111/qiniu/server/repository/oss"
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

func (s *VideoSrv) VideoCreate(ctx context.Context, req types.VideoCreateReq, videoF *multipart.FileHeader, image *multipart.FileHeader, userId int64) (resp interface{}, err error) {
	vdao := dao2.NewVideoDao(ctx)

	fileF, err := videoF.Open()
	f := make([]byte, videoF.Size)

	fileF.Read(f)
	url, err := oss.Add(100001, "testTitle", f)
	if err != nil {
		return
	}

	video := &model.Video{
		AuthorId:        userId,
		CoverURL:        url + "?vframe/jpg/offset/1",
		CommentCount:    0,
		FavoriteCount:   0,
		CollectionCount: 0,
		DanmakuCount:    0,
		PlayURL:         url,
		Title:           req.Title,
		CategoryId:      req.CategoryId,
	}
	err = vdao.CreateVideo(video)
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
	//通过用户id获取视频
	videos, err := vdao.GetVideoByUId(req.UserId)
	if err != nil {
		return
	}
	r := BuildVideos(ctx, videos)

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

// 视频流
func (s *VideoSrv) VideoFeed(ctx context.Context, userId int64) (resp interface{}, err error) {
	dao := dao2.NewVideoDao(ctx)
	var videos []*model.Video
	//随机推荐
	randVideos, err := dao.VideoFeed(0)
	if err != nil {
		return
	}
	videos = append(videos, randVideos...)
	//用户推荐
	if userId != 0 {
		//根据用户画像推荐
		videoOfPersonaIds := cache.GetTopTags(ctx, cache.PersonasKey(uint(userId)))
		videoOfPersonas, _ := dao.GetVideoByIds(videoOfPersonaIds)
		if len(videoOfPersonas) > 0 {
			videos = append(videos, videoOfPersonas...)
		}

	}
	r := BuildVideos(ctx, randVideos)

	resp = types.DataList{
		Item:  r,
		Total: uint(len(r)),
	}
	return

}

// VideoHot 热门视频
func (s *VideoSrv) VideoHot(ctx context.Context, userId int64) (resp interface{}, err error) {
	dao := dao2.NewVideoDao(ctx)
	videosId, err := cache.GetTop30Videos(ctx)
	if err != nil {
		return
	}
	videos, err := dao.GetVideoByIds(videosId)
	if err != nil {
		return
	}
	r := BuildVideos(ctx, videos)

	resp = types.DataList{
		Item:  r,
		Total: uint(len(r)),
	}
	return

}

// 视频热门队列
func HotVideo(ctx context.Context) (err error) {
	dao := dao2.NewVideoDao(ctx)
	videos, err := dao.GetHotVideo()
	if err != nil {
		return
	}
	for _, video := range videos {
		var score float64
		score = Score(video)
		cache.AddPopularVideo(ctx, int64(video.ID), score, video.CreatedAt)
	}
	return nil
}
func Score(video *model.Video) float64 {
	now := time.Now()
	days := int(math.Floor(now.Sub(video.CreatedAt).Hours() / 24))

	var baseScore float64
	if days <= 7 {
		baseScore = float64(7 - days)
	} else {
		baseScore = 0
	}

	return baseScore + float64(video.FavoriteCount)*0.2 + float64(video.CollectionCount)*0.3 + float64(video.DanmakuCount)*0.1
}
func BuildVideos(ctx context.Context, videos []*model.Video) (r []*types.GetFavResp) {

	//因为是一个作者，提出来共用

	for _, video := range videos {
		data, _ := BuildVideo(ctx, video)

		r = append(r, data)

	}

	return
}
func BuildVideo(ctx context.Context, video *model.Video) (data *types.GetFavResp, err error) {
	if video == nil {
		return nil, errors.New(e2.GetMsg(e2.ERRORNULL))
	}
	udao := dao2.NewUserDao(ctx)
	cdao := dao2.NewCateDao(ctx)

	var username, categoryName string
	user, err := udao.GetUserById(uint(video.AuthorId))
	if err != nil {
		username = "未知"
	} else {
		username = user.UserName
	}
	//获取视频标签
	c, err := cdao.GetCateById(int64(video.CategoryId))
	if err != nil {
		categoryName = "其他"
	} else {
		categoryName = c.CategoryName
	}
	data = &types.GetFavResp{
		VideoId:         int64(video.ID),
		AuthorId:        video.AuthorId,
		CreateTime:      video.CreatedAt.Unix(),
		AuthorName:      username,
		PlayCount:       0,
		CoverURL:        video.CoverURL,
		PlayURL:         video.PlayURL,
		FavoriteCount:   video.FavoriteCount,
		CollectionCount: video.CollectionCount,
		Title:           video.Title,
		Category:        categoryName,
	}

	return
}
