package main 
import(

	"fmt"
	_"github.com/garyburd/redigo/redis"
)


type  test struct{

	age int
	name string
}



var  maptest map[int]*test



func main(){
     
	mapt1 := make( map[int]*test,10)
	t1 :=&test{
		age:12,
		name:"jason",
	}
	n:=1
	mapt1[1]=t1
	mapt1[n].age=100
	val := mapt1[1]
	fmt.Println(*val)
	
}