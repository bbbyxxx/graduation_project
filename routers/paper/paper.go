package routers_paper

import (
	handler_paper "lab_device_management_api/handler/paper"

	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.POST("/add_paper", handler_paper.AddPaper)

	e.POST("/update_paper", handler_paper.UpdatePaper)

	e.GET("/mget_paper", handler_paper.MGetPaper)
}
