package handler_person

import (
	"encoding/json"
	"errors"
	"fmt"
	"lab_device_management_api/dal"
	handler_verify "lab_device_management_api/handler/verify"
	"lab_device_management_api/model"
	"lab_device_management_api/proto/person/person"
	"lab_device_management_api/rpc"
	"log"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//已经登录了
	if v, exist := c.Get("is_login"); v == true && exist {
		return
	}
	var (
		resp  *person.LoginResponse
		login model.Login
		err   error
	)
	c.ShouldBindJSON(&login)
	fmt.Println(login)

	//1.参数校验
	err = checkLoginParams(c, &login)
	if err != nil {
		log.Printf("checkLoginParams is failed,err:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	//2.请求rpc服务，查库判断是否有此人
	resp, err = rpc.Login(c, nil, &login)
	if err != nil {
		log.Printf("Login is failed,err:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	//3.登录成功，先查询一下此人token是否还在，还没有过期
	tokenExist, err := dal.Get(nil, login.MultiId)
	if err != nil {
		fmt.Printf("token is not exist,warn: %s", err)
	}
	//如果存在，直接返回
	if tt1, ok := tokenExist.(string); ok {
		if len(tt1) > 0 {
			dataMap := make(map[string]string)
			dataMap["token"] = tt1
			data, _ := json.Marshal(dataMap)
			c.JSON(http.StatusOK, gin.H{
				"message": resp.Message,
				"data":    string(data),
			})
			return
		}
	} else if tt2, ok := tokenExist.(int); ok {
		if tt2 > 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": resp.Message,
				"data":    tt2,
			})
			return
		}
	}

	//4.登录成功，返回前端token，用于后续跳转页面身份认证。放入中间件中取校验token
	//这里采用放入redis中，过期时间为三天,multi_id为key，以multi_id生成uuid作为token
	token := uuid.NewV4()
	fmt.Printf("Successfully parsed: %s", token)

	err = dal.SetEx(nil, login.MultiId, token, 60*60*24*3)
	if err != nil {
		log.Printf("[Login] SetEx is failed,err:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	dataMap := make(map[string]string)
	dataMap["token"] = token.String()
	data, _ := json.Marshal(dataMap)

	c.JSON(http.StatusOK, gin.H{
		"message": resp.Message,
		"data":    string(data),
	})
}

func checkLoginParams(c *gin.Context, login *model.Login) error {
	if s := strings.TrimSpace(login.MultiId); s == "" {
		return errors.New("学号或职工号不能为空")
	}
	if s := strings.TrimSpace(login.Password); s == "" {
		return errors.New("密码不能为空")
	}
	if login.Code == 0 {
		return errors.New("code不合法")
	}
	//校验验证码
	verify, err := handler_verify.Verify(login)
	if !verify || err != nil {
		return errors.New("验证码校验失败")
	}
	return nil
}
