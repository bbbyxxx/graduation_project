package handler_device

import (
	"lab_device_management_api/model"
	"lab_device_management_api/rpc"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateDevice(c *gin.Context) {
	var (
		modelDevice model.Device
		err         error
	)
	c.ShouldBindJSON(&modelDevice)

	log.Printf("modelDevice is %v", modelDevice)
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
	err = checkParams(modelDevice)
	if err != nil {
		log.Printf("[UpdateDevice] is failed,err:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	modelDevice.DeviceNumberModelId = modelDevice.DeviceNumberId + "-" + modelDevice.DeviceModelId
	//2.调用 RPC 服务更新设备信息
	resp, err := rpc.UpdateDevice(c, nil, &modelDevice, multiId)
	if err != nil {
		log.Printf("[UpdateDevice] call rpc err:%v\n", err)
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
