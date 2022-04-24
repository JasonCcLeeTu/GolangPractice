package main

import(
	
     
	"encoding/json"
	"fmt"
)




type  student  struct{
	Name string
	Age int
	Address []string
}



func main(){
        
	str:="{\"Name\":\"jason\",\"Age\":28,\"Address\":[\"張張張\",\"橙橙橙\"]}"
    var student student
	err:= json.Unmarshal([]byte(str),&student)
	if err!=nil{
		fmt.Println(nil)
	}

	fmt.Println(student)
  
  

}