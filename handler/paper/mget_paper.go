package handler_paper

import (
	"encoding/json"
	"lab_device_management_api/model"
	"lab_device_management_api/rpc"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MGetPaper(c *gin.Context) {
	var (
		modelPersonDevicePaper model.PersonDevicePaper
		err                    error
	)
	c.ShouldBindJSON(&modelPersonDevicePaper)
	modelPersonDevicePaper.PaperNumber = c.Query("paper_number")
	modelPersonDevicePaper.DeviceNumberModelId = c.Query("device_number_model_id")
	log.Printf("modelPersonDevicePaper is %v\n", modelPersonDevicePaper)
	_, exist := c.Get("multi_id")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "未登录",
			"data":    "",
		})
		return
	}

	//不用校验参数是否全为空，因为 multi_id 不能为空
	//1.调用RPC服务
	resp, err := rpc.MGetPaper(c, nil, &modelPersonDevicePaper)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	content, _ := json.Marshal(resp.Paper)
	c.JSON(http.StatusOK, gin.H{
		"message": resp.Message,
		"data":    string(content),
	})
	return
}
