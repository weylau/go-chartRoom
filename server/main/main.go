package main

import (
	"net"
	"fmt"
	"os"

	"code/chartRoom/server/processes"
	"code/chartRoom/server/model"
	"time"
)




func main() {
	initPool("192.168.1.105:6379", 16, 0, 300*time.Second)
	ud := &model.UserDao{}
	model.MyUserDao = ud.InitUserDao(pool)

	fmt.Println("服务端开启8889端口")
	listener, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("tcp端口监听失败", err)
		os.Exit(0)
	}

	for {
		fmt.Println("等待客户端来连接服务端")
		conn, err := listener.Accept()
		if err!=nil {
			fmt.Println("listener.Accept error:", err)
		}
		processor := &processes.Processor{
			Conn:conn,
		}
		go processor.Process()
	}

}