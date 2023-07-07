package tool

import (
	"CloudRestaurant/model"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Orm struct {
	*xorm.Engine
}

var DbEngine *Orm

func OrmEngine(cfg *Config) (*Orm, error) {

	database := cfg.Database //没必要通过传参的方式，直接GetConfig().Database更简单
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		return nil, err
	}

	engine.ShowSQL(database.ShowSql)

	err = engine.Sync2(&model.SmsCode{}, &model.Member{},
		&model.FoodCategory{}, &model.Shop{}, &model.Service{}, &model.ShopService{}, &model.Goods{})
	if err != nil {
		return nil, err
	}

	orm := &Orm{}
	orm.Engine = engine

	DbEngine = orm

	//插入初始化店铺数据
	InitShopData()
	//插入初始化商品数据
	InitGoodsData()
	return orm, nil
}

func InitShopData() {
	shops := []model.Shop{{Id: 1, Name: "嘉禾一品(温水都城)", Address: "北京温都水城F1", Longitude: 116.36868, Latitude: 40.222, Phone: "13444850035", Status: 1, RecentOrderNum: 106, RatingCount: 961, Rating: 4.7, PromotionInfo: "欢迎光临", ImagePath: "", MinimumOrderAmount: 0, DeliveryFee: 5, OpeningHours: "8:30/20:30", IsNew: true, IsPremiun: true},
		{Id: 2, Name: "皇城停车场", Address: "北京停车场F1", Longitude: 117.4348, Latitude: 42.222, Phone: "1543760035", Status: 1, RecentOrderNum: 204, RatingCount: 461, Rating: 4.4, PromotionInfo: "嘿嘿嘿", ImagePath: "", MinimumOrderAmount: 0, DeliveryFee: 2, OpeningHours: "6:30/17:30", IsNew: true, IsPremiun: true},
		{Id: 3, Name: "按摩店", Address: "北京按摩店F2", Longitude: 416.36868, Latitude: 10.622, Phone: "1758782033", Status: 1, RecentOrderNum: 112, RatingCount: 523, Rating: 3.4, PromotionInfo: "顾客你好", ImagePath: "", MinimumOrderAmount: 0, DeliveryFee: 12, OpeningHours: "7:00/18:00", IsNew: true, IsPremiun: true},
	}
	//事务 开启新session
	session := DbEngine.NewSession()
	defer session.Close()
	//事务开始，执行(回滚),提交事务
	err := session.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, shop := range shops {
		_, err := session.Insert(&shop)
		if err != nil {
			fmt.Println(err.Error())
			session.Rollback()
			return
		}
	}
	err = session.Commit()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func InitGoodsData() {
	goods := []model.Goods{
		{Name: "小小鲜肉包", Description: "滑蛋牛肉粥(1份)+小小鲜肉包(4只)", SellCount: 14, Price: 25, OldPrice: 29, ShopId: 1},
		{Name: "港式牛杂", Description: "港式车仔面(1份)", SellCount: 8, Price: 8, OldPrice: 18, ShopId: 1},
		{Name: "人肉叉烧包", Description: "人肉叉烧包(6只)", SellCount: 20, Price: 50, OldPrice: 100.5, ShopId: 479},
	}
	//事务 开启新session
	session := DbEngine.NewSession()
	defer session.Close()
	//事务开始，执行(回滚),提交事务
	err := session.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, shop := range goods {
		_, err := session.Insert(&shop)
		if err != nil {
			fmt.Println(err.Error())
			session.Rollback()
			return
		}
	}
	err = session.Commit()
	if err != nil {
		fmt.Println(err.Error())
	}
}
