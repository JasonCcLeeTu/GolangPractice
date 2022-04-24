package test
import(
	"fmt"
	"strconv"
)

var Letter string

type account struct{

	accountid string
	balance float64
	password int
	Id int

}



func (ac  *account) Deposit(money float64,pass int){
          if ac.password != pass{
			  fmt.Println("密碼錯誤")
			  return
		  }
            ac.balance+=money
			fmt.Printf("存錢成功，餘額%f \n",ac.balance)

}

func (ac *account) Withdraw(money float64,pass int){

	if ac.password != pass{
		fmt.Println("密碼錯誤")
		return
	}
	if !(20 < ac.balance-money)  {
		fmt.Println("餘額不足")
		return
	}
	ac.balance -=money
	fmt.Printf("提取成功，餘額%f",ac.balance)
}


func (ac *account) SetPassWord(pass int){
	if  len(strconv.Itoa(pass)) != 6 {
		fmt.Println("設定密碼錯誤")
		return
	}
	ac.password = pass
	fmt.Println("密碼設定成功")

}

func (ac *account) SetAccount(accountId string) {

	if len(accountId) >=6 && len(accountId) <=10{
	
		 ac.accountid=accountId
		  return
	}else{
		fmt.Println("帳號長度錯誤")
		return
	}

}

func NewAccount() *account{

	return new(account)

	

}
