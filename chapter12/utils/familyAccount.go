package utils

import(
	"fmt"
)



type familyAccount struct{
    haveDetail bool 
	balance float64
	detial string
	quit bool
	
	 
}

func (this *familyAccount)getAccountDetail(){

	fmt.Println("---------------當前收支明細紀錄------------------")
	if !this.haveDetail{
            fmt.Println("~~~目前沒有收支明細!!!!")
	}else{
		fmt.Println("收支\t\t帳戶金額\t\t收支金額\t\t說明")
		fmt.Println(this.detial)
	}
}


func (this *familyAccount) recordDetail(){
        var money float64
		var illustrate string
	fmt.Println("本次收入金額")
	fmt.Scanln(&money)
	fmt.Println("本次收入說明")
	fmt.Scanln(&illustrate)
	this.balance+=money
	this.haveDetail =true
	this.detial += fmt.Sprintf("\n收入\t\t %f\t\t %f\t\t %s",this.balance,money,illustrate)
}

func (this *familyAccount) payDetail(){
  	var money float64
  	var illustrate string
	fmt.Println("本次支出金額")
	fmt.Scanln(&money)
	fmt.Println("本次支出說明")
	fmt.Scanln(&illustrate)
	this.balance-=money
	this.haveDetail =true
	this.detial += fmt.Sprintf("\n支出\t\t %f\t\t %f\t\t %s",this.balance,money,illustrate)
}

func (this *familyAccount) quitSys(){
      var ans string
	  fmt.Println("您確定要退出? y/n")
	  for {
		
		fmt.Scanln(&ans)
		if ans=="y"  {
			 this.quit =true
			 break
		}else if ans=="n"{
			 break
		}else{
			fmt.Println("請重新輸入")
		}
	  }
	
}


func NewfamilyAccount() *familyAccount{

	return  &familyAccount{
		balance:10000.0,
	}
}


func (this *familyAccount) MainMenu(){

     var choice int
for{

	fmt.Println("---------------家庭收支記帳軟體------------------")
	fmt.Println("               1.收支明細")
	fmt.Println("               2.登記收入")
	fmt.Println("               3.登記支出")
	fmt.Println("               4.退出軟體")
	fmt.Println("請選擇<1-4>")
	fmt.Scanln(&choice)

           switch choice{

		   case 1:
               this.getAccountDetail()
		   case 2:
			   this.recordDetail()
		   case 3:
		     	this.payDetail()
		   case 4:
			  this.quitSys()

		   }
		if this.quit{
			break
		}
            

}

   fmt.Println("完成登出")

}