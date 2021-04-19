package handler_verify

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//获取验证码图片
func GetCaptchaPicture(c *gin.Context) {
	captchaId := c.Param("captchaId")
	fmt.Println("GetCaptchaPng : " + captchaId)
	ServeHTTP(c.Writer, c.Request)
}
