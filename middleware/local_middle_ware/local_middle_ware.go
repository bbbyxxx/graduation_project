package local_middle_ware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"lab_device_management_api/dal"
	"lab_device_management_api/model"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token model.Token
		)

		t := time.Now()
		fmt.Println("中间件开始执行了")

		//获取请求体内容，因为请求体读一次后面读取就是空的，很坑！！！
		buf := make([]byte, 1024)
		body, _ := c.Request.Body.Read(buf)
		fmt.Printf("body is %v\n", body)
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(buf[:]))

		//得到请求参数
		fmt.Println(string(buf))
		_ = c.ShouldBindJSON(&token)
		fmt.Println(token)

		//回写请求体
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(buf[:]))

		tokenRedis, err := dal.Get(nil, token.MultiId)
		if err != nil {
			log.Println("[MiddleWare] Get is nil,warn:%v", err)
		}
		if tt, ok := tokenRedis.(string); ok {
			if len(tt) > 0 {
				if strings.Compare(tt, token.Token) == 0 {
					c.Set("is_login", true)
					c.Set("multi_id", token.MultiId)
				}
			} else {
				c.Set("is_login", false)
			}
		}

		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("TIME:", t2)

		c.Next()
	}
}
