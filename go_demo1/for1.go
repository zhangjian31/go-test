package main
// 连个for循环只差了一个冒号，结果大不相同
import "fmt"
func main() {
	for1()
	for2()
}

func for1() {
	var i=10
	fmt.Println("------for1------")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("------------")
	fmt.Println(i)
	fmt.Println("------------")
	fmt.Println("------for1------")
}
func for2() {
	var i=10
	fmt.Println("------for2------")
	for i = 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("------------")
	fmt.Println(i)
	fmt.Println("------------")
	fmt.Println("------for2------")
}