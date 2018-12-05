package main
import "fmt"

var vname1,vname2,vname3 =1,"aa",false
var vname4,vname5,vname6 int
var x,y int = 1,2
var (
	a int 
	b string
)
	



func  main() {
 	fmt.Println(vname1,vname2,vname3)

 	//只能在函数内赋值
 	vname4,vname5,vname6=2,3,4
 	fmt.Println(vname4,vname5,vname6)
	//不带声明格式的，只能在函数内赋值
 	vname7,vname8 :=9,10
 	fmt.Println(vname7,vname8)

 	x,y=11,22
 	a,b = 33,"aaa"
	fmt.Println(x,y)
	fmt.Println(a,b)


	var s1 string="aaa"
	fmt.Println(*&s1)

 }