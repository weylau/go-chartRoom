package processes

import (
	"fmt"
	"os"
	"net"
	"code/chartRoom/client/utils"
)

func ShowMenu() {
	fmt.Println("----恭喜成功登陆----")
	fmt.Println("1、显示在线用户列表")
	fmt.Println("2、发送消息")
	fmt.Println("3、信息列表")
	fmt.Println("4、退出系统")
	fmt.Println("请输入1-4")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("您输入的信息有误")

	}


}

func serverProcessMsg(conn net.Conn) (err error) {
	transfer := utils.Transfer{
		Conn:conn,
	}
	for {
		msg, err := transfer.ReadPkg()
		if err != nil {
			fmt.Println("err", err)
			return err
		}
		fmt.Println("读取客户端数据：", msg)
	}
}