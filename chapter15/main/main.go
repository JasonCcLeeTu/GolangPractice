package main

import(
	"reflect"
	"fmt"
)

type Monster struct{

	Name string `json:"monsterName"`
	Age int     `json:"monsterAge"`
}

func (monster *Monster) Sum(s1 int,s2 int) int{
	return s1+s2
}

func (monster *Monster) ShowInfo(){

	   fmt.Printf("名稱:%s,年紀:%d",monster.Name,monster.Age)
}

func (monster *Monster) SetInfo(name string,age int){
	monster.Name=name
	monster.Age=age
}



func  reOperation (inter interface{}){

  //  rType :=  reflect.TypeOf(inter).Elem()
    rVal := reflect.ValueOf(inter).Elem()
	kind := rVal.Kind()

	 if kind !=  reflect.Struct{
		 fmt.Println("it is not struct ,is",kind)
		  return
	 }
	//  fieldNum :=  rVal.NumField()
    
	//  for i:=0;i<fieldNum;i++ {
    //       fmt.Printf(" %v \n",rVal.Field(i))    
	// 	  tag:= rType.Field(i).Tag.Get("json")
	// 	  if tag !="" {
	// 		  fmt.Println(tag)
	// 	  }
	//  }

  
	 // reflect.ValueOf(inter).Method(1).Call(nil)
    //    var param []reflect.Value
	//    param =append(param,reflect.ValueOf(100))
	//    param =append(param,reflect.ValueOf(200))

	//    res :=  reflect.ValueOf(inter).Method(2).Call(param)

	//    fmt.Println( res[0].Int())

	 rVal.FieldByName("Name").SetString("張家偉")
	 fmt.Println(rVal.FieldByName("Name").String())



	






}

func main(){

  monster :=Monster{
	  Name:"jason",
	  Age: 28,
  }


   // monster.ShowInfo()

  reOperation(&monster)
	




}