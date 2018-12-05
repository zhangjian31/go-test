package main
import (
"fmt"
)
/* 定义结构体 */
type Circle struct{
	r float64
}

func main() {
	var c1 Circle
	c1.r =10.00
	fmt.Println(c1.getArea())
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.r 即为 Circle 类型对象中的属性
	return 3.14*c.r*c.r
}