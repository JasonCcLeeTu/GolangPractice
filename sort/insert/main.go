package main
import(
	"fmt"

)


func main(){

   array:=&[...]int{1,2,3,8,7,4,6,5}
   insertSort(array)
   fmt.Println("結果:",*array)

}




func insertSort(array *[8]int){

    for i:=1;i<len(array) ;i++{

		sortIndex:=i-1
		sortValue:=array[i]
	
	
		for sortIndex>-1 &&  sortValue > array[sortIndex]{
			array[sortIndex+1]=array[sortIndex]
			sortIndex--
		}
	    
		if sortIndex+1!=i{

			array[sortIndex+1]=sortValue
		}

		fmt.Println(array)
	}

	

}