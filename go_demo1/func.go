package main
import "fmt"
func main() {
	// var result = max(10,20)
	// fmt.Printf("max=%d\n",result)

	// var_1,var_2 :=dis2("aaa",1)
	// fmt.Printf("%s %d\n",var_1,var_2)

	 // a,b :=swap("aaa","bbb")
	 // fmt.Printf("%s   %s\n",a,b)

	var a string ="aaa"
	var b string = "bbb"
	fmt.Printf("before--> %s   %s\n",a,b)
	swap2(&a,&b)
	fmt.Printf("after--->%s   %s\n",a,b)
}

func swap(x,y string)(string,string){
	return y,x
}

func swap2(x *string ,y *string){
	var temp string 
	temp = *x
	*x = *y
	*y = temp
}

func dis(str string,a int) {
	fmt.Printf("%s %d\n",str,a)
}

//多个返回值用括号
func dis2(str1 string,a1 int) (string,int){
	return "str="+str1, a1+a1+a1
}

func max(num1,num2 int) int{
	var result int
	if num1>num2 {
		result =num1
	}else {
		result=num2
	}
	return result
}