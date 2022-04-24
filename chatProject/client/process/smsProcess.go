package process

import(
	"go_Code/chatProject/common/message"
	"encoding/json"
	"go_Code/chatProject/client/utils"
	"fmt"
)

type  smsProcess struct{

	
}


func  (this *smsProcess) sendMessage(content string) (err error){

        var  mes   message.Message
		mes.Type= message.SendMesType

		var sms message.SendMes
		sms.Content = content
		sms.UsrId = currentUser.UsrId

		 data,err := json.Marshal(sms)
		 if err!=nil{
			 fmt.Println("Marshal fail:",err)
			 return
		 }
		 mes.Data=string(data)

	    data,err =json.Marshal(mes)
		if err!=nil{
			fmt.Println("Marshal fail:",err)
			return
		}

		 rwP := utils.RWPkg{
			 Con:currentUser.Con,
		 }

		 err = rwP.WritePkg(data)

		 if err!=nil{
			 fmt.Println("WritePkg fail:",err)
			 return
		 }
		 return


}

func  (this *smsProcess) AcceptMessage(ms *message.Message) (err error){
         var  sms message.SendMes
       err=  json.Unmarshal([]byte(ms.Data),&sms)
	   if err!=nil{
		fmt.Println("Unmarshal fail:",err)
		return
	   }
	    fmt.Printf("編號:%d 說: %s\n",sms.UsrId,sms.Content)  
		return
	  

}
