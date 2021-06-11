package znet

import (
	"fmt"
	"net"
	"zjxzinx/zinx/ziface"
)

type Connection struct {
	Conn *net.TCPConn
	ConnID uint32
	isClosed bool
	HandlerFunc ziface.HandlerFunc
	//告知当前链接状态的chan
	ExitChan chan bool
}

// NewConnection 创建链接对象
func NewConnection(conn *net.TCPConn,connID uint32, callbackApi ziface.HandlerFunc )  *Connection{
	c := &Connection{
		Conn:        conn,
		ConnID:      connID,
		HandlerFunc: callbackApi,
		isClosed:    false,
		ExitChan:    make(chan bool,1),

	}

	return c
}
//开始
func (c *Connection)Start(){

}

// Stop 关闭链接是否资源
func (c *Connection)Stop()  {
	fmt.Println("conn closed,connid:",c.ConnID)
	if c.isClosed ==true {
		return
	}
	c.isClosed = true
	//关闭socket链接
	c.Conn.Close()
	//释放资源
	close(c.ExitChan)
}

// GetTCPConnection 获取TCP链接对象
func (c *Connection)GetTCPConnection() *net.TCPConn {
	return c.Conn

}

// GetConnectionID 获取链接ID
func (c *Connection)GetConnectionID() uint32  {
	return c.ConnID
}

// GetRemoteAddr 获取远程addr
func (c *Connection)GetRemoteAddr() net.Addr   {
	return c.Conn.RemoteAddr()

}

// Send 发送数据
func (c *Connection)Send() error {
	return nil
}
