package main

import(

	"fmt"

	
	
)



func putIntChan(intChan chan int){

	for i:=1;i<=80;i++{
		intChan<- i
	}
	close(intChan)
}

func putIntChan2(intChan chan int){

	for i:=1;i<=800;i++{
		intChan<- i
	}
	
}


func getPrimeChan(intChan chan int,primeChan chan int,existChan chan bool){

          for{
			   isPrime := true
               num,ok:=<-intChan
			   if !ok {
				break
				}
			   for i:=2;i<num;i++{
                     if num%i==0{
						 isPrime = false
						 break
					 }
			   }

			   if isPrime {
				   primeChan<-num
			   }
                
			

		  }

		  existChan<-true


}


func main(){
   
	intChan := make(chan int,1000)
	primeChan :=make(chan int ,4000)
	existChan:=make(chan bool,4)
	testChan :=make(chan int,1000)

	go putIntChan(intChan)
	go putIntChan2(testChan)

	for i:=0;i<4;i++{
		go getPrimeChan(intChan,primeChan,existChan)
	}
     
	go func(){

		for i:=0;i<4;i++{
			<-existChan
		}
	
		
	}()
      
	
	 for {
           
		select {

		case val:=<-primeChan:
			fmt.Println("~",val)

		case val:=<-testChan:
			fmt.Println("testChan:",val)
		default:
			

			 return


		}


	 }
     
	 fmt.Println("成功退出")


	 


	

 
	
	
}