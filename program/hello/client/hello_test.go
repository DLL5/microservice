package main

import (
	pb "awesomeProject/proto/hello"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"testing"
)

const testAddress = "127.0.0.1:50052"

func Test_Hello(t *testing.T) {
	// 连接
	//conn, err := grpc.Dial(Address, grpc.WithInsecure())
	inse := insecure.NewCredentials()
	conn, err := grpc.Dial(testAddress, grpc.WithTransportCredentials(inse))
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer func() { _ = conn.Close() }()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)
	fmt.Println(res.Message)
	if err != nil {
		grpclog.Fatalln(err)
	}

	grpclog.Infoln(res.Message)
	if res.Message != "Hello gRPC." {
		t.Errorf("gRPC 调用 hello服务的SayHello方法错误！")
	}

}

//
//func Benchmark_Hello(b *testing.B) {
//	inse := insecure.NewCredentials()
//	conn, err := grpc.Dial(testAddress, grpc.WithTransportCredentials(inse))
//	if err != nil {
//		grpclog.Fatalln(err)
//	}
//	defer func() { _ = conn.Close() }()
//
//	// 调用方法
//	req := &pb.HelloRequest{Name: "gRPC"}
//	for i := 0; i < b.N; i++ {
//		// 初始化客户端
//		c := pb.NewHelloClient(conn)
//		res, err := c.SayHello(context.Background(), req)
//		//fmt.Println(res.Message)
//		if err != nil {
//			grpclog.Fatalln(err)
//		}
//
//		grpclog.Infoln(res.Message)
//	}
//}

func helloTotal() {
	inse := insecure.NewCredentials()
	conn, err := grpc.Dial(testAddress, grpc.WithTransportCredentials(inse))
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer func() { _ = conn.Close() }()

	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}

	// 初始化客户端
	c := pb.NewHelloClient(conn)
	res, err := c.SayHello(context.Background(), req)
	//fmt.Println(res.Message)
	if err != nil {
		grpclog.Fatalln(err)
	}

	grpclog.Infoln(res.Message)

}

//func Benchmark_HelloTotal(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		helloTotal()
//	}
//}
