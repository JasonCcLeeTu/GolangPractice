package main

import (
   "fmt"
   "strconv"

)


func main(){
	var str string = "123456789"
   
	n1,_ := strconv.ParseInt(str,10,32)
	fmt.Printf("%T 內容:%v \n",n1,n1)
	

	var str2 string ="true"
	var b bool 
	b,_ = strconv.ParseBool(str2)
	fmt.Printf("%T %v \n ",b,b)
	
    var str3 string ="123.41234"
	var f float64
	f,_ = strconv.ParseFloat(str3,64)
	fmt.Printf("%T %v \n" ,f,f)

	var n3 int = 222
	 var str4 string
	str4 = strconv.Itoa(n3) 
	fmt.Printf("%T %v",str4,str4)

}