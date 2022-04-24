package monsterM

import(
	"encoding/json"
	"io/ioutil"
	"fmt"

)

type Monster struct{

	Name string
	Age int
	Skill string
}

func(this *Monster) Store(path string) bool {

	str,err:= json.Marshal(this)
	if err!=nil{
		fmt.Println(nil)
		return false
	}
    err =   ioutil.WriteFile(path,str,0666)

	
	 if err !=nil{
		 fmt.Println(nil)
		 return false
	 }

	 return true

	 

}

func (this *Monster) ReStore(path string) bool{

    file,err:=  ioutil.ReadFile(path)
	if err!=nil{
		fmt.Println(err)
		return false
	}
	err=json.Unmarshal(file,this)
	if err!=nil{
		fmt.Println(err)
		return false
	}

	return true


}
