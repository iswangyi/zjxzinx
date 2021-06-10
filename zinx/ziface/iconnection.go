package ziface

import "net"

// IConnection 定义连接模块的抽象层
type IConnection interface {
	// Start 启动链接
	Start()
	// Stop 停止连接，结束当前的链接工作

	Stop()

	// GetTCPConnection 获取当前链接的scoket
	GetTCPConnection() *net.TCPConn

	// GetConnectionID 获取当前的连接ID
	GetConnectionID() uint32

	// GetRemoteAddr 获取远程客户端的TCP状态
	GetRemoteAddr() net.Addr
	// Send 发送数据
	Send(data []byte)  error
}



//定义一个处理链接的方法

type HandlerFunc func(*net.TCPConn,[]byte,int) error