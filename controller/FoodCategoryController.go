package controller

import (
	"CloudRestaurant/service"
	"CloudRestaurant/tool"

	"github.com/gin-gonic/gin"
)

type FoodCategorycController struct {
}

func (fcc *FoodCategorycController) Router(engine *gin.Engine) {
	engine.GET("/api/food_category", fcc.foodCategory)
}

func (fcc *FoodCategorycController) foodCategory(context *gin.Context) {
	//调用service获得食品种类信息
	foodCategoryService := &service.FoodCategoryService{}
	categories, err := foodCategoryService.Categories()
	if err != nil {
		tool.Failed(context, "食品种类获取失败")
		return
	}
	tool.Success(context, categories)

	//转换格式：主要是路径
	for _, category := range categories {
		if category.ImageUrl != "" {
			category.ImageUrl = "http://192.168.0.90:80" + "/" + category.ImageUrl
		}
	}
	tool.Success(context, categories)
}
