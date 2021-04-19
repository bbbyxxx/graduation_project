package handler_verify

import (
	"errors"
	"fmt"
	"lab_device_management_api/model"
	"log"
	"strconv"
	"strings"

	"github.com/dchest/captcha"
)

//验证-验证码
func Verify(login *model.Login) (bool, error) {
	captchaId := login.CaptchaId
	captchaId = strings.ReplaceAll(captchaId, ".png", "")
	code := login.Code
	fmt.Println(captchaId, " ", code)
	codeString := strconv.Itoa(int(code))
	fmt.Println(captchaId, " ", codeString)
	if captchaId == "" || codeString == "" {
		return false, errors.New("参数错误")
	}
	if captcha.VerifyString(captchaId, codeString) {
		return true, nil
	} else {
		log.Println("[Verify] VerifyString 验证校验码失败")
		return false, nil
	}
}
