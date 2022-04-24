package main

import(
	"fmt"
)



func quickSort(array *[]int) []int{
      
	  
	   var left,right []int
	   arrayVal:=*array
	   if len(arrayVal)<2{
		   return arrayVal
	   }
       pivot:=arrayVal[0]
	   CompareVal:=arrayVal[1:]

	   for _,val:=range CompareVal{

		   if val >=pivot{
           right=  append(right,val)
		   }else{
			   left=append(left,val)
		   }
	   }

	   mid:=append([]int{pivot},quickSort(&right)...)
	   return append(quickSort(&left),mid...)


	  
	 

}

func main(){

    num:=&[]int{56,4,64,43,21,23,0,23,6,10,30,26,1,3,-1}
  // num:=&[]int{56,4,64,43,21,0,23,6,10,30,26,1,3,-1}
    
	
    res:= quickSort(num)

    fmt.Println(res)





}