package main


import(
    "fmt"
)


type student struct{
	id int
	name string
	nextStudent *student
}

func main(){
          
	    head:=&student{}
		s1:=&student{
			id:1,
			name:"jason",
		}
		s2:=&student{
			id:2,
			name:"jacky",
		}
		s3:=&student{
			id:3,
			name:"judy",
		}
		s4:=&student{
			id:4,
			name:"david",
		}

		add(head,s1)
		add(head,s2)
		add(head,s3)
		add(head,s4)
		show(head)
		head =delete(head,1)
		
		show(head)
		head =delete(head,3)
		show(head)
		
			


}


func delete(head *student,id int) *student{
           
	   if head.nextStudent == nil{
		   fmt.Println("沒有值可刪除")
		   return head
	   }

	   if head.nextStudent == head && head.id==id {
		      head.nextStudent=nil
			  return head
	   }


      temp:=head
	  beforetemp:=head

	  for {
		  if beforetemp.nextStudent==head{
			  break
		  }
		  beforetemp=beforetemp.nextStudent
	  } 

	  for{
		   
           if temp.id==id{
                if temp==head{
					head=temp.nextStudent
				}
				beforetemp.nextStudent=temp.nextStudent


		   }
		   if temp.nextStudent==head{
			   break
		   }
		   temp=temp.nextStudent
		   beforetemp=beforetemp.nextStudent

	  }

	  return head

}

func add(head *student,newStudent *student){

          if head.nextStudent==nil{
			  head.id=newStudent.id
			  head.name=newStudent.name
			  head.nextStudent=head
			  fmt.Println("加入成功")
			  return
		  }
		  temp:=head

		  for{
                
			   if temp.nextStudent==head{
				  
				   break
			   }

			   temp=temp.nextStudent

		  }
		  temp.nextStudent=newStudent
		  newStudent.nextStudent=head

}


func show(head *student){

	if head.nextStudent==nil{
		fmt.Println("沒有內容")
		return
	}

	temp:=head

	for{
		fmt.Printf("id:%d name:%s -->",temp.id,temp.name)
         if temp.nextStudent==head{
			 break
		 }
		 
		 temp=temp.nextStudent
	}
	fmt.Println()
}




