package main

import (
	"fmt"
	"lab_device_management_api/dal"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	fmt.Println("redis conn success")

	dal.SetEx(nil, "abc", "sada", 60)
	dal.Get(nil, "abc")

	defer c.Close()
}
