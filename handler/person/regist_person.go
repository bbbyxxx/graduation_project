package handler_person

import (
	"errors"
	"fmt"
	"lab_device_management_api/model"
	"lab_device_management_api/rpc"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegistPerson(c *gin.Context) {
	var person model.Person
	err := c.BindJSON(&person)
	if err != nil {
		log.Printf("[RegistPerson] Bind is failed,err:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "参数绑定错误",
			"data":    "",
		})
		return
	}

	fmt.Println(person)

	//1.检查参数
	err = checkParams(&person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	//2.请求rpc方法,将信息入库
	resp, err := rpc.RegisterPerson(c, nil, &person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	//3.将响应返回给前端
	c.JSON(http.StatusOK, gin.H{
		"message": resp.Message,
		"data":    "",
	})
}

func checkParams(person *model.Person) (err error) {
	if person == nil {
		return errors.New("参数有误")
	}
	if person.IsDeleted == 0 {
		if person.MultiId = strings.TrimSpace(person.MultiId); person.MultiId == "" {
			return errors.New("学号或职工号不能为空")
		}
		pattern := "^[0-9]*$" //反斜杠要转义
		result, _ := regexp.MatchString(pattern, person.MultiId)
		if !result {
			return errors.New("学号或职工号非法")
		}
		if person.Password = strings.TrimSpace(person.Password); person.Password == "" {
			return errors.New("密码不能为空")
		}
		//学生年级班级不为空
		if person.Indentity == 2 {
			if person.Grade <= 0 || person.Grade >= 5 {
				return errors.New("年级范围不符合规定")
			}
			if person.Class <= 0 || person.Class >= 18 {
				return errors.New("班级范围不符合规定")
			}
			if person.Major = strings.TrimSpace(person.Major); person.Major == "" {
				return errors.New("专业不能为空")
			}
		}
		if person.Phone = strings.TrimSpace(person.Phone); person.Phone == "" {
			return errors.New("手机号不能为空")
		}
		pattern1 := "0\\d{2,3}-[1-9]\\d{6,7}"
		result1, _ := regexp.MatchString(pattern1, person.Phone)
		pattern = "1[3-9]\\d{9}"
		result, _ = regexp.MatchString(pattern, person.Phone)
		if !result && !result1 {
			return errors.New("手机号长度不合法")
		}
		if person.Indentity <= 0 || person.Indentity >= 4 {
			return errors.New("身份不合法")
		}
	}
	return nil
}
