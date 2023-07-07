package model

//商家
type Shop struct {
	Id   int64  `xorm:"pk autoincr" json:"id"`
	Name string `xorm:"varchar(32)" json:"name"`
	//宣传
	PromotionInfo string `xorm:"varchar(100)" json:"promotionInfo"`
	Address       string `xorm:"varchar(100)" json:"address"`
	Phone         string `xorm:"varchar(32)" json:"phone"`
	//店铺状态
	Status int `xorm:"int" json:"status"`
	//经度
	Longitude float64 `xorm:"float" json:"longitude"`
	//维度
	Latitude float64 `xorm:"float" json:"latitude"`
	//商家店铺
	ImagePath string `xorm:"varchar(255)" json:"imagePath"`

	IsNew bool `xorm:"bool" json:"is_new"`

	IsPremiun bool `xorm:"bool" json:"is_premiun"`

	//商铺评分
	Rating float32 `xorm:"float" json:"rating"`
	//评分总数
	RatingCount int64 `xorm:"int" json:"rating_count"`
	//当前订单总数
	RecentOrderNum int64 `xorm:"int" json:"recent_order_num"`
	//配送起送价
	MinimumOrderAmount int32 `xorm:"int" json:"minimum_order_amount"`
	//配送费
	DeliveryFee int32 `xorm:"int" json:"delivery_fee"`
	//营业时间
	OpeningHours string `xorm:"varchar(20)" json:"opening_hours"`
	Supports     []Service
}
