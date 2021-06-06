package main

import "zjxzinx/zinx/znet"

//基于zinx开发的服务器程序
func main() {
	//创建一个server 句柄，使用zinx 的api
	s := znet.NewServer("[zinxv0.1]")
	s.Serve()
}
