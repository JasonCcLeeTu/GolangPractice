package chapter04

import(
	"fmt"
	_ "unsafe"
)

func mainTr(){

var i int32 =100

var f float32 = float32(i)

fmt.Printf("f=%v",f)

}