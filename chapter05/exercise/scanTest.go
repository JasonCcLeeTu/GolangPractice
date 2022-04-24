package main

import(

	"fmt"
)

func main(){

    var name string
	var age int
	var  salary float32
	var pass bool

  //fmt.Println("請輸入 姓名、年齡、薪水、是否通過考試")
//   fmt.Println("請輸入 姓名")
//    fmt.Scanln(&name)

//    fmt.Println("請輸入 年齡")
//    fmt.Scanln(&age)
//    fmt.Println("請輸入 薪水")
//    fmt.Scanln(&salary)
//    fmt.Println("請輸入  是否通過考試")
//    fmt.Scanln(&pass)

   

    fmt.Println("請輸入 姓名 年紀 薪水 是否通過考試")
	fmt.Scanf("%s %d %f %t",&name,&age,&salary,&pass)

	fmt.Printf("姓名: %v 年紀: %v 薪水: %v 是否通過考試: %v",name,age,salary,pass)
   

}