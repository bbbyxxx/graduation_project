package routers_person

import (
	handler_person "lab_device_management_api/handler/person"

	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.POST("/regist_person", handler_person.RegistPerson)

	e.POST("/login", handler_person.Login)

	e.POST("/update_person", handler_person.UpdatePerson)

}
