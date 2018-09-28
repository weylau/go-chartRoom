package processes

import (
	"net"
	"code/chartRoom/common/message"
	"encoding/json"
	"fmt"
	"code/chartRoom/server/utils"
	"code/chartRoom/server/model"
)


type UserProcess struct {
	Conn net.Conn

}


func (this *UserProcess) ServiceProcessLogin(msg *message.Message) (err error) {
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("json.Unmarshal fail ", err)
		return
	}
	var resMsg message.Message
	var loginResMsg message.ResponseResMsg
	user, err := model.MyUserDao.Login(loginMsg.UserName, loginMsg.UserPwd)


	if err != nil {
		loginResMsg.Code = 500
		loginResMsg.Error = "用户名或密码错误！"
		fmt.Println(loginResMsg.Error)
	} else {
		fmt.Println("登录成功", loginMsg.UserName,user)
		loginResMsg.Code = 200
	}
	resMsg.Type = message.LoginResMsgType
	data, err := json.Marshal(loginResMsg)
	resMsg.Data = string(data)
	if err != nil {
		fmt.Println("json.Marshal error", err)
		return
	}
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("message格式化错误！", err)
		return
	}
	transfer := &utils.Transfer{
		Conn : this.Conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("service writePkg error", err)
		return
	}
	return
}



func (this *UserProcess) ServiceProcessRegister(msg *message.Message) (err error) {
	var registerMsg message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data), &registerMsg)
	if err != nil {
		fmt.Println("json.Unmarshal fail ", err)
		return err
	}
	var resMsg message.Message
	var responseResMsg message.ResponseResMsg
	user, err := model.MyUserDao.Register(registerMsg.User.UserId, registerMsg.User.UserName, registerMsg.User.UserPwd)


	if err != nil {
		responseResMsg.Code = 500
		responseResMsg.Error = err.Error()
		fmt.Println(responseResMsg.Error)
	} else {
		fmt.Println("注册成功", user)
		responseResMsg.Code = 200
	}
	resMsg.Type = message.RegisterMsgType
	data, err := json.Marshal(responseResMsg)
	resMsg.Data = string(data)
	if err != nil {
		fmt.Println("json.Marshal error", err)
		return err
	}
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("message格式化错误！", err)
		return err
	}
	transfer := &utils.Transfer{
		Conn : this.Conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("service writePkg error", err)
		return err
	}
	return
}