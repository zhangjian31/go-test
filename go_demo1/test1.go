package main
import "fmt"

var a string ="哈哈哈"
var b ="xxx"
var c bool
func main(){
	/*var age int
	age = 10
	// age := 10

	name :="this is "
	xiaoming := "xiaoming"
	var result string
	result =name + xiaoming

	fmt.Println("Hello World" ,age,result)

	fmt.Println("--------------")
	*/

	//变量声明
	var v_name string
	var num int
	fmt.Println(v_name,num)

	//变量推断
	var v_b = "fafa"
	fmt.Println(v_b)

	//省略var,此时的变量不能是声明过的
	v_c :=10
	fmt.Println(v_c)

	fmt.Println(a,b,c)






}
