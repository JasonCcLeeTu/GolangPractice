package main
import(
	"fmt"
	"go_Code/exercise01/control"
	"go_Code/exercise01/model"
)

type CustomerView struct{
    customerservice   *control.CustomerService
	choice int
	quit bool
}


func (view *CustomerView) showInfo(){
     fmt.Println("--------------客戶列表-------------")
	 fmt.Println("編號\t\t姓名\t\t性別\t\t年齡\t\t電話\t\t信箱")
	 for i:=0;i<len(view.customerservice.CustomerList);i++{
		 fmt.Println(view.customerservice.CustomerList[i].GetList())
	 }
	 fmt.Println("--------------客戶列表完成-------------")
	   
}


func (view *CustomerView) addInfo(){
       var name,gender,phone,email string
	   var age int
     fmt.Println("---------------添加客戶---------------")
	 fmt.Println("姓名")
	 fmt.Scanln(&name)
	 fmt.Println("性別")
	 fmt.Scanln(&gender)
	 fmt.Println("年齡")
	 fmt.Scanln(&age)
	 fmt.Println("電話")
	 fmt.Scanln(&phone)
	 fmt.Println("信箱")
	 fmt.Scanln(&email)

	 newcustomer:= model.NewCustomer2( name ,gender ,
		age ,phone,email)
	
		view.customerservice.Add(newcustomer)
	fmt.Println("---------------添加完成---------------")




}

func (view *CustomerView) deleteInfo(){
        var index int
		var ans string
	  fmt.Println("------------刪除客戶-------------")
	  for {
		fmt.Println("請輸入要刪除的編號(-1退出)")
		fmt.Scanln(&index)
		if index==-1{
			return
		}
		
		 for {
		  fmt.Println("確定是否刪除y/n ?")
		  fmt.Scanln(&ans)
		  if ans=="y" || ans=="Y"{
	       if res :=  view.customerservice.Delete(index);!res{
			   fmt.Println("-----------------刪除失敗找不到資料-----------")
		   }else{
			   fmt.Println("----------------刪除完成--------------------")

		   }
		   return
			  
		  }else if ans=="n"||ans=="N"{
			  break
		  }else{
			  fmt.Println("確認刪除輸入錯誤，請重新輸入")
		  }
		 }
	  }
	  
	  
	 




}

func (view *CustomerView) update(){
     var number int
	 var name,gender,phone,email,age string

	  customerMap := make(map[string]interface{})
	
	 for {
		fmt.Println("-------------------修改客戶---------")
		fmt.Println("請輸入修改的編號(-1退出)")
		fmt.Scanln(&number)
		if number==-1{
			return
		}
		  res, index:=view.customerservice.UpdateShow(number)
		 if res==nil{
			 fmt.Println("--------------修改失敗---------")
			 continue
		 }

		 fmt.Printf("姓名(%s):",res.Name)
		 fmt.Scanln(&name)
		 customerMap["Name"] =name
		 fmt.Printf("性別(%s):",res.Gender)
		 fmt.Scanln(&gender)
		 customerMap["Gender"] = gender
		 fmt.Printf("年齡(%d):",res.Age)
		 fmt.Scanln(&age)
		 customerMap["Age"] =age
		 fmt.Printf("電話(%s):",res.Phone)
		 fmt.Scanln(&phone)
		 customerMap["Phone"]=phone
		 fmt.Printf("信箱(%s):",res.Email)
		 fmt.Scanln(&email)
		 customerMap["Email"] =email

		 view.customerservice.Update(index,customerMap)
		 fmt.Println("----------修改成功---------")
         break
		 
		 

	 }
	



}


func (view *CustomerView) quitSys() {
         var strQuit string
		 fmt.Println("確定是否退出y/n?")
	for{
		

		fmt.Scanln(&strQuit)
		if strQuit=="y"||strQuit=="Y"{
			view.quit=true
			break
		}else if strQuit=="n" || strQuit=="N"{
			 break
		}else{
			fmt.Println("輸入錯誤請重新輸入")
		}
	}
}




func ( view  *CustomerView) MainView(){

         for{

           fmt.Println("--------------客戶訊息管理軟件-------------")
		   fmt.Println("\t\t1.添加客戶")
		   fmt.Println("\t\t2.修改客戶")
		   fmt.Println("\t\t3.刪除客戶")
		   fmt.Println("\t\t4.客戶列表")
		   fmt.Println("\t\t5.退出")
		   fmt.Println("請選擇(1-5)")


            fmt.Scanln(&view.choice)

			switch view.choice{

			case 1:
				view.addInfo()
			case 2:
				view.update()
			case 3:
				view.deleteInfo()
			case 4:
				view.showInfo()
			case 5:
				view.quitSys()

			}

			if view.quit{
				break
			}



		 }


} 


func main(){
	
    // customerview := CustomerView{}
	// customerview.customerservice =control.NewService()
	// customerview.choice=0
		
	// customerview.MainView()
	//  customerMap:= make(map[string]interface{})
	//  customerMap["1"]=123
	//  customerMap["2"]="abc"
	//  fmt.Println(customerMap)

	//  var a map[string]int 
	//  a= make(map[string]int)
	//  a["1"]=0
	//  a["2"]=2
	//  a["3"]=3
	//  fmt.Println(a)
	// var b []int
	// b = make([]int,1)
	// b[0]=1
	// b= append(b,2)
	// //  b[2]=3
	// fmt.Println(b)

	model.TestStruct()


	

}


