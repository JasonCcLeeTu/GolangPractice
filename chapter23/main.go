package main

import(
	"fmt"
	"errors"
	"strconv"
)


type stack struct{
	maxSize int 
	topFlag int
	array  [30]int

}


func (this *stack) put(thing int) (err error){

       if this.topFlag==this.maxSize-1{
		fmt.Println("stack is full")
		  return  errors.New("stack is full")
		}

		this.topFlag++
		this.array[this.topFlag]=thing
		
		  return


}

func (this *stack) pop()  (thing int,err error){
	
	if this.topFlag == -1{
		fmt.Println("stack is empty")
		return -1,errors.New("stack is empty")
	
	}
	
	  thing = this.array[this.topFlag]
	  this.topFlag--
	  return

}


func (this *stack) show() (err error){
     
	if this.topFlag==-1{
        return errors.New("show fail:stack is empty")
	}

    curFlag:=this.topFlag
	for  curFlag>-1{
        fmt.Println(this.array[ curFlag])
		curFlag--
	}
	return
}

func (this *stack) judgeOper(thing int) bool {
	  if thing==42||thing==43||thing==45||thing==47 {
         return true
	  }
	  return false
}


func (this *stack) judgePriority(thing int) int{
	  switch thing {
	  case 42,47:
		return 1
	  case 43,45:
		return 0
	  default:
		return 0
	  }

	 
}


func (this *stack) calculate(num1, num2 ,oper int) int {
     
	switch oper{
	case 42:
	  return num2*num1
	case 47:
		return num2/num1
	case 43:
		return num2+num1
	case 45 :
		return num2-num1
	default:
		return -1
	}


}







func  main(){
	
	str:="20+10*60-40"
	index:=0
	combination :=""

	opr:=stack{
		maxSize:30,
		topFlag:-1,
	}
	num:=stack{
		maxSize:30,
		topFlag:-1,
	}
    
	

	for {
		val:=int(str[index:index+1][0]) 
		strVal:=str[index:index+1][0]

		isOtr := opr.judgeOper(val)
		if isOtr{
           if opr.topFlag==-1{
			   opr.put(val)
		   }else{
			    
			  beforeVal:=opr.array[opr.topFlag]
			  if opr.judgePriority(beforeVal) > opr.judgePriority(val){
				     oprbefore,_:=opr.pop()
					 n1,_:=num.pop()
					 n2,_:=num.pop()
					 res:=opr.calculate(n1,n2,oprbefore)
					 num.put(res)
					 opr.put(val)

			  }else{
					opr.put(val)
			  }

			}
				  
		}else{
			combination+=string(strVal)
			if  index==len(str)-1{
				
					integer64,_:=strconv.ParseInt(combination,10,64)
					num.put(int(integer64))
				
			}else{
				if opr.judgeOper(int(str[index+1:index+2][0])){
                    
				
					integer64,_:=strconv.ParseInt(combination,10,64)
					num.put(int(integer64))
					combination=""
				}

			}
			
		}

		  index++
          fmt.Println("num:",num.array)
		  fmt.Println("opr",opr.array)
		  fmt.Println("-------------------")

		  if index==len(str){
			  break
		  }

	}


	for opr.topFlag>-1{

		n1,_:=num.pop()
		n2,_:=num.pop()
		curOpr,_:=opr.pop()
		res:=opr.calculate(n1,n2,curOpr)
		num.put(res)
	}
     res,_:=num.pop()
	fmt.Println(res)
	
			
			   
			
        










	
     
	

}