package process

import(
	"net"
	"encoding/json"
	"go_Code/chatProject/common/message"
	"fmt"
	_"encoding/binary"
	_"io"
	"go_Code/chatProject/server/utils"
	"go_Code/chatProject/server/model"
)

type ProcDetail struct {

	Con net.Conn
	Id int
	
}



func (this *ProcDetail)  UpdateOnlineList(id int) (err error){

         for index,val := range  UsrOlineList.usrOnList {
			    
			if index==id {
				continue
			}

		 err =	val.SendOnlineToOther(id)
				if err!=nil{
					fmt.Println("SendOnlineToOther error:",err)
					return
				}		
		 }

		 return


}



func  (this *ProcDetail)  SendOnlineToOther(id int) (err error){
          
	         var  mes message.Message
			 mes.Type =message.ReNewListMesType

			 var  renewInfo message.RenewList
			 renewInfo.UsrId = id
			 renewInfo.Status = message.OnlineStatus

			 data, err := json.Marshal(renewInfo)
			 if err!=nil{
				 fmt.Println("json.Marshal fail:",err)
				 return
			 }
			 mes.Data =string(data)

			 data ,err = json.Marshal(mes)
			 if err!=nil{
				fmt.Println("json.Marshal fail:",err)
				return
			}
               
			 rwP := utils.RWPkg{
				Con:this.Con,
			 }
             err = rwP.WritePkg(data)
			 if err!=nil{
				 fmt.Println("WritePkg fail:",err)
				 return
			 }
	        
			 return


}



func (this *ProcDetail) ServerProcessRegister(ms *message.Message) (err error){

	   var RegisterMes message.RegisterMe
	   err =  json.Unmarshal([]byte(ms.Data),&RegisterMes)
	   if err!=nil{
		   fmt.Println("json.Unmarshal fail err",err)
		   return
	   }

       
		 
		 var reMes message.Message
		 var  reRegMes  message.RegisterReMes

	     reMes.Type =message.RegisterReMesType

		 err =  model.MyUsrDao.Register(&RegisterMes.User)

		

		 if err!=nil{
              if err==model.USR_ERROR_EXIST{
				reRegMes.Code =500
				reRegMes.Error=err.Error()
			  }

		 }else{
			reRegMes.Code =200
			fmt.Println("註冊成功")
		 }


        data,  err := json.Marshal(reRegMes)
		if err!=nil{
			fmt.Println("json.Marshal fail:",err)
			return
		}
		reMes.Data = string(data)

		data,err =json.Marshal(reMes)
		if err!=nil{
			fmt.Println("Marshal fail:",err)
			return
		}

		rwPackage := &utils.RWPkg{
			Con:this.Con,
		}

		err=  rwPackage.WritePkg(data)
		return
		 

	
}

func (this *ProcDetail) ServerProcessLogin(ms *message.Message) (err error){
	var loginMes  message.LoginMes 
   err =  json.Unmarshal([]byte(ms.Data),&loginMes)
   if err!=nil{
	   fmt.Println("json.Unmarshal fail err",err)
	   return
   }

   var reMes message.Message
   reMes.Type=message.LoginResMesType
   var loginResMes message.LoginReMes
      
    _ ,err =  model.MyUsrDao.Login(loginMes.UserId,loginMes.UserPwd)

	if err!=nil{

		if err == model.USR_ERROR_NONEXIST {
		 loginResMes.Code =403
		 loginResMes.Error =err.Error()
		  
		}else if err == model.USR_ERROR_EXIST {
		 loginResMes.Code =500
		 loginResMes.Error =err.Error()
		 }else{
		 loginResMes.Code =505
		 loginResMes.Error ="未知訊息錯誤"
		 }

		 
	}else{
		loginResMes.Code = 200
		this.Id =loginMes.UserId
		UsrOlineList.AddUsrOnList(this)
	    for usrid,_:=range UsrOlineList.usrOnList {
			loginResMes.UsrList = append( loginResMes.UsrList ,usrid) 
		}
		 this.UpdateOnlineList(loginMes.UserId)
	     
		
	   fmt.Println("登入成功")
	}
	

	 
    
	

data,  err := json.Marshal(loginResMes)
  if err!=nil{
	  fmt.Println("json.Marshal fail:",err)
	  return
  }
  reMes.Data = string(data)

  data,err =json.Marshal(reMes)
  if err!=nil{
	  fmt.Println("Marshal fail:",err)
	  return
  }

   rwPackage := &utils.RWPkg{
	   Con:this.Con,
   }

  err=  rwPackage.WritePkg(data)
  return
  

}
