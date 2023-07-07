package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
	"strconv"
)

type ShopService struct {
}

func (ShopService *ShopService) GetService(shopId int64) *[]model.Service {
	dao := dao.ShopDao{Orm: tool.DbEngine}
	shopService := dao.QueryServiceByShopID(shopId)
	return shopService
}

func (ShopService *ShopService) ShopList(long, lat string) []model.Shop {
	longitude, err := strconv.ParseFloat(long, 64)
	if err != nil {
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return nil
	}

	shopDao := dao.NewShopDao()
	return shopDao.QueryShops(longitude, latitude)
}
