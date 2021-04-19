package handler_person

import (
	"fmt"
	"lab_device_management_api/model"
	"lab_device_management_api/rpc"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdatePerson(c *gin.Context) {
	var person model.Person
	err := c.BindJSON(&person)
	if err != nil {
		log.Printf("[UpdatePerson] Bind is failed,err:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "参数绑定错误",
			"data":    "",
		})
		return
	}

	//判断是否登录状态
	if v, exist := c.Get("is_login"); v != true || exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "请登录",
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

	//2.更新人物信息
	resp, err := rpc.UpdatePerson(c, nil, &person)
	if err != nil {
		log.Printf("[UpdatePerson] UpdatePerson is failed,err:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp.Message,
		"data":    "",
	})
}
