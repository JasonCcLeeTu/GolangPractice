package main
import(
	"fmt"
	"os"
	"errors"
)



type  queue struct{

	maxsize int
	head int
	tail int
	array [5]int 


}


func main(){
    var key string
	
	queue := queue{
		maxsize:5,
		head:0,
		tail:0,

	}
	for{

		fmt.Println("1.put 新增")
		fmt.Println("2.get 取出")
		fmt.Println("3.show 顯示")
		fmt.Println("4.leave 離開")
		fmt.Println("請輸入:")
		fmt.Scanf("%s\n",&key)
	
	
		switch key{
		case "put":
			var num int
			fmt.Println("請輸入放入的數字:")
			fmt.Scanf("%d\n",&num)
			 err:=queue.put(num)
			 if err!=nil{
				 fmt.Println(err)
			 }
		case "get":
			num,err:=queue.get()
			if err!=nil{
				fmt.Println(err)
			}
			fmt.Println("取出:",num)
		case "show":
			queue.show()
			
		case "leave":
			os.Exit(0)
	
		}
	}

}


func (this *queue) put(num int) (err error){
       
	if (this.tail+1)%this.maxsize==this.head{
		return errors.New("array is full")
	}

	this.array[this.tail]=num
	this.tail=(this.tail+1)%this.maxsize
	
	return

}


func (this *queue) get() (num int,err error){

	if this.head==this.tail {
		return 0,errors.New("array is empty")
	}

	num=this.array[this.head]
	this.head++
	return 
}


func (this *queue) show()  {
    content:=(this.tail-this.head+this.maxsize)%this.maxsize

	h:=	this.head
	for i:=0;i<content;i++{
		
		fmt.Printf("array[%d]=%d \t",h,this.array[h])
		
		h= (h+1)%this.maxsize
			
	}
	fmt.Println()
	return
}





