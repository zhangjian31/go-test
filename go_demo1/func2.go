package main
import (
"fmt"
"math"
)
func main() {
	//函数定义后可作为值来使用
	fmt.Println("###############")
	getSqrt :=func (x float64) float64{
		return math.Sqrt(x)
	}
	fmt.Println(getSqrt(9))
	fmt.Printf("%f\n",getSqrt(9))

// Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。
// 匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
	fmt.Println("###############")
	nextNumber:=getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber2:=getSequence()
	fmt.Println(nextNumber2())
	fmt.Println(nextNumber2())
	fmt.Println(nextNumber2())

	fmt.Println("###############")
	add_func:=add(2,3)
	fmt.Println(add_func())
	fmt.Println(add_func())
	fmt.Println(add_func())

	fmt.Println("###############")
	add2_func:=add2(1,2)
	fmt.Println(add2_func(1,1))
	fmt.Println(add2_func(0,0))
	fmt.Println(add2_func(2,2))
}

// 闭包是匿名函数
func getSequence() func () int {
	i:=0
	return func() int{
		i+=1
		return i
	}
}

func add(x1,x2 int) func()(int,int){
	i:=0
	return func()(int,int){
		i++
		return i,x1+x2
	}
}

func add2(x1,x2 int) func (x3,x4 int)(int,int,int) {
	i:=0
	return func (x3,x4 int)(int,int,int) {
		i++
		return i,x1+x2,x3+x4
	}
} 

