package utils

import(
	"net"
	"encoding/json"
	"go_Code/chatProject/common/message"
	"fmt"
	"encoding/binary"
	_"io"
)

type   RWPkg struct{

	Con net.Conn
	buf [8096]byte
}


func (this *RWPkg) WritePkg(data []byte) (err error){
	//bytes :=make([]byte,8)
	      bufs :=this.buf[:8]
		var pklen uint32 
		pklen = uint32(len(data))
	  binary.BigEndian.PutUint32(bufs[:8],pklen)
     
	  n,err :=  this.Con.Write(this.buf[:8])
	  
	  
	  if n!=len(bufs) || err!=nil{
		  fmt.Println("conn write error:",nil)
		  return
	  }
       
	 
	  n,err= this.Con.Write(data)

	  if  n!=int(pklen) ||err!=nil{
		  fmt.Println("con wirte data error",err)
		  return
	  }
	  return
          
}


func (this *RWPkg)  ReadPkg() (mes message.Message,  err error){

	//buf:=make([]byte,8096)
	bufs :=this.buf[:8]
	fmt.Println("讀取客戶端消息")
	n,err:=   this.Con.Read(bufs[:8])
  
	if   err!=nil{
		
		return
	}
   
	
	
	var dataLen uint32
	 dataLen= binary.BigEndian.Uint32(bufs[:n])
	 
	 
	 n,err= this.Con.Read(bufs[:dataLen])
    
	 if n!=int(dataLen)||err!=nil{
		return
	 }

	err = json.Unmarshal(bufs[:dataLen],&mes)

	if err!=nil{
		return
	}
	
    

    return


}


