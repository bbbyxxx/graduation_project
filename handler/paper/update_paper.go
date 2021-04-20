package handler_paper

import (
	"lab_device_management_api/model"
	"lab_device_management_api/rpc"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdatePaper(c *gin.Context) {
	var (
		modelPaper model.Paper
		err        error
	)
	c.ShouldBindJSON(&modelPaper)
	log.Printf("modelDevice is %v\n", modelPaper)
	multiIdInterface, exist := c.Get("multi_id")
	multiId, _ := multiIdInterface.(string)
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "未登录",
			"data":    "",
		})
		return
	}

	//1.检查参数
	err = checkParams(&modelPaper)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    "",
		})
	}

	//2.调用rpc服务，更新论文信息
	resp, err := rpc.UpdatePaper(c, nil, &modelPaper, multiId)
	if err != nil {
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
	return
}
