package main
import(
	"fmt"
   _  "encoding/json"

  
 
)



type  number struct{

}

func (num *number) judge(numb int) {

     if numb%2==0{
          fmt.Println("偶數")
     }else{
          fmt.Println("奇數")
     }
}


func (num *number) calculation(arr *[3][3]int) {

      for i:=0 ;i<len(*arr);i++{

          for j:=0;j<len((*arr)[i]);j++{
               if i < j{
             
                    temp :=(*arr)[i][j]
                    (*arr)[i][j]=(*arr)[j][i]
                    (*arr)[j][i]=temp
                
                
     
               }
          }

         
      }
   

    
    
     

}

func main(){

  
var map1 *map[string]int = &map[string]int{

    "~abc":123,
    "~ert":456,
    "~ert":456,
}

fmt.Println(*map1)
}