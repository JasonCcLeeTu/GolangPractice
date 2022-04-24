package main

import(

	"fmt"
	"runtime"
)

func main(){
   

	cpuNum := runtime.CPUNum()
	fmt.Println(cpuNum)

}