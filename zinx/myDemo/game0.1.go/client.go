package main

import (
	"fmt"
	"net"
	"time"
)

//模拟客户端

func main() {
	//1、连接原创服务器，得到conn
	conn, err := net.Dial("tcp4", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("connection err ,check ip:port", err)
	} else {
		fmt.Println("连接成功")
	}
	for {
		//2、链接调用write，写数据
		_, err := conn.Write([]byte("服务器收到请回答"))
		if err != nil {
			fmt.Println("conn Write error, send  ", err)
		}
		bufServer := make([]byte, 512)
		cnt, err := conn.Read(bufServer)
		if err != nil {
			fmt.Println("read err", err)
		} else {
			fmt.Printf("服务器返回内容：%q\n", bufServer[:cnt])
		}
		time.Sleep(1 * time.Second)
	}

}
