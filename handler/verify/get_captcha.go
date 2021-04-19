package handler_verify

import (
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

//获取验证码
func GetCaptcha(c *gin.Context) {
	length := captcha.DefaultLen
	captchaId := captcha.NewLen(length)
	var captcha CaptchaResponse
	captcha.CaptchaId = captchaId
	captcha.ImageUrl = "/captcha/" + captchaId + ".png"
	c.JSON(http.StatusOK, captcha)
}
