package main

import (
	"fmt"
	routers_device "lab_device_management_api/routers/device"
	routers_paper "lab_device_management_api/routers/paper"
	routers_person "lab_device_management_api/routers/person"
	routers_verify "lab_device_management_api/routers/verify"
)

//主函数：加载路由及初始化
func main() {
	Include(routers_person.Routers, routers_verify.Routers, routers_device.Routers, routers_paper.Routers)
	r := Init()
	if err := r.Run("127.0.0.1:8001"); err != nil {
		fmt.Println("startup service failed,err:%v\n", err)
	}
}
