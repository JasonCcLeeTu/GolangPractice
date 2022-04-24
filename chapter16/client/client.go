package main

import (

	"net"
	"fmt"
	"os"
	"bufio"
	"strings"
	
)


func main(){
    
	conn,err :=  net.Dial("tcp","192.168.1.101:8888")
	defer conn.Close()

	if err != nil{

		fmt.Println("用戶端發送訊息錯誤:",err)
	}

	
	for{
        
		reader := bufio.NewReader(os.Stdin)
		str,err:=  reader.ReadString('\n')
		if err!=nil{
			fmt.Println("ReadString錯誤",err)
		}
  
         str =strings.Trim(str," \r\n")
		 if str=="exit"{
			 fmt.Println("退出")
			 break
		 }
		n,err := conn.Write([]byte(str))
  
		fmt.Printf("客戶端發送%d個byte數據\n",n)
  

	}
    


}