package testMonster

import(
	"testing"
     "go_Code/exercise02/monster"
	
	"fmt"
)


func TestStore(t *testing.T){
           
     monster := monsterM.Monster {
		 Name:"jason",
		 Age:28,
		 Skill:"覺醒",
	 }

	 path:="d:/monster.txt"
	 res:= monster.Store(path)
	 if !res {
		 t.Fatalf("Store()執行錯誤")
	 } 
	 t.Logf("Store()執行成功")

}


func TestReStore(t *testing.T){

	path:="d:/monster.txt"
	var monster monsterM.Monster 

	res:= monster.ReStore(path)
	if !res{
		t.Fatalf("ReStore()執行錯誤")
	}

	t.Logf("~~~ReStore()執行成功")
	fmt.Println(monster)

}