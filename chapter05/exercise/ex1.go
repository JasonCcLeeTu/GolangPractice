package main

import(
	"fmt"
	"strconv"
)


func main(){

  var a float64

  var str string ="ina"

  a ,_ = strconv.ParseFloat(str,64)
  fmt.Println(a)


}