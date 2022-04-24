package process

import(
	"fmt"
	"go_Code/chatProject/common/message"
	"encoding/json"
	"encoding/binary"
	"net"
	"go_Code/chatProject/client/utils"
	"os"
	_"go_Code/chatProject/client/model"
	
)

type  UsrProc struct{

}

func  (this *UsrProc) Register(userId int, userPwd string, userName string) (err error){

	con,err:=	net.Dial("tcp","0.0.0.0:8888")
	defer con.Close()
	if err!=nil{
		fmt.Println("Dial error")
		return
	}

	var Sendmessage message.Message
	Sendmessage.Type =message.RegisterMesType
	

    RegisterMes := message.RegisterMe{}
      user := message.User{}
	  user.UsrId=userId
	  user.UsrPwd=userPwd
	  user.UsrName=userName
	  RegisterMes.User = user
    
	Byte,err:=json.Marshal(RegisterMes)

	if err!=nil{
		fmt.Println("Marshal loginMes error:",err)
		return
	}
	Sendmessage.Data=string(Byte)

	 data,err:= json.Marshal(Sendmessage)
   
	 if err!=nil{
	   fmt.Println("Marshal Sendmessage error:",err)
	   return
   }
    


   bytes :=make([]byte,8)
	  
		var pklen uint32 
		pklen = uint32(len(data))
	  binary.BigEndian.PutUint32(bytes[:8],pklen)
       
	  n,err :=  con.Write(bytes)
	  if n!=len(bytes) || err!=nil{
		  fmt.Println("conn write error:",nil)
		  return
	  }
       
	  _,err= con.Write(data)

	  if err!=nil{
		  fmt.Println("con wirte data error",err)
		  return
	  }




	    rwPackage:=  &utils.RWPkg{
			Con:con,
		}
	    mes,err := rwPackage.ReadPkg()
		if err!=nil{
			fmt.Println("readPkg fail",err)
		}


		 resgisterReMes  := message.RegisterReMes{}
		  err =json.Unmarshal([]byte(mes.Data), &resgisterReMes)
		  if err!=nil{
			  return
		  }

		  if  resgisterReMes.Code == 200{
			  fmt.Println("註冊成功")
			  os.Exit(0)
		  }else{
			  fmt.Println(resgisterReMes.Error)
			  os.Exit(0)
		  }

		  return


	
	
 
} 

func (this *UsrProc) Login(userId int,userPwd string) (err error) {
    
   con,err:=	net.Dial("tcp","0.0.0.0:8888")
   defer con.Close()
   if err!=nil{
	   fmt.Println("Dial error")
	   return
   }

    var Sendmessage message.Message
	Sendmessage.Type =message.LoginMesType

    loginMes := message.LoginMes{}
	loginMes.UserId=userId
	loginMes.UserPwd=userPwd
     
	 Byte,err:=json.Marshal(loginMes)

       

	 if err!=nil{
		 fmt.Println("Marshal loginMes error:",err)
		 return
	 }
	 Sendmessage.Data=string(Byte)

	  data,err:= json.Marshal(Sendmessage)
	
	  if err!=nil{
		fmt.Println("Marshal Sendmessage error:",err)
		return
	}
      //  var lenbyte [4]byte
	//	bytes := lenbyte[:4]
	bytes :=make([]byte,8)
	  
		var pklen uint32 
		pklen = uint32(len(data))
	  binary.BigEndian.PutUint32(bytes[:8],pklen)
       
	  n,err :=  con.Write(bytes)
	  if n!=len(bytes) || err!=nil{
		  fmt.Println("conn write error:",nil)
		  return
	  }
       
	  _,err= con.Write(data)

	  if err!=nil{
		  fmt.Println("con wirte data error",err)
		  return
	  }
	    rwPackage:=  &utils.RWPkg{
			Con:con,
		}
	    mes,err := rwPackage.ReadPkg()
		if err!=nil{
			fmt.Println("readPkg fail",err)
		}
		
		var  loginReMes   message.LoginReMes
		 err=json.Unmarshal([]byte(mes.Data),&loginReMes)
		 if loginReMes.Code == 200 {

			for _,val:=range loginReMes.UsrList{
				if val == userId{
					continue
				}
				usr := &message.User{
					UsrId:val,
					UsrStatus:message.OnlineStatus,
				}
				onlineUsrList[val]=usr
				fmt.Printf("目前 %d 在線\n",val)

			}
			fmt.Println("登入成功")
			currentUser.UsrId=userId
			currentUser.Con =con
			
			
			go Monitor(con)
			for{
				ShowMenu()
			}
		 }else if loginReMes.Code == 500 {	
			fmt.Println(loginReMes.Error)
		 }else if loginReMes.Code == 403 {
			fmt.Println(loginReMes.Error)
		 }else {
                fmt.Println(loginReMes.Error)
		 }

        
	 

    return
    
    

}