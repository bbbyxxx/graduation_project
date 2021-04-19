package main

import (
	"context"
	"fmt"
	"lab_device_management_person/proto/test/wuqing"
	"net"
	"net/http"

	"golang.org/x/net/trace"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	//Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

type helloService struct{}

var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *wuqing.HelloRequest) (out *wuqing.HelloResponse, err error) {
	resp := new(wuqing.HelloResponse)
	//resp.Message = fmt.Sprintf("Hello %s. token info:appid=%s,appkey=%s", in.Name, appid, appkey)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	return resp, nil
}

func auth(ctx context.Context) error {
	// 解析metadata中的信息并验证
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "无 Token 认证信息")
	}

	var (
		appid  string
		appkey string
	)

	if val, ok := md["appid"]; ok {
		appid = val[0]
	}

	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}

	if appid != "101010" || appkey != "i am key" {
		return grpc.Errorf(codes.Unauthenticated, "Token 认证信息无效:appid=%s,appkey=%s", appid, appkey)
	}
	return nil
}

//interceptor 拦截器
func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	err := auth(ctx)
	if err != nil {
		return nil, err
	}
	//继续处理请求
	return handler(ctx, req)
}

func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	go http.ListenAndServe(":50051", nil)
	fmt.Println("Trace listen on 50051")
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen:%v", err)
	}

	var opts []grpc.ServerOption

	//TlS认证
	creds, err := credentials.NewServerTLSFromFile("keys/server.pem", "keys/server.key")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials:%v", err)
	}

	opts = append(opts, grpc.Creds(creds))
	// 注册interceptor
	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	//实例化 grpc server
	s := grpc.NewServer(opts...)

	go startTrace()

	//注册HelloService
	wuqing.RegisterHelloServer(s, HelloService)

	fmt.Println("listen on ", Address, " with tls")

	s.Serve(listen)
}
