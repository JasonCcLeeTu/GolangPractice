package model

import(
	"net"
	"go_Code/chatProject/common/message"
)

type CurUser struct {

	Con net.Conn
	 message.User 
	
}