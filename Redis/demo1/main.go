package main

import (
	"fmt"
	"github.com/go-redis/redis"
)
var r *redis.Client
func initRedis(){
	r = redis.NewClient(&redis.Options{
		Addr:               "localhost:6379",
		Password:           "12345",
		DB:                 0,

	})

	str,err := r.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str,",θΏζ₯ζε!")
}

func main() {
	initRedis()

	status := r.Set("test","hello",0)
	fmt.Println(status)
}
