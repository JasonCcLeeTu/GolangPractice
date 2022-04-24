package main

import(
	"net"
	_"encoding/json"
	_"go_Code/chatProject/common/message"
	"fmt"
	_"encoding/binary"
	_"io"
	"go_Code/chatProject/server/model"
	"time"
)




	  
       



func init(){
	CreatePool(0,16,time.Second*300,"0.0.0.0:6379")
    model.MyUsrDao = model.NewUsrDao(pool)

}
          

			 


func main(){
       fmt.Println("port:8888 監聽中")
	 listener,err:=net.Listen("tcp","0.0.0.0:8888")

	 if err !=nil{
		 fmt.Println("Listen error:",err)
	 }

	 defer listener.Close()

     for{
		fmt.Println("等待客戶端連接")
		con,err:=listener.Accept()
		if err!=nil{
			fmt.Println("listener error:",err)
		}
       
		 prAssign := &ProcAssign{
			 Con:con,
		 }
        go prAssign.Process()


	 }


}
