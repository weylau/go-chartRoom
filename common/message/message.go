package message

const(//消息类型定义
	LoginMsgType  =  "LoginMsg" //登录
	RegisterMsgType  =  "RegisterMsg" //注册
	LoginResMsgType  =  "LoginResMsg" //登录结果返回
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type User struct {
	UserId int `json:"user_id"`
	UserPwd string `json:"user_pwd"`
	UserName string `json:"user_name"`
}

type LoginMsg struct {
	UserName string `json:"user_name"`
	UserPwd string `json:"user_pwd"`
}

type RegisterMsg struct {
	User User `json:"user"`
}

type ResponseResMsg struct {
	Code int `json:"code"` //200成功 其他为错误
	Error string `json:"error"`
}
