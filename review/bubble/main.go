package main

import(
	"fmt"
)


func main(){
		array:=[...]int{10,44,2,-1,4,3,8,20,1,0,3}

		for j:=0;j<len(array)-1;j++{

			for i:=0;i<len(array)-j-1;i++{
					if array[i]>array[i+1]{
						array[i],array[i+1]=array[i+1],array[i]
					}
	
			}
		}

		fmt.Println(array)
}