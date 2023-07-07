package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
)

type GoodsService struct {
}

func (gs *GoodsService) GetGoodsByshopID(shopId int64) *[]model.Goods {
	dao := dao.GoodsDao{Orm: tool.DbEngine}
	goods := dao.GetGoods(shopId)
	return goods
}
