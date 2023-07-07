package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
)

type ShopDao struct {
	*tool.Orm
}

func NewShopDao() *ShopDao {
	return &ShopDao{tool.DbEngine}
}

const DEFAULT_RANGE = 5

func (shopDao *ShopDao) QueryShops(longitude, latitude float64) []model.Shop {
	var shops []model.Shop
	err := shopDao.Engine.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ?", longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE).Find(&shops)
	if err != nil {
		return nil
	}
	return shops
}

//根据商户id查询它有哪些服务
func (shopDao *ShopDao) QueryServiceByShopID(shopId int64) *[]model.Service {
	var services []model.Service
	shopDao.Table("service").Join("INNER", "shop_service", "service.id=shop_service.shopid").Find(&services)
	return &services
}
