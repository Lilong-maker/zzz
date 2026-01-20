package service

import (
	"context"
	"zzz/service/basic/config"
	__ "zzz/service/basic/proto"
	"zzz/service/handler/model"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedUserServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) GoodsAdd(_ context.Context, in *__.GoodsAddReq) (*__.GoodsAddResp, error) {

	var good model.Goods
	err := good.Findgood(config.DB, in.Name)
	if err != nil {
		return &__.GoodsAddResp{
			Code: 400,
			Msg:  "商品不存在",
		}, nil
	}
	err = good.GoodAdd(config.DB)
	if err != nil {
		return &__.GoodsAddResp{
			Code: 400,
			Msg:  "商品添加失败",
		}, nil
	}
	return &__.GoodsAddResp{
		Code: 200,
		Msg:  "商品添加成功",
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) GoodsUpdate(_ context.Context, in *__.GoodsUpdateReq) (*__.GoodsUpdateResp, error) {

	var good model.Goods
	err := good.FindGoods(config.DB, in.Id)
	if err != nil {
		return &__.GoodsUpdateResp{
			Code: 400,
			Msg:  "商品不存在",
		}, nil
	}
	ModGoods := model.Goods{
		Name: in.Name,
		Num:  int(in.Num),
	}
	err = ModGoods.UpdateGoods(config.DB, in.Id)
	if err != nil {
		return &__.GoodsUpdateResp{
			Code: 400,
			Msg:  "修改失败",
		}, nil
	}
	return &__.GoodsUpdateResp{
		Code: 200,
		Msg:  "修改成功",
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) GoodsList(_ context.Context, in *__.GoodsListReq) (*__.GoodsListResp, error) {

	var good model.Goods
	err := good.FindGoods(config.DB, in.Id)
	if err != nil {
		return &__.GoodsListResp{
			Code: 400,
			Msg:  "商品不存在",
		}, nil
	}
	return &__.GoodsListResp{
		Name: good.Name,
		Num:  int32(good.Num),
		Code: 200,
		Msg:  "列表查询成功",
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) GoodsDelete(_ context.Context, in *__.GoodsDeleteReq) (*__.GoodsDeleteResp, error) {

	var good model.Goods
	err := good.FindGoods(config.DB, in.Id)
	if err != nil {
		return &__.GoodsDeleteResp{
			Code: 400,
			Msg:  "商品不存在",
		}, nil
	}
	err = good.DeleteGoods(config.DB, in.Id)
	if err != nil {
		return &__.GoodsDeleteResp{
			Code: 400,
			Msg:  "删除失败",
		}, nil
	}
	return &__.GoodsDeleteResp{
		Code: 200,
		Msg:  "删除成功",
	}, nil
}
