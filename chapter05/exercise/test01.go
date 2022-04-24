package main

import(
	"fmt"
)

func main(){

  a:=10
  b:=5

  a = a+b
  b = a-b
  a = a-b

  fmt.Printf("a=%v,b=%v",a,b)

}