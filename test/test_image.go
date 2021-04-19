package main

import (
	"fmt"
	services "lab_device_management_api/service"
	"net/http"

	"github.com/qiniu/api.v7/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["images"]
		// 遍历所有图片
		for _, file := range files {
			// 逐个存
			if err := c.SaveUploadedFile(file, "images/"+file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
			//上传图片到七牛云
			key, _, err := services.UploadImageToQiNiuYun("images/"+file.Filename, file.Filename)
			if err != nil {
				return
			}
			domain := "https://image.test.com"
			publicAccessURL := storage.MakePublicURL(domain, key)
			fmt.Println(publicAccessURL)
		}

		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})
	//默认端口号是8080
	r.Run(":8000")
}
