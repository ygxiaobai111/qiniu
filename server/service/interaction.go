package service

import (
	"context"
	"sync"
	"www.github.com/ygxiaobai111/qiniu/server/types"
)

var InterSrvIns *InterSrv
var InterSrvOnce sync.Once

type InterSrv struct {
}

// GetUserSrv 返回userSrv对象
func GetInterSrv() *InterSrv {
	InterSrvOnce.Do(func() {
		InterSrvIns = &InterSrv{}
	})
	return InterSrvIns
}

func (s *InterSrv) GetFavlist(ctx context.Context, req *types.GetFavlistReq) (resp interface{}, err error) {
	return

}
func (s *InterSrv) GetFavorite(ctx context.Context, req *types.GetFavoriteReq) (resp interface{}, err error) {

	return
}
func (s *InterSrv) GetComment(ctx context.Context, req *types.GetCommentReq) (resp interface{}, err error) {
	return

}
func (s *InterSrv) GetBarrage(ctx context.Context, req *types.GetBarrageReq) (resp interface{}, err error) {

	return
}
func (s *InterSrv) FavlistCreate(ctx context.Context, req *types.FavlisCreatetReq) (resp interface{}, err error) {

	return
}
func (s *InterSrv) FavlistAdd(ctx context.Context, req *types.FavlistAddReq) (resp interface{}, err error) {
	return

}
func (s *InterSrv) FavlistDel(ctx context.Context, req *types.FavlistDelReq) (resp interface{}, err error) {
	return

}
func (s *InterSrv) DelFavlist(ctx context.Context, req *types.DelFavlistReq) (resp interface{}, err error) {
	return

}
func (s *InterSrv) CommentCreate(ctx context.Context, req *types.CommentCreateReq) (resp interface{}, err error) {
	return

}
func (s *InterSrv) Favorite(ctx context.Context, req *types.FavoriteReq) (resp interface{}, err error) {
	return

}
func (s *InterSrv) Barrage(ctx context.Context, req *types.BarrageReq) (resp interface{}, err error) {
	return

}
