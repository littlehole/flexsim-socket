package main

import (
	"fmt"
	"io"
	"net"
)

const input int = 300   //3位
const packNeed int = 10 // 2位
const redNeed int = 30
const greenNeed int = 30
const blueNeed int = 30

func process(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte(fmt.Sprintf("input%dpack%dred%dblue%dgreen%d", input, packNeed, redNeed, blueNeed, greenNeed)))
	for {
		//创建一个切片
		buf := make([]byte, 1024)
		//1.等待客户端通过conn发送信息学
		//2.如果客户端没有wirte[发送]，那么协程就阻塞在这里
		//fmt.Printf("服务器在等待客户端%s 发送信息\n",conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("the connetction is closed")
				conn.Close()
			} else {
				fmt.Printf("Read Error: %s\n", err)
			}
			return
		}
		//3.显示客户端发送的内容到服务器的终端
		fmt.Printf("客户端:%s 发送信息:%s。\n", conn.RemoteAddr().String(), string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close()
	//循环等待客户端来连接
	for {
		fmt.Println("等待客户端来连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err=", err)
		} else {
			fmt.Printf("Accept suc con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}
