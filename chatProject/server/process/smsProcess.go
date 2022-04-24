package process

import(
     "go_Code/chatProject/common/message"
	 "encoding/json"
	 "go_Code/chatProject/server/utils"
	 "fmt"
)


type SmsProcess struct {


}


func (this *SmsProcess) SendContentToAll(ms *message.Message)(err error){
     var   sendMs  message.SendMes 
     err= json.Unmarshal([]byte(ms.Data),&sendMs)
	 if err!=nil{
		fmt.Println("Unmarshal fail:",err)
		 return
	 }

     data,err:= json.Marshal(ms)
	 if err!=nil{
		 fmt.Println("Marshal fail:",err)
		 return
	 }
	 
	 var   rwP utils.RWPkg
	 for  _,info:= range  UsrOlineList.usrOnList{
             if  info.Id ==  sendMs.UsrId{
				 continue
			 }
			 rwP.Con=info.Con
			err= rwP.WritePkg(data)
			if err!=nil{
				fmt.Println("WritePkg fail:",err)
				return
			}
			 

	 }
	 return


}

