package process

import(
   "fmt"
   "os"
   "net"
   "go_Code/chatProject/client/utils"
   "go_Code/chatProject/common/message"
   "encoding/json"
)


func  Monitor(con net.Conn){

	     rwPackage:= utils.RWPkg{
			 Con:con,
		 }

		 for{
			 fmt.Println("持續監聽伺服器端")
			 
			 mes,err:= rwPackage.ReadPkg()
			 
			 if err!=nil{
				 fmt.Println("Monitor ReadPkg fail:",err)
				 return
			 }
			
			 switch mes.Type{

				case message.ReNewListMesType :
					
					var  reList message.RenewList
					
					   json.Unmarshal([]byte(mes.Data),&reList)
					   
					 

                     updateOnlineList(&reList)
				case message.SendMesType:
					var  sms smsProcess
					err :=sms.AcceptMessage(&mes)
					if err!=nil{
						fmt.Println("AcceptMessage fail:",err)
						
					}
					  
				default:
					fmt.Println("無匹配狀態")
				}
			
			 }
       
		
}


func ShowMenu(){

    fmt.Println("--------恭喜xxx登入成功----------")
	fmt.Println("--------1. 顯示在線用戶列表----------")
	fmt.Println("--------2. 發送消息----------")
	fmt.Println("--------3. 信息列表----------")
	fmt.Println("--------4. 退出系統----------")
	fmt.Println("請選擇(1-4):")
    var key int
	var content string
	sms1 := &smsProcess{}
	fmt.Scanf("%d\n",&key)
	
	switch key {

	case 1:
		showOnlineList()
		
		
	case 2:
	
		fmt.Println("輸入想發送的消息:")
		fmt.Scanf("%s\n",&content)
	   sms1.sendMessage(content)
		
		
		
		
		
	case 3:
		
		
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系統")
		os.Exit(0)


	}
		




}