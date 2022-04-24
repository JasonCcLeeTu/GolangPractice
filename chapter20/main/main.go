package main

import(
   "fmt"
   "errors"
   
)


type heroNode struct{
	id int
	name string
	nextHero *heroNode
}









func   main(){

	    headNode:=&heroNode{}
        hero1:=&heroNode{
			id:1,
			name:"jason",
		}
		hero2:=&heroNode{
			id:2,
			name:"david",
		}
		hero3:=&heroNode{
			id:3,
			name:"benson",
		}
		hero4:=&heroNode{
			id:4,
			name:"judy",
		}

		add(headNode,hero3)
		add(headNode,hero4)
		add(headNode,hero1)
		add(headNode,hero2)
		show(headNode)
		delete(headNode,3)
		show(headNode)






}


func add(head *heroNode,newHero *heroNode) (err error){
        
	temp:= head
	flag:=false
	   for {
              
		       if temp.nextHero==nil{
				flag=true
				   break
			   }else if temp.nextHero.id > newHero.id{
				   flag=true
				   break
			   }else if temp.nextHero.id==newHero.id{
				   return errors.New("編號已經存在")
			   }
			   temp=temp.nextHero
           
	   }
	   if flag{
		   newHero.nextHero=temp.nextHero
		   temp.nextHero=newHero
	   }

	   return


}


func delete(head *heroNode,id int) (err error){

	temp:=head
   
	for{

		if temp.nextHero==nil{
		 err =	fmt.Errorf("沒有找到id:%d",id)
			return
		}else if temp.nextHero.id==id{
			
			break
		}

		temp=temp.nextHero
	}


	

	 temp.nextHero=temp.nextHero.nextHero
	 return

	 


}

func show(head *heroNode) {
        temp:=head
	
   for{

	   if temp.nextHero==nil{
		   break
	   }
		
		fmt.Printf("id:%d:%s --> ",temp.nextHero.id,temp.nextHero.name)
		temp=temp.nextHero
	}

   
   fmt.Println()
   return


}