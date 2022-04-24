package main


import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"io"
	"strings"
)


type  chess struct{

	row int
	col int
	val int
}

func main(){
     
	var  chessList []chess
	array:=[11][11]int{}
	array[1][2]=1
	array[2][3]=2
     
	
	for _,val1:=range array{
		for _,val2:=range val1{
			fmt.Print(val2)
		}
		fmt.Println()
	}
	record := chess{
		row:11,
		col:11,
		val:0,
	}

     chessList = append(chessList,record)

	 for index1,val1:=range  array{

		for index2,val2:=range val1{
			if val2!=0{
				record:= chess{
					row:index1,
					col:index2,
					val:val2,
				}
				chessList = append(chessList,record)
			}
			

		}
	 }
    

	 path := write(chessList)
	 read(path)


    

}


func  write (chessList []chess) string {
         path:="d:\\chess.txt"
		
        file ,err := os.OpenFile(path,os.O_CREATE |os.O_WRONLY,0666 )
		if err!=nil{
			fmt.Println("openfile fail:",err)
		}
	    writer := bufio.NewWriter(file)
		
		        
		for _,val:=range chessList{

			writer.WriteString(strconv.Itoa(val.row)+"\t")
			writer.WriteString(strconv.Itoa(val.col)+"\t")
			writer.WriteString(strconv.Itoa(val.val)+"\n")
			writer.Flush()

		}
		
   
        return path
		
		
}


func read (path string)  {
     
	 array:=[11][11]int{}
    file,err:= os.OpenFile(path,os.O_RDONLY,0666)
	if err!=nil{
		fmt.Println("OpenFile fail:",err)
	}
	reader:= bufio.NewReader(file)
	  for{

		  
		  str ,err :=  reader.ReadString('\n')
		  
		  if err==io.EOF {
			break
		  }else if err!=nil{
			fmt.Println("ReadString fail:",err)
		  }
		  str =strings.Trim(str,"\n")
		   var s []string
		   s= strings.Split(str,"\t")
		//    for _,val :=range s{
		// 	fmt.Printf("%q",val)
		//    }
		 
                
			   row,_:=strconv.Atoi(s[0])
			   col,_:=strconv.Atoi(s[1])
			   value,_:=strconv.Atoi(s[2])
			   if row>=11 || col>=11{
				   continue
			   }
			   array[row][col]=value
		   
		  
		
	  }

	  fmt.Println("\n\n")
	  for _,val1:=range array{
		for _,val2:=range val1{
			fmt.Print(val2)
		}
		fmt.Println()
	}

	  


}