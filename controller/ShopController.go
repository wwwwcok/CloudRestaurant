package controller

import (
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ShopController struct {
}

func (sc *ShopController) Router(engine *gin.Engine) {
	engine.GET("/api/shops", sc.GetShopList)
}

func (sc *ShopController) GetShopList(context *gin.Context) {
	longitude := context.Query("longitude")
	latitude := context.Query("latitude")

	fmt.Println("---------------------------------------longitude,latitude:", longitude, latitude)
	if longitude == "" || latitude == "" {
		//如果为空就返回，默认经纬度
		longitude = "416"
		latitude = "10"
	}

	fmt.Println("---------------------------------------longitude,latitude:", longitude, latitude)

	shopService := service.ShopService{}
	shops := shopService.ShopList(longitude, latitude)
	if len(shops) == 0 {
		tool.Failed(context, "没有获取到商户信息")
		return
	}

	for _, shop := range shops {
		TheshopServies := shopService.GetService(shop.Id)
		if len(*TheshopServies) == 0 {
			shop.Supports = nil
		} else {
			shop.Supports = *TheshopServies
		}
	}

	tool.Success(context, shops)

}
