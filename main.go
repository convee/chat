package main

import (
	"chat/cache"
	"chat/conf"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	conf.InitConfig("config.toml")
	cache.InitRedis()
	name, _ := redis.String(cache.GetRedis().Do("get", "name"))
	fmt.Println(name)

}
