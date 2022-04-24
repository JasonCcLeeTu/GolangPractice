package main


import(
	"fmt"
	"os"
)


type student struct{

	name string
	age int
	next *student
}


type link struct{
	head *student
}

func (this *link) delete(age int){

	cur:=this.head
	tail:=this.head
	for{
		if cur.age ==age{
			if  cur==this.head{
               this.head=cur.next
			}else{
                tail.next=cur.next
			}
			break
		}
		tail=cur
		cur=cur.next
		if cur==nil{
			fmt.Println("沒該資料")
			break
		}
	}
	
}


func (this *link) search(age int){

	cur:=this.head
	for{
		if cur.age ==age{
			fmt.Println("找到了")
			break
		}
		cur=cur.next
		if cur==nil{
			fmt.Println("沒找到")
			break
		}
	}
	
}


func (this *link) show(){
	cur:=this.head
	for{
		
		fmt.Printf("節點%d:名稱:%s,年紀:%d -> ",cur.age%100,cur.name,cur.age)
        cur=cur.next
		if cur==nil{
			break
		}
	}
	
	return
}

func (this *link) insert(member *student) (err error) {

	if this.head==nil{
		this.head=member
		return
	}

	curNode:=this.head
	beflNode:=this.head

	for{

		

		if curNode.age>=member.age {
			if curNode.next==nil{
				curNode.next=member
				break
			}
		}else {
			if curNode.next==nil{
				member.next=curNode
				this.head=member
				break
			}else{
				if curNode==beflNode{
					
					this.head=member
					member.next=curNode
				}else{
					beflNode.next=member
				
					member.next=curNode
				}
                
				 
				
				
				 break
			}

		}

		

		beflNode=curNode
		curNode=curNode.next
		if curNode==nil{
			break
		} 


	}

	return



}

type storage struct{
	 list [100]link
}


func (this *storage) add(name string,age int){
      node:= this.judge(age)
	   newSt:=&student{
		   name:name,
		   age:age,
	   }
	  
	 err:=  this.list[node].insert(newSt)
	 if err!=nil{
		 fmt.Println("insert fail:",err)
	 }
}

func (this *storage)  judge(age int) int{

	return age%100

}


func (this *storage) show(){
	for _,val:= range this.list{
		
		 if (val!= link{}) {
		 	val.show()
			 fmt.Println()
		}
		
	}
	
	
}

func (this *storage) search(age int){
	   node:= this.judge(age)

	   this.list[node].search(age)
	   

		
	
	
	
}

func (this *storage) delete(age int){
	node:= this.judge(age)

	this.list[node].delete(age)
	

	 
 
 
 
}





func main(){
	 var key string
	 depository :=storage{}

	for{

		fmt.Println("put輸入:")
		fmt.Println("sear查詢:")
		fmt.Println("del刪除:")
		fmt.Println("show顯示:")
		fmt.Println("exit離開:")
		fmt.Scanf("%s \n",&key)
	
		switch key{
		case "put":
			var age int
			var name string
			fmt.Println("輸入名稱:")
			fmt.Scanf("%s \n",&name)
			fmt.Println("輸入年齡:")
			fmt.Scanf("%d \n",&age)
			depository.add(name,age)
			
		case "sear":
			var age int
			fmt.Println("輸入年齡:")
			fmt.Scanf("%d \n",&age)
			depository.search(age)
		case "del":
			var age int
			fmt.Println("輸入年齡:")
			fmt.Scanf("%d \n",&age)
			depository.delete(age)
		case "show":
			depository.show()
		case "exit":
			os.Exit(0)
		}
	}














	


     

}
