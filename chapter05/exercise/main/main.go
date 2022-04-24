package main


import (
	"fmt"
	U "go_Code/chapter05/exercise/util"
)

func main(){
    
	var negative,positive int

    positive,negative =  U.JudgeNum()
	fmt.Printf("~~正數有%d個,負數有%d個\n",positive,negative)

}
