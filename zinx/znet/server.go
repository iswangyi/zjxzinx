package znet

import (
	"fmt"
	"net"
)

// Server iServer接口的实现
type Server struct {
	//服务器名称
	Name string
	//服务器绑定ip版本
	Ipversion string
	//服务器监听IP
	Ip string
	//服务器监听Port
	port int
}

// Start Iserver的接口方法实现
func (s *Server) Start() {
	//1、获取一个TCP的addr
	fmt.Println("[ZinxServer] lisetn at ", s.Ipversion, s.Ip, s.port)
	go func() {
		addr, err := net.ResolveTCPAddr(s.Ipversion, fmt.Sprintf("%s:%d", s.Ip, s.port))
		if err != nil {
			fmt.Println("ResolveIPaddr err:", err)
			return
		}

		//2、监听服务器地址

		listener, err := net.ListenTCP(s.Ipversion, addr)
		if err != nil {
			fmt.Println("ListenTCP error", err)
			return
		}
		fmt.Println("start zinx server succ", s.Name)

		//3、阻塞的等待客户端连接，处理客户端连接业务

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("listener AcceptTCP err", err)
				continue
			}
			//已经建立连接，处理业务，做一个512字节长度回显业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("conn Read buf err", err)
						continue
					}
					//回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("Write buf err", err)
						continue
					}

				}
			}()

		}
	}()

}
func (s *Server) Stop() {
	//释放服务器资源

}
func (s *Server) Serve() {
	s.Start()
	//TODU 做一些业务处理,start 提供监听功能，serve处理业务
	//阻塞行为
	select {}
}

// NewServer 初始化Server的方法
func NewServer(name string) *Server {
	s := &Server{
		Name:      name,
		Ipversion: "tcp4",
		Ip:        "0.0.0.0",
		port:      8080,
	}
	return s
}
