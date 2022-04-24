package main

import(
	"fmt"

)


func main(){
	array:=[...]int{1,2,3,4,5,6}



    
	for j:=0;j<len(array);j++{

		max:=array[j]
		maxIndex:=j
	
		for i:=j+1;i<len(array);i++{
			if array[i]>max{
				max=array[i]
				maxIndex=i
			}
		}
		
		array[j], array[maxIndex]  = max,array[j]

	}	

	fmt.Println(array)



}