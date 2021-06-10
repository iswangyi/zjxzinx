package znet

import (
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



func NewConnection(conn *net.Conn,connID uint32, callbackApi ziface.HandlerFunc )  *Connection{
	c := &Connection{
		Conn:        conn,
		ConnID:      connID,
		HandlerFunc: callbackApi,
		isClosed:    false,
		ExitChan:    make(chan bool,1),

	}
	return c
}

func (c *Connection)Start(){

}

func (c *Connection)Stop()  {

}
func (c *Connection)GetTCPConnection() *net.TCPConn {
	return c.Conn

}

func (c *Connection)GetConnectionID() uint32  {
	return c.ConnID
}
func (c *Connection)GetRemoteAddr() net.TCPConn   {
	c.Conn.RemoteAddr()

}
func (c *Connection)Send() error {
	return nil
}
