package main

import (
	"context"
	"fmt"
	"net"

	pb "example.com/learn-grpc/hello-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// 绝对地址
	creds, err1 := credentials.NewServerTLSFromFile(
		"/root/workspace/learn-grpc/key/test.pem",
		"/root/workspace/learn-grpc/key/test.key",
	)

	if err1 != nil {
		fmt.Printf("证书错误：%v", err1)
		return
	}

	// 开启端口
	listen, _ := net.Listen("tcp", ":9090")
	// 创建grpc服务
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	// 在grpc服务端中注册我们自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})

	// 启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Println(err)
		return
	}
}

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		ResponseMsg: "hello," + req.RequestName,
	}, nil
}
