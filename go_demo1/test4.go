package main
import "fmt"
func main() {
	var a int =10
	var b int =21
	if a!=10 &&b!=20{
		fmt.Println("!a!b")
	}else if a!=10 &&b==20{
		fmt.Println("!ab")
	}else if a==10 &&b!=20{
		fmt.Println("a!b")
	}else if a==10 &&b==20{
		fmt.Println("ab")
	}

	if a==10 {
		if b==20 {
			fmt.Println("right")
		}else{
			fmt.Println("wrong")
		}
	}

	switch(a){
	case 10:
		fmt.Println("10")
	case 20:
		fmt.Println("20")
	default:
		fmt.Println("sb")
	}
}