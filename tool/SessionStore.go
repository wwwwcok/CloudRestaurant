package tool

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitSession(Engine *gin.Engine) {
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(""))
	if err != nil {
		fmt.Println(err.Error())
	}
	Engine.Use(sessions.Sessions("mysession", store))
}

//Set
//Get

func Setsess(context *gin.Context, key interface{}, value interface{}) error {
	//获取当前session
	session := sessions.Default(context)

	session.Set(key, value)
	return session.Save()

}

func Getsess(context *gin.Context, key interface{}) interface{} {
	session := sessions.Default(context)
	return session.Get(key)
}
