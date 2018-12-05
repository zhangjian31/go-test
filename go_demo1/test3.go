package main
import "fmt"
const const_a int = 100
const c_a,c_b ="aaa",333

//常量用作枚举
const(
	SUCCESS =100
	FAIL =404
)
func main() {
	const const_b int = 200
	fmt.Println(const_a,const_b)
	fmt.Println(c_a,c_b)

//类型要一致PI&r
	const PI float32 =3.14
	var r float32 =2
	var result float32 = PI * r * r 
	fmt.Printf("面积为 : %f\n", result)

//常量可以定义不使用
	fmt.Println(SUCCESS)


	fmt.Println(len("abc"))

	// iota，特殊常量，可以认为是一个可以被编译器修改的常量。




	

}