package main

import(
	"fmt"
)

func main(){
	num:=[]int{1,5,7,4,55,-1,14,0,30,8,4}

	for i:=0;i<len(num)-1;i++{
         max:=num[i+1]
		 for j:=i;j>=0;j--{
			 if max>num[j]{
				 num[j+1]=num[j]
				 if j==0{
					 num[j]=max
					 break
				 }
				 
			 }else if max <= num[j]{
                  num[j+1]=max
				  break
			 }

		 }
	}

	fmt.Println(num)

}