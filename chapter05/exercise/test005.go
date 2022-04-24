package main

import(
	"fmt"
	_"math"
)

func main(){
   
   var height float64 
   var money  float64
   var handsome bool

   fmt.Println("請輸入條件 身高 金錢  是否是帥哥")
   fmt.Scanf("  %f  %f  %t ",&height,&money,&handsome)

   if height >= 180 &&  money >= 10000000 && handsome {
	   fmt.Println("我一定嫁給他")
   }else if height >= 180 ||  money >= 10000000 || handsome{
	   fmt.Println("比上不足，比下有餘")
   }else{
	   fmt.Println("我不嫁")
   }

	
}