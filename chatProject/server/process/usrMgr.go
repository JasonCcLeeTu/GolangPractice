package process

import(
   "fmt"
)

var (
	UsrOlineList *usrMrg
)

type usrMrg struct{

	usrOnList map[int]*ProcDetail
}


func init(){

	UsrOlineList = &usrMrg{
		usrOnList:make(map[int]*ProcDetail,1024),
	}
}


func (this *usrMrg) AddUsrOnList(usr *ProcDetail){

		this.usrOnList[usr.Id]=usr

}

func (this *usrMrg) DelUsrOnList(id int) {
       
	 delete(this.usrOnList,id)

}

func (this *usrMrg) SearchOnList(id int)(err error, usr *ProcDetail){

	   val,ok:= this.usrOnList[id]
	   if ok {
		   usr =val
	   }else{
		   err = fmt.Errorf("成員:%d不在線",id)
	   }
	   return
}


func (this *usrMrg) SearchAllList()( usr map[int]*ProcDetail){

	 return  this.usrOnList
}

