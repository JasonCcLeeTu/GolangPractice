package message



const(
     
	LoginMesType="LoginMes"
	LoginResMesType="LoginResMes"
	RegisterMesType="RegisterMes"
	RegisterReMesType ="RegisterReMes"
	ReNewListMesType= "ReNewListMesType"
	SendMesType ="SendMesType"

)


const(

	OnlineStatus = iota
	OfflineStatus 
	
)

type Message struct{

	 Type string  `json:"type"`
	 Data string `json:"data"`
}


type LoginMes struct{

	UserId int   `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginReMes struct {

	 Code int  `json:"code"`
	 Error string  `json:"error"`
	 UsrList  []int 
}


type  RegisterMe struct{
	  User User
}

type RegisterReMes struct{

	Code int  `json:"code"`
	Error string  `json:"error"`
	

}


type  RenewList struct{

	 UsrId int
	 Status int
}


type SendMes struct{
      
	  Content string
	  User
	  

}