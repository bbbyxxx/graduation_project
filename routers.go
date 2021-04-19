package main

import (
	"io"
	"lab_device_management_api/middleware/local_middle_ware"
	"os"

	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

//注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

//初始化
func Init() *gin.Engine {
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	r := gin.New()
	//注册全局中间件
	r.Use(local_middle_ware.MiddleWare())
	for _, opt := range options {
		opt(r)
	}
	return r
}
