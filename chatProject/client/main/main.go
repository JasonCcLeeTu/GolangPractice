package main

import(
	"fmt"
	"os"
	"go_Code/chatProject/client/process"
)

var(
	userId int
	userPwd string
	userName string
)


func main(){

	var choice int
	
	for true {

	
		fmt.Println("------歡迎登入多人聊天系統---------")
		fmt.Println("\t\t\t 1 登入聊天系統")
		fmt.Println("\t\t\t 2 註冊用戶")
		fmt.Println("\t\t\t 3 退出系統")
		fmt.Println("請選擇(1-3)")

		fmt.Scanf("%d \n",&choice)

		switch choice {

		case 1:
			
			fmt.Println("請輸入用戶id:")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("請輸入用戶密碼:")
			fmt.Scanf("%s\n",&userPwd)
            
			 userProcess:=&process.UsrProc{}
			 err:=userProcess.Login(userId,userPwd)
			 if err!=nil{
				 fmt.Println("Login fail:",err)
			 }
		case 2:
			fmt.Println("請輸入用戶id:")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("請輸入用戶密碼:")
			fmt.Scanf("%s\n",&userPwd)
			fmt.Println("請輸入用戶名稱:")
			fmt.Scanf("%s\n",&userName)
			userProcess:=&process.UsrProc{}
			err :=userProcess.Register(userId,userPwd,userName)
			if err!=nil{
				fmt.Println("Register fail:",err)
			}

			
		      
		case 3:
			fmt.Println("退出系統")
          os.Exit(0)
		default:
			fmt.Println("輸入錯誤")

		}
	}	

	// if choice==1{
	// 		fmt.Println("登入")
	// 		fmt.Println("請輸入用戶id:")
	// 		fmt.Scanf("%d\n",&userId)
	// 		fmt.Println("請輸入用戶密碼:")
	// 		fmt.Scanf("%s\n",&userPwd)

	// 		 err:=login(userId,userPwd)
		

	// }else if choice==2{


	// }





}