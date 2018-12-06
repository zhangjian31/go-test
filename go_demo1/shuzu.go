package main

import "fmt"

func main() {
	
	//数组声明
	var balance1 [10] int

	//数组初始化
	//1.{}中的元素个数不能大于【】中的数字
	var balance2 = [5]int{1,2,3,4,5}
	//2.[]中不设置数组大小，根据元素个数设置书序大小
	var balance3 = [...]int{1,2,3,4,5,6,7,8,9}

	fmt.Println(balance1)
	fmt.Println(balance2)
	fmt.Println(balance3)

	var n [10]int
	var i,j int
	for i = 0; i < 10; i++ {
		n[i]=i+100
	}

	for j = 0; j < 10; j++ {
		fmt.Println(n[j])
	}


}