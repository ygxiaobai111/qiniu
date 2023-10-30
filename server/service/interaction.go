package service

import (
	"context"
	"errors"
	"log"
	"sync"
	"www.github.com/ygxiaobai111/qiniu/server/pkg/e"
	dao2 "www.github.com/ygxiaobai111/qiniu/server/repository/db/dao"
	"www.github.com/ygxiaobai111/qiniu/server/repository/db/model"
	"www.github.com/ygxiaobai111/qiniu/server/types"
)

var InterSrvIns *InterSrv
var InterSrvOnce sync.Once

type InterSrv struct {
}

func GetInterSrv() *InterSrv {
	InterSrvOnce.Do(func() {
		InterSrvIns = &InterSrv{}
	})
	return InterSrvIns
}

func (s *InterSrv) GetFavlist(ctx context.Context, req *types.GetFavlistReq) (resp interface{}, err error) {

	cdao := dao2.NewCollectionDao(ctx)
	udao := dao2.NewUserDao(ctx)
	user, err := udao.GetUserById(uint(req.UserId))
	if err != nil {
		return nil, errors.New(e.GetMsg(e.ERROR))
	}
	//当只携带userid则返回所有收藏夹
	if req.FavlistId == 0 {
		var collections []*model.Collection
		collections, err = cdao.GetCollections(req.UserId)
		if err != nil {
			return nil, err
		}
		var getFavlistResp []types.GetFavlistResp
		for _, collection := range collections {
			var videos []*types.GetFavResp
			if err != nil {
				return
			}
			videos = BuildVideos(ctx, collection.Videos)

			r := types.GetFavlistResp{
				UserName:       user.UserName,
				CollectionName: collection.Name,
				CreateTime:     collection.CreatedAt.Unix(),
				Favlist:        videos,
				Total:          int64(len(videos)),
			}

			getFavlistResp = append(getFavlistResp, r)
		}
		resp = types.DataList{
			Item:  getFavlistResp,
			Total: uint(len(getFavlistResp)),
		}
		return

	}
	//返回单个数据
	collection, err := cdao.GetCollection(req.FavlistId)
	if err != nil {
		return nil, err
	}
	var videos []*types.GetFavResp
	if err != nil {
		return
	}
	videos = BuildVideos(ctx, collection.Videos)

	resp = types.GetFavlistResp{
		UserName:       user.UserName,
		CollectionName: collection.Name,
		CreateTime:     collection.CreatedAt.Unix(),
		Favlist:        videos,
		Total:          int64(len(videos)),
	}
	return

}
func (s *InterSrv) GetFavorite(ctx context.Context, req *types.GetFavoriteReq) (resp interface{}, err error) {
	fdao := dao2.NewFavDao(ctx)

	videos := fdao.ListFav(ctx, req.UserId)
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
func (s *InterSrv) GetComment(ctx context.Context, req *types.GetCommentReq) (resp interface{}, err error) {
	commentdao := dao2.NewCommentDao(ctx)
	userdao := dao2.NewUserDao(ctx)
	comments, err := commentdao.CommentList(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	var r []types.GetCommentResp
	for _, comment := range comments {
		author, _ := userdao.GetUserById(uint(comment.UserId))
		data := types.GetCommentResp{
			UserName:   author.UserName,
			Avatar:     author.Avatar,
			Content:    comment.Content,
			CreateTime: comment.CreatedAt.Unix(),
		}
		r = append(r, data)
	}
	resp = types.DataList{
		Item:  r,
		Total: uint(len(r)),
	}

	return

}
func (s *InterSrv) GetBarrage(ctx context.Context, req *types.GetBarrageReq) (resp interface{}, err error) {
	ddao := dao2.NewDanDao(ctx)

	dans, err := ddao.GetDanByVid(req.VideoId)
	if err != nil {
		return
	}
	var r []types.GetBarrageResp
	for _, dan := range dans {

		data := types.GetBarrageResp{
			Content:   dan.Content,
			Color:     dan.Color,
			Timestamp: dan.Timestamp,
		}
		r = append(r, data)
	}
	resp = types.DataList{
		Item:  r,
		Total: uint(len(r)),
	}
	return

}

func (s *InterSrv) FavlistCreate(ctx context.Context, req *types.FavlisCreatetReq, userId uint) (resp interface{}, err error) {
	dao := dao2.NewCollectionDao(ctx)
	collection := &model.Collection{

		Name:      req.FavlistName,
		IsPrivate: req.Type,
		UserID:    userId,
		Videos:    []*model.Video{},
	}
	err = dao.Create(collection)
	return
}
func (s *InterSrv) FavlistAdd(ctx context.Context, req *types.FavlistAddReq) (resp interface{}, err error) {
	cdao := dao2.NewCollectionDao(ctx)

	c, err := cdao.GetCollection(req.FavlistId)
	if err != nil {
		log.Println("err1:", err)
		return
	}
	err = cdao.AddVideo(c, uint(req.VideoId))
	return

}
func (s *InterSrv) FavlistDel(ctx context.Context, req *types.FavlistDelReq) (resp interface{}, err error) {
	cdao := dao2.NewCollectionDao(ctx)

	c, err := cdao.GetCollection(req.FavlistId)
	if err != nil {
		return
	}
	err = cdao.RemoveVideo(c, uint(req.VideoId))
	return

}
func (s *InterSrv) DelFavlist(ctx context.Context, req *types.DelFavlistReq, userId int64) (resp interface{}, err error) {
	cdao := dao2.NewCollectionDao(ctx)

	err = cdao.DelCollection(req.FavlistId, userId)
	return
	return

}
func (s *InterSrv) CommentCreate(ctx context.Context, req *types.CommentCreateReq, userId int64) (resp interface{}, err error) {
	dao := dao2.NewCommentDao(ctx)
	comment := model.Comment{
		VideoID: req.VideoId,
		UserId:  userId,
		Content: req.Content,
	}
	b, err := dao.SaveComment(ctx, comment)
	if err != nil {
		return
	}
	if b != true {
		err = errors.New(e.GetMsg(e.ERROR))
	}
	return

}
func (s *InterSrv) Favorite(ctx context.Context, req *types.FavoriteReq, userId int64) (resp interface{}, err error) {
	dao := dao2.NewFavDao(ctx)
	switch req.Type {
	case 1:
		err = dao.CreateFav(ctx, req.VideoId, userId) //关注
	case 2:
		err = dao.DeleteFav(ctx, req.VideoId, userId) //取关
	default:
		err = errors.New(e.GetMsg(e.InvalidParams))
	}
	return

}
func (s *InterSrv) Barrage(ctx context.Context, req *types.BarrageReq, userId uint64) (resp interface{}, err error) {
	dao := dao2.NewDanDao(ctx)
	dan := model.Danmaku{
		VideoID:   req.VideoID,
		UserID:    uint(userId),
		Content:   req.Content,
		Color:     req.Color,
		Timestamp: req.Timestamp,
	}
	err = dao.Create(dan)
	return

}
