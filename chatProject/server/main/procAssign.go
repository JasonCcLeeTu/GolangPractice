package  main


import(
	"net"
	_"encoding/json"
	"go_Code/chatProject/common/message"
	"fmt"
	_"encoding/binary"
	"io"
	"go_Code/chatProject/server/utils"
	"go_Code/chatProject/server/process"
)


type ProcAssign struct{

	Con net.Conn
}


func (this *ProcAssign) ServerProcessMes( ms *message.Message) (err error){
             
	switch ms.Type{

	case message.LoginMesType:
		  prDetail:= &process.ProcDetail{
			  Con:this.Con,
			  
		  }
		err=  prDetail.ServerProcessLogin(ms)
		if err!=nil{
			fmt.Println("ServerProcessLogin() error",err)
			return
		}
	case message.RegisterMesType :
		prDetail:= &process.ProcDetail{
			Con:this.Con,
		}
	  err=  prDetail.ServerProcessRegister(ms)
	  if err!=nil{
		  fmt.Println("ServerProcessRegister() error",err)
		  return
	  }
		
    case message.SendMesType:
          smsP:=&process.SmsProcess{}
		 err= smsP.SendContentToAll(ms)
		 if err!=nil{
			 fmt.Println("sendContentToAll fail:",err)
			 return
		 }
	default:
	   fmt.Println("消息類型不存在，無法處理")
	}
	return

}



func (this *ProcAssign) Process(){

	// var loginMes message.Message
	defer this.Con.Close()

	 for{
		
	   rwPkage :=&utils.RWPkg{
		   Con:this.Con,
	   }
	  mes,err:=  rwPkage.ReadPkg()
	  
	 	 if err!=nil{
			if err==io.EOF{
				fmt.Println("客戶端退出，服務端也退出")
				return
			}else{
				fmt.Println("readPkg erro:",err)
				return
			}
		 }

	 	 err = this.ServerProcessMes(&mes)
			if err!=nil{
				return
			}
    
	}
}