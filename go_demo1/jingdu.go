package main

import(
	"fmt"
)
func main() {
	a1:=1.69
	b1:=1.7
	c1:=a1*b1
	fmt.Println(c1)

	a2:=1690
	b2:=1700
	c2:=a2*b2
	fmt.Println(c2)
	fmt.Println(float64(c2)/1000000)
	fmt.Println(float64(123456789)/1000000)
}