package main
import(
_	"go_Code/chapter07/account"
	"fmt"
	
)

type  student struct{

}

func checkType(item... interface{}){

     for _,val:=range item{

		switch val.(type){
        case int64 :
			fmt.Println("int64,",val)
		case float64 :
			fmt.Println("float64,",val)
		case string :
			fmt.Println("string,",val)
		case byte :
			fmt.Println("byte,",val)
		case student :
			fmt.Println("student,",val)
		case *student :
			fmt.Println("*student,",val)
		default :
			fmt.Println("不包括,",val)

		}
	 }
   

}





func main(){

	a:=12
    b:="adad"
	c:='r'
	d:=12.5
    s:=student{}
	e:=&student{}

	checkType(a,b,c,d,e,s)
   
  
}