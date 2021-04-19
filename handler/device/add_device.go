package handler_device

import (
	"errors"
	"lab_device_management_api/model"
	"lab_device_management_api/rpc"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddDevice(c *gin.Context) {
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

	//获取上传的图片
	//form, err := c.MultipartForm()
	//if err != nil {
	//	c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
	//}
	//// 获取所有图片
	//files, _ := form.File["images"]
	//// 遍历所有图片
	//for _, file := range files {
	//	// 逐个存
	//	file.Filename = modelDevice.DeviceNumberId + "-" + modelDevice.DeviceNumberModelId + file.Filename
	//	if err := c.SaveUploadedFile(file, "images/"+file.Filename); err != nil {
	//		c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
	//		return
	//	}
	//}
	//后面支持：图片通过url地址回传给前端，前端可直接打开

	//1.检查参数
	err = checkParams(modelDevice)
	if err != nil {
		log.Printf("[AddDevice] is failed,err:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	//2.调用Rpc服务,完成设备添加
	resp, err := rpc.AddDevice(c, nil, &modelDevice, multiId)
	if err != nil {
		log.Printf("[AddDevice] call rpc err:%v\n", err)
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

func checkParams(device model.Device) (err error) {
	if device.IsDeleted == 0 {
		if s := strings.TrimSpace(device.DeviceNumberId); s == "" {
			return errors.New("device_number_id 不能为空")
		}
		if s := strings.TrimSpace(device.DeviceModelId); s == "" {
			return errors.New("device_model_id 不能为空")
		}
		if s := strings.TrimSpace(device.LaboratoryFloorId); s == "" {
			return errors.New("laboratory_floor_id 不能为空")
		}
		if device.LaboratoryId == 0 {
			return errors.New("laboratory_id 不能为空")
		}
		if device.DeviceStatus == 0 {
			return errors.New("device_status 不能为空")
		}
		if s := strings.TrimSpace(device.DeviceUseDesc); s == "" {
			return errors.New("device_use_desc 不能为空")
		}
	}
	return
}
