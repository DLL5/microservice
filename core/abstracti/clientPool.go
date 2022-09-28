package abstracti

import (
	"google.golang.org/grpc"
	"sync"
)

// GrpcConnPoolI 微服务架构，使用连接池的场景相对较多，封装一下pool连接池的使用方法
// 使用说明：
// 调用func NewGrpcConnPool(cp GrpcConnPoolI) *sync.Pool 函数即可
// example:
//
//	var Pool = abstracti.NewGrpcConnPool(&helloPool{}) // 接收返回的连接池对象
//	type helloPool struct{} // 定义一个空结构体，实现该接口GrpcConnPoolI方法
//	func (p *helloPool) NewConn() *grpc.ClientConn {}
type GrpcConnPoolI interface {
	NewConnection() *grpc.ClientConn
}

// NewGrpcConnPool 生成grpc连接池对象
func NewGrpcConnPool(cp GrpcConnPoolI) *sync.Pool {
	return &sync.Pool{New: func() any { return cp.NewConnection() }}
}
