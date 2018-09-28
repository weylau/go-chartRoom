package processes

import (
	"net"
	"fmt"
	"encoding/json"
	"code/chartRoom/common/message"
	"code/chartRoom/client/utils"
	"errors"
)

type UserProcess struct {

}

func (this *UserProcess) Register(user_id int, user_name string, user_pwd string) (err error) {
	//连接服务端
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("连接服务端失败", err)
		return err
	}
	defer conn.Close()

	//消息组装
	msgType := message.RegisterMsgType
	user := message.User{
		UserId:user_id,
		UserName:user_name,
		UserPwd:user_pwd,
	}
	registerMsg := message.RegisterMsg{
		User:user,
	}

	data, err := json.Marshal(registerMsg)
	if err != nil {
		fmt.Println("registerMsg格式化错误！", err)
		return err
	}

	var msg = message.Message{
		Type:msgType,
		Data:string(data),
	}

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("message格式化错误！", err)
		return err
	}

	transfer := &utils.Transfer{
		Conn:conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("客户端写入数据失败")
		return err
	}

	msg, err = transfer.ReadPkg()
	if err != nil {
		fmt.Println("客户端读取数据失败")
		return err
	}

	var registerResMsg message.ResponseResMsg
	err = json.Unmarshal([]byte(msg.Data), &registerResMsg)
	if err != nil {
		fmt.Println("json.Unmarshal error ", err)
		return err
	}
	if registerResMsg.Code == 200 {
		fmt.Println("注册成功")
	} else {
		err = errors.New(registerResMsg.Error)
		fmt.Println(registerResMsg.Error)
	}
	return err
}


func (this *UserProcess) Login(userName string, userPwd string) (err error) {
	//fmt.Printf("用户名：%v,密码：%v", userName, userPwd)

	//连接服务端
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("连接服务端失败", err)
		return err
	}
	defer conn.Close()

	//消息组装
	msgType := message.LoginMsgType
	loginMsg := message.LoginMsg{
		UserName:userName,
		UserPwd:userPwd,
	}

	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("loginMsg格式化错误！", err)
		return err
	}

	var msg = message.Message{
		Type:msgType,
		Data:string(data),
	}

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("message格式化错误！", err)
		return err
	}

	transfer := &utils.Transfer{
		Conn:conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("客户端写入数据失败")
		return err
	}

	msg, err = transfer.ReadPkg()
	if err != nil {
		fmt.Println("客户端读取数据失败")
		return err
	}

	var loginResMsg message.ResponseResMsg
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)
	if err != nil {
		fmt.Println("json.Unmarshal error ", err)
		return err
	}
	if loginResMsg.Code == 200 {
	} else {
		err = errors.New(loginResMsg.Error)
		fmt.Println(loginResMsg.Error)
	}
	return err
}



