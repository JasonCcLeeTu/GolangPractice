package main


import(
	
	"fmt"
)



func putNum(intChan chan int){
          
	 for i:=1;i<8000;i++{
		 intChan<-i
	 }
      
	 close(intChan)
	 


}

func caculator(intChan chan int ,putChan chan int ,finishChan chan bool){


	for{
		   flag:=true
		   num,ok:=<-intChan
		   if !ok{
			   break
		   }

		   for i:=2;i<num;i++{
			   if num%i==0{
                  flag= false
				  break
			   }
		   }

		   if flag{
			   putChan<-num
		   }


	}


	 finishChan<-true



}


func main(){

     intchan:=make(chan int,1000)
	 reschan:=make(chan int,1000)
	
	 finTakeChan:=make(chan bool,8)
     
	 
	 go putNum(intchan)
	 for i:=0;i<8;i++{
         go caculator(intchan,reschan,finTakeChan)
	 }

	 go func(){

		 for i:=0;i<8;i++{
			 <-finTakeChan
		 }
	
		 close(reschan)
	 }()

	 for{

       res,ok:=<-reschan
	   if !ok{
		   break
	   }
	   fmt.Println(res)

	 }


	 


}