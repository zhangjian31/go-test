package main
import "fmt"
func main() {
var result = max(10,20)
	fmt.Printf("max=%d\n",result)
}

func max(num1,num2 int) int{
	var result int
	if num1>num2 {
		result =num1
	}else {
		result=num2
	}
}