package model

import(
	"fmt"
	"go_Code/exercise01/test"
)

type Customer struct{

	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}


func TestStruct(){

	t1 := TestSs{"jason",22}
	fmt.Println(t1)

}


func (customer *Customer) GetList() string{
    return    fmt.Sprintf("%d\t\t%s\t\t%s\t\t%d\t\t%s\t\t%s",
						 customer.Id,customer.Name,customer.Gender,customer.Age,customer.Phone,customer.Email)
}



func NewCustomer(id int, name string,gender string,
		   age int,phone string,email string) Customer{

         newCustomer := Customer{
			 Id: id,
			 Name:name,
			 Gender:gender,
			 Age: age,
			 Phone: phone,
			 Email:email,
		 }

		 return  newCustomer



}

func NewCustomer2( name string,gender string,
	age int,phone string,email string) Customer{

  newCustomer := Customer{
	  
	  Name:name,
	  Gender:gender,
	  Age: age,
	  Phone: phone,
	  Email:email,
  }

  return  newCustomer



}