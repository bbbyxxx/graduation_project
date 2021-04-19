package main

import (
	"fmt"
	handler_device "lab_device_management_device/handler/device"
	handler_paper "lab_device_management_device/handler/paper"
	"lab_device_management_device/proto/device/device"
	"lab_device_management_device/proto/paper/paper"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/keepalive"
)

const (
	//grpc 服务地址
	Address = "127.0.0.1:8002"
)

var (
	handlerDevice = handler_device.Device{}
	handlerPaper  = handler_paper.Paper{}
)

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen : %v", err)
		return
	}

	//实例化 grpc server
	s := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute, //这个连接最大的空闲时间，超过就释放，解决proxy等到网络问题（不通知grpc的client和server）
	}))

	//注册服务方法
	device.RegisterDeviceServer(s, &handlerDevice)
	paper.RegisterPaperServer(s, &handlerPaper)

	fmt.Println("Listen on " + Address)
	s.Serve(listen)
}
