package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
)

type GoodsDao struct {
	*tool.Orm
}

func (gd *GoodsDao) GetGoods(shopId int64) *[]model.Goods {
	var goods []model.Goods
	err := gd.Where("shop_id = ?", shopId).Find(&goods)
	if err != nil {
		return nil
	}
	return &goods
}
