package main

import(
	"fmt"
)


func main(){
    
	num:=[]int{10,8,4,7,0,1,-1,46,12,5,2,11,10}
    var smallest,index int
	for i:=0;i<len(num)-1;i++{
		for j:=i;j<len(num);j++{
			if j==i {
				smallest=num[j]
				index=j
			}else if smallest > num[j]{
				smallest=num[j]
				index=j
			}

		}
		num[i],num[index]=num[index],num[i]
		
	}

	fmt.Println(num)



}