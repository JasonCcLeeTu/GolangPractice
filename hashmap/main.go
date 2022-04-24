package main

import(
         "fmt"
		 "errors"
		 "os"
)


type storeMaterial struct{
	member [7]linkList
}


type linkList struct{
	 head *content
}

type content struct{
	id int
	name string
	next *content
}

func (this *storeMaterial) findStorePoint(id int) int {
     return   id%7

}

func (this *storeMaterial) add(id int , name string) (err error){
         
	  stPoint:= this.findStorePoint(id)
	  newContent:=&content{
		  id:id,
		  name:name,
	  }
	  this.member[stPoint].insert(newContent)
	
	  return

}


func (this *storeMaterial) show(){
      
	for _,val:=range this.member{
	  err:=	val.show()
	  if err!=nil{
		  fmt.Println(err)
	  }
	}
	

}

func (this *linkList) insert(cont *content) {
           
	  cur:=this.head
	  tail:=this.head
       
	  for{

		  if cur==nil{ //目前沒有值
			  	this.head=cont
				 
			 	break
		  }else{
			  if cont.id < cur.id{ //小到大 
	             if cur==this.head{ //判斷是否為頭且資料要放到頭前面 等於新資料當頭
                       cont.next=cur
					   this.head=cont
					  
					   break
				 }else{  //剩下的都是插入兩個值中間
                       tail.next=cont
					   cont.next=cur
					 
					   break
				 }
			  }else if cur.next ==nil{  //都沒找到代表加入的值最大直接放最後面
				  
                   cur.next=cont
				 
				   break
			  }
	
		  }
           tail=cur
		   cur=cur.next



	  }
	   
	  return
   
} 

func (this *linkList) show() (err error) {
       cur:=this.head
	   if cur==nil{
		   return errors.New("沒有資料顯示")
	   }
	   for{
		   if cur==nil{
			   break
		   }
		   fmt.Printf("第%d個鏈表= id:%d name:%s --> ",cur.id%7,cur.id,cur.name)
		   cur= cur.next
		
	   }
      
	   fmt.Println()
	   return

}


func (this *storeMaterial) search(id int)  (con *content) {
    point:=	this.findStorePoint(id)
	start:= this.member[point].head
	
	con=getContent(start,id)
	
	return

}


func  getContent (head *content,id int) (con *content){
	fmt.Println(head,id)
	if head==nil{
		fmt.Println("沒有資料可查詢")
		return
	}
	
	temp:=head
	for{
		if temp==nil{
			break
		}
		if temp.id==id{
			
			con=temp
			return 
		}
		temp=temp.next
	
	}
	return

}



func main(){
        
	 var material storeMaterial
    for{
		key:=""
        fmt.Println("put 加入")
		fmt.Println("sear 查找")
		fmt.Println("show 顯示")
		fmt.Println("leave 離開")
		fmt.Scanf("%s \n",&key)
		
		switch key{
		case "put":
			var id int
			var name string
			fmt.Println("輸入id:")
			fmt.Scanf("%d\n",&id)
			fmt.Println("輸入name:")
			fmt.Scanf("%s\n",&name)
			material.add(id,name)
			fmt.Println("添加成功")
		case  "sear":
			var id int
			fmt.Println("輸入id:")
			fmt.Scanf("%d\n",&id)
			res:=material.search(id)
			if res!=nil{
				fmt.Println(res)
			}
		case "show":
			material.show()
		case "leave":
			os.Exit(0)
		default:

		}

		


	}
   


   
}