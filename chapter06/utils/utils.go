package  utils
import(
	_"fmt"
)

var Age int
var Name string
var(
	Sum = func (a int, b int) int {
		return a+b
	}
)

func init(){
  
    Age=18
	Name="Jason"
    
}