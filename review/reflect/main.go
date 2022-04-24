package main

import(
	"reflect"
	"fmt"
)


type student struct{

	name string
	age int
}

func main(){

  var s student
  s.name="jason"
  s.age=22
  

  val:= reflect.ValueOf(s)
   fmt.Println(val.Field(0))


}