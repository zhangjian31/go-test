package main
import "fmt"
func main() {
	for i := 0; i < 3; i++ {
		fmt.Printf("i=%d\n",i)
	}
	var a int =5
	var b int =10
	for a<b {
		fmt.Printf("i=%d\n",a)
		a++
	}

    arr :=[6]string{"a","b","c","d","e"}
	for index,value := range arr{
		fmt.Printf("arr[%d]=%s\n",index,value)
	}

	var num int = 0
	for true{
		if num>10 {
			fmt.Printf("%d>10 finish\n",num)
			break
		}else{
			fmt.Printf("%d>10?\n",num)
		}
		num++
	}
}