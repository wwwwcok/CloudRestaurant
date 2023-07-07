package controller

import (
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GoodsController struct {
}

func (gs *GoodsController) Router(engine *gin.Engine) {
	engine.GET("/api/goods", gs.GetGoods)
}

func (gs *GoodsController) GetGoods(context *gin.Context) {
	shopId := context.Query("shop_id")
	if shopId == "" {
		tool.Failed(context, "请求解析错误")
	}

	//开始获取Goods
	id, _ := strconv.Atoi(shopId)
	goodsService := service.GoodsService{}
	goods := goodsService.GetGoodsByshopID(int64(id))

	if len(*goods) == 0 {
		tool.Failed(context, "没有查询到商品信息")
		return
	}
	//查到的商品信息
	tool.Success(context, goods)

}
