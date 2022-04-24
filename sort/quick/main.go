package main

import(
	"fmt"

)





func quickSort(array *[10]int,leftOr ,rightOr int){
       left:=leftOr
	   right:=rightOr
     mid:=(left+right)/2
	 label:=array[mid]

	 for left<right {
       
		for array[left]<label {
		
			left++
			
		}

		for  array[right]>label {
			
			right--
		}

        for left== right{
			break
		}

		array[right],array[left] = array[left],array[right]

	   
		if array[right]==label{
			left++
		}else if  array[left]==label{
			right--
		}
		
	
	
		

	}
      

	  

		  right--
		  left++
	  
	  
      
	
		
	  
	 
     if rightOr>left{
		quickSort(array,left,rightOr)
	 }

	  
     if leftOr<right{
		quickSort(array,leftOr,right)
	 }
	 
	 

    
}

func main(){

       array:=&[...]int{ -1, 2, 1, 13, 14, 15, -2 ,24, 29, 0}
       quickSort(array,0,9)
	   fmt.Println(array)


}