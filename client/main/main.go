package main

import (
	"fmt"
	"os"
	"code/chartRoom/client/processes"
)

var (
	userId int
	userName string
	userPwd string
)

func main() {
	//聊天系统的客户端
	var key int
	var loop bool = true
	for loop {
		fmt.Println("======欢迎进入多人聊天系统======")
		fmt.Println("======1、登录系统======")
		fmt.Println("======2、注册账号======")
		fmt.Println("======3、退出系统======")
		fmt.Println("请输入上面的操作（1-3）")
		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("登录系统")
			loop = false
		case 2:
			fmt.Println("注册账号")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
			//Exit让当前程序以给出的状态码code退出。一般来说，状态码0表示成功，非0表示出错。
			//程序会立刻终止，defer的函数不会被执行。
			os.Exit(0)
		default:
			fmt.Println("您输入有误，请重新输入！！！")

		}
	}

	if key == 1 {//登录操作
		fmt.Println("请输入用户名！")
		fmt.Scanln(&userName)
		fmt.Println("请输入密码！")
		fmt.Scanln(&userPwd)
		userProcess := &processes.UserProcess{}
		err := userProcess.Login(userName, userPwd)
		if err != nil {
			fmt.Println("用户名或密码错误")
		} else {
			fmt.Println("登录成功")
		}
	} else if key == 2 {
		fmt.Println("请输入用ID！")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户名！")
		fmt.Scanln(&userName)
		fmt.Println("请输入密码！")
		fmt.Scanln(&userPwd)
		userProcess := &processes.UserProcess{}
		err := userProcess.Register(userId, userName, userPwd)
		if err != nil {
			fmt.Println("用户名或密码错误")
		} else {
			fmt.Println("注册成功")
		}
	}



}
