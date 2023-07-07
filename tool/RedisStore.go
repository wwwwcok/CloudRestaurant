package tool

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
)

type RedisStore struct {
	client *redis.Client
}

var RediStore RedisStore

func InitRedisStore() *RedisStore {
	config := GetConfig().RedisConfig
	client := redis.NewClient(&redis.Options{ //在option结构体里填入配置参数
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.Db,
	})
	err := client.Set("testt", "测试redis连接", 0).Err()
	if err != nil {
		fmt.Println("云餐厅项目测试redis连接失败:", err)
	}

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	RediStore := RedisStore{client: client}
	base64Captcha.SetCustomStore(&RediStore)
	return &RediStore
}

//必须要有的这两个方法

func (s *RedisStore) Set(id string, value string) {
	err := s.client.Set(id, value, time.Minute*10).Err()
	if err != nil {
		log.Println(err)
	}
}

// customizeRdsStore implementing Get method of  Store interface
func (s *RedisStore) Get(id string, clear bool) (value string) {
	val, err := s.client.Get(id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	if clear {
		err := s.client.Del(id).Err()
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	return val
}
