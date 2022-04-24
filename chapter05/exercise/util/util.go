package util

import(
	"fmt"
)

func JudgeNum() (int,int){
	var num,negative,positive int
	for{
		fmt.Println("請輸入整數")
		fmt.Scanln(&num)
		if num > 0 {
			positive++
		}else if num <0 {
			negative++
		}else{
			break
		}
		
	}

	return positive,negative
}