package main

import(
	"net"
	"fmt"
	"io"
)


func Operation(con net.Conn){

     defer con.Close()
	 for {

	   buf  :=  make([]byte,1024)
	   n,err:= con.Read(buf)

	   if  err==io.EOF{
          fmt.Println("客戶端斷線")
		  return
	   }else if err!=nil{
		
			fmt.Println("server端Read發生錯誤:",err)
			return
	   }
	  
	   
        fmt.Printf("用戶端: %v 訊息: \n",con.RemoteAddr().String())
		
		
	    fmt.Print(string(buf[:n])+"\n")

	 }
      




}


func main(){
     
 	 listen,err := net.Listen("tcp","0.0.0.0:8888")
    
	  
	  if err!=nil{
		  fmt.Println("Listen錯誤:",err)
	  }
	  fmt.Println("開始接收數據")

	  for {
		
	    conn,err :=	listen.Accept()

		if err !=nil{
			fmt.Println("接收錯誤",err)
			return
		}else{
			fmt.Printf("用戶端:%v 連接成功\n",conn.RemoteAddr().String())
		}
        
		go  Operation(conn)

		

	  }
	  defer listen.Close()
	  






}