package routers_verify

import (
	handler_verify "lab_device_management_api/handler/verify"

	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	//1.获取验证码
	//http://localhost:xxxx/captcha
	e.GET("/captcha", handler_verify.GetCaptcha)

	//2.获取验证码图片
	//http://localhost:xxxx/captcha/gHEIwh7nWreTFb53MkVk.png
	e.GET("/captcha/:captchaId", handler_verify.GetCaptchaPicture)

}
