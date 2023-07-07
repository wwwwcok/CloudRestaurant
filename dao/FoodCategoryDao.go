package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
)

type FoodCategoryDao struct {
	*tool.Orm
}

func NewFoodCategoryDao() *FoodCategoryDao {
	return &FoodCategoryDao{tool.DbEngine}
}

func (fcd *FoodCategoryDao) QueryCategories() ([]model.FoodCategory, error) {
	var categories []model.FoodCategory

	if err := fcd.Engine.Find(&categories); err != nil {
		return nil, err
	}

	return categories, nil
}
