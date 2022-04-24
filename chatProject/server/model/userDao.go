package  model

import (
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"fmt"
	"go_Code/chatProject/common/message"
	
)


var  (
	MyUsrDao *UsrDao
)

type  UsrDao struct{

	 pool *redis.Pool
}

func  NewUsrDao(pool *redis.Pool)  (usrDao *UsrDao){

	usrDao = &UsrDao{
		 pool:pool,
	 }
	 return
}




func (this *UsrDao)Register(usr *message.User) ( err error){

	con:=  this.pool.Get()
	defer con.Close()
  
	_ ,err = this.getUsrById(con,usr.UsrId)

	if err ==nil{
		err = USR_ERROR_EXIST
	   return
	}

	 res,err:=json.Marshal(usr)
	 if err!=nil{
		 return
	 }

	_,err = con.Do("hset","users",usr.UsrId,string(res))
	if err!=nil{
		fmt.Println("註冊錯誤",err)
		return 
	}
	return

  
}



  



func (this *UsrDao)getUsrById(con redis.Conn,id int) (usr *User,err error){
             
	      res,err := redis.String(con.Do("hget","users",id))
		  if err!=nil{
			  if err==redis.ErrNil {
				  err = USR_ERROR_NONEXIST
			  }
			  return
		  }
           
		   usr = &User{}
		  err=  json.Unmarshal([]byte(res),usr)
		  if err!=nil{
           fmt.Println("Unmarshal error:",err)
			return 
		  }
		  return 





}

func (this *UsrDao)Login(id int,pwd string) (usr *User, err error){

      con:=  this.pool.Get()
	  defer con.Close()
	
      usr ,err = this.getUsrById(con,id)

	  if err !=nil{
         return
	  }
	  if usr.UsrPwd!=pwd {
		  err=USR_ERROR_PWD
	  }
	  return

	
    

}