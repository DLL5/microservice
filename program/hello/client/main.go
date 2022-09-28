package main

import (
	"awesomeProject/core/abstracti"
	pb "awesomeProject/proto/hello" // 引入proto包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	// 连接
	//conn, err := grpc.Dial(Address, grpc.WithInsecure())
	//inse := insecure.NewCredentials()
	//conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(inse))
	//if err != nil {
	//	grpclog.Fatalln(err)
	//}

	//conn := NewConnection()
	//defer func() {
	//	_ = conn.Close()
	//}()
	//
	//// 初始化客户端
	//c := pb.NewHelloClient(conn)
	//
	//// 调用方法
	//req := &pb.HelloRequest{Name: "gRPC"}
	//res, err := c.SayHello(context.Background(), req)
	//fmt.Println(res.Message)
	//if err != nil {
	//	grpclog.Fatalln(err)
	//}
	//
	//grpclog.Infoln(res.Message)
	SampleCallHello("ldl")

}

var Pool = abstracti.NewGrpcConnPool(&helloPool{})

type helloPool struct{}

// Pool hello服务client客户端的连接池
//var Pool = &helloPool{}.newClientPool(&helloPool{})

//// newHelloClientPool 生成连接池
//func NewClientPool(cp *abstracti.GrpcConnPoolI) *sync.Pool {
//	return &sync.Pool{New: func() any { return cp.NewConnection() }}
//}

// NewConnection 新建一个默认的客户端连接，实现接口abstracti.GrpcConnPoolI
func (p *helloPool) NewConnection() *grpc.ClientConn {
	return p.newDefaultConn()
}

// NewDefaultConn 新建一个默认的客户端连接
func (p *helloPool) newDefaultConn() *grpc.ClientConn {
	inse := insecure.NewCredentials()
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(inse))
	if err != nil {
		grpclog.Error(err)
	}
	return conn
}

// SampleCallHello 一个简单调用Hello服务的程序
func SampleCallHello(name string) (*pb.HelloResponse, string) {
	// 获取一个连接
	conn := Pool.Get().(*grpc.ClientConn)
	defer Pool.Put(conn)
	//conn := NewConnection()
	//defer conn.Close()
	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: name}
	res, err := c.SayHello(context.Background(), req)
	//fmt.Println(res.Message)
	if err != nil {
		grpclog.Fatalln(err)
	}

	grpclog.Infoln(res.Message, res)
	//time.Sleep(20 * time.Millisecond)
	return res, res.Message
}
