package processes

import (
	"net"
	"fmt"
	"code/chartRoom/common/message"
	"code/chartRoom/server/utils"
	"io"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) ServerProcessMes(msg *message.Message) (err error) {
	switch msg.Type {
	case message.LoginMsgType:
		userProcess := &UserProcess{
			Conn:this.Conn,
		}
		err1 := userProcess.ServiceProcessLogin(msg)
		if err1 != nil {
			err = err1
			fmt.Println("login errro", err)
			return
		}
	case message.RegisterMsgType:
		userProcess := &UserProcess{
			Conn:this.Conn,
		}
		err1 := userProcess.ServiceProcessRegister(msg)
		if err1 != nil {
			err = err1
			fmt.Println("register errro", err)
			return
		}
	default:
		fmt.Println("消息类型不存在，无法处理")

	}
	return
}


func (this *Processor) Process() (err error)	 {
	defer this.Conn.Close()
	for {
		transfer := utils.Transfer{
			Conn:this.Conn,
		}
		msg, err := transfer.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出")
				return err
			}
			fmt.Println("客户端退出……", err)
			return err
		}
		fmt.Println("读取客户端数据：", msg)
		this.ServerProcessMes(&msg)
	}

}
