package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"code/chartRoom/common/message"
)

type Transfer struct {
	Conn net.Conn
	Buf [8096]byte
}




func (this *Transfer) ReadPkg() (msg message.Message, err error) {
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		return msg, err
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if int(pkgLen) != n || err != nil {
		return msg, err
	}
	err = json.Unmarshal(this.Buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("json.Unmarshal error")
		return msg, err
	}
	return msg, nil
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))

	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	_, err = this.Conn.Write(this.Buf[:4])
	if err != nil {
		fmt.Println("conn.Write faild :", err)
		return
	}

	//发送消息本身
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write faild :", err)
		return
	}
	return
}