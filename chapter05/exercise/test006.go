package main

import(
	"fmt"
	
	
)

func judgeNum() (int,int){
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

func main(){
    
	var negative,positive int

    positive,negative =  judgeNum()
	fmt.Printf("正數有%d個,負數有%d個\n",positive,negative)

}


