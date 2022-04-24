package process

import(
	"go_Code/chatProject/common/message"  
	"fmt"   
	"go_Code/chatProject/client/model"
	
)


var(
	 
	onlineUsrList  map[int]*message.User = make(map[int]*message.User,10)
	currentUser  model.CurUser
)

     
func showOnlineList(){
           fmt.Println("當前用戶列表:")
	 for id,_:=range onlineUsrList{
		 fmt.Printf("目前%d在線\n",id)
	 }
}


func  updateOnlineList(reLsit *message.RenewList) (err error){

		
	
	      _,ok:= onlineUsrList[reLsit.UsrId]
		  if !ok{
			   usr:=&message.User{
				UsrId:reLsit.UsrId,
			   }
			   onlineUsrList[reLsit.UsrId]=usr
		  }

            onlineUsrList[reLsit.UsrId].UsrStatus=message.OnlineStatus

		//    user.Status=message.OnlineStatus
		//    onlineUsrList[reLsit.UsrId]=user

		showOnlineList()
         return

}

