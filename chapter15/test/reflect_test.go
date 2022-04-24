package test

import(
	
	"testing"
	_"reflect"
	"fmt"
)


type student struct{

	Name string
	Age int
}

func a(b []int){

   b[0]=12
    fmt.Printf("func a: %p \n ",&b)

}

func TestReflect(t *testing.T){
 
	   var b []int
	   b = append(b,1)
      a(b)
	  fmt.Printf("main: %p \n",&b)



}