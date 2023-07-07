package main

import (
	"CloudRestaurant/controller"
	"CloudRestaurant/tool"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//编写中间件函数
func Test_middleware() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.Set("测试中间件", "中间件被触发")
	}
}

func main() {

	ginServer := gin.Default()
	cfg := tool.ParseConfig("./config/app.json")

	//实例化orm及数据库
	_, err := tool.OrmEngine(cfg)
	if err != nil {
		return
	}
	//实例化redis
	tool.InitRedisStore()
	//设置跨域访问
	ginServer.Use(Cors())
	//设置session
	tool.InitSession(ginServer)

	RegisterRouter(ginServer)

	fmt.Println("__________________________________--saas", cfg.AppHost+":"+cfg.AppPort)
	err = ginServer.Run(cfg.AppHost + ":" + cfg.AppPort)
	if err != nil {
		fmt.Println("错误——————————————————————", err)
	}
}

func RegisterRouter(ginServer *gin.Engine) {
	Hellocontrol := controller.HelloController{}
	Hellocontrol.Router(ginServer)
	Membercontrol := controller.MemberController{}
	Membercontrol.Router(ginServer)
	FoodCategoryController := controller.FoodCategorycController{}
	FoodCategoryController.Router(ginServer)
	ShopController := controller.ShopController{}
	ShopController.Router(ginServer)
}

//跨域访问响应
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headerKeys []string

		for key, _ := range context.Request.Header {
			headerKeys = append(headerKeys, key)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("Access-Control-Allow-Origin, Access-Control-Allow-header, %s", headerStr)
		} else {
			headerStr = "Access-Control-Allow-Origin, Access-Control-Allow-Header"
		}
		if origin != "" {
			context.Header("Access-Control-Allow-Origin", "*")                                  //任意都能访问
			context.Header("Access-Control-Allow-Method", "POST,OPTIONS,GET,UPDATE,PUT,DELETE") //访问方法
			context.Header("Access-Control-Allow-Headers", "Authoriztion,Content-Length,X-CSRF-Token,Token,session")
			context.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin, Access-Control-Allow-Header")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false")
			context.Set("content-type", "application/json") //返回值类型设置为json

		}

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "OPTIONS request!!!!!")
		}

	}
}
