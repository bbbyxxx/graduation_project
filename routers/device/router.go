package routers_device

import (
	handler_device "lab_device_management_api/handler/device"

	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.POST("/add_device", handler_device.AddDevice)

	e.POST("/update_device", handler_device.UpdateDevice)

}
