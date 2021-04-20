package handler_device

import (
	"encoding/json"
	"lab_device_management_api/rpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MGetDeviceByDNMIS(c *gin.Context) {
	var (
		err error
	)
	multiIdInterface, exist := c.Get("multi_id")
	multiId, _ := multiIdInterface.(string)
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "未登录",
			"data":    "",
		})
		return
	}
	deviceNumberModelId := c.Query("device_number_model_id")
	if deviceNumberModelId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "device_number_model_id不能为空",
			"data":    "",
		})
		return
	}

	//1.调用rpc服务获取数据
	resp, err := rpc.MQueryDeviceByDNMIS(c, nil, []string{deviceNumberModelId}, multiId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    "",
		})
	}

	content, _ := json.Marshal(resp.Device)
	c.JSON(http.StatusOK, gin.H{
		"message": resp.Message,
		"data":    string(content),
	})
}
