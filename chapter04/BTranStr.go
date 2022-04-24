package chapter04

import(
	"fmt"
    "strconv"

)

var Name string ="jason"

func main(){

	var n1 int = 12
    var n2 float32 = 2.334
	// var char byte = 'h'
	 var b bool = true
	var str string


	//   str =  fmt.Sprintf("%d",n1)
	//   fmt.Printf("種類%T %v \n",str,str)

	//   str = fmt.Sprintf("%f",n2)
	//   fmt.Printf("種類%T %v \n",str,str)

	//   str = fmt.Sprintf("%c",char)
	//   fmt.Printf("種類%T %v \n",str,str)

	//   str = fmt.Sprintf("%t",b)
	//   fmt.Printf("種類 %T %v \n",str,str)


	str = strconv.FormatInt(int64(n1),10)
	fmt.Printf("%T  %q \n", str,str)


	str = strconv.FormatFloat(float64(n2),'f',6,32)
	fmt.Printf("%T %q \n",str,str)

   str = strconv.FormatBool(b)
   fmt.Printf("%T %q",str,str)




}