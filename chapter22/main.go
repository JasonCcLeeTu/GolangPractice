package main

import(
	"fmt"

)

type student struct{
	id int
	next *student
}


func main(){

     head:=insert(5)
	 show(head)
	 delete(head,2,3)


}


func insert(account int) *student{

      if account<1{
		  fmt.Println("輸入的人數不能小於1")
		  return nil
	  }

	  first:=&student{}
       temp:=first
	  for i:=1;i<=account;i++{
		
		  if i==1{
			  first.id=i
			  first.next=temp
		  }else{
			newStu:=&student{}
			newStu.id=i
			newStu.next=temp.next 
			temp.next=newStu

		  }
		  temp=temp.next
		  
		  
	  }
	  return first


}


func show(head *student){

	  if head.next==nil{
		  fmt.Println("沒有學生在裡面")
		  return
	  }
     temp:=head
      for {
			fmt.Printf("學生id: %d -> ",temp.id)
           if temp.next==head{
			   break
		   }

		   temp=temp.next
	  }
	  fmt.Println()

}



func delete(head *student,start int,account int){
    
	if head.next==nil{
		fmt.Println("沒有任何內容")
		return
	}
	

	
	temp:=head
	before:=head

	for{

		 
		if before.next==head{
			break
		}
		before=before.next
	}
	
	for  i:=0;i<start-1;i++  {
   
		
		temp=temp.next
		before=before.next
	}

	
	for {
		   
		if temp==before{
			fmt.Println("取出id",temp.id)
			break
		}
		for i:=0;i<account;i++{
			temp=temp.next
			before=before.next

		}
		fmt.Printf("取出id: %d \n",temp.id)
         temp=temp.next
		 before.next=temp
		
		
	}

	

	
	   



}