package main

import(
	"fmt"
)

func  main() {
	var arr = []int{1,2,3,4,5}
	var avg float32
	avg = getAverage(arr,5)
	fmt.Printf("平均值=%f\n",avg)
}

func getAverage(arr []int,size int) float32{

	var sum int
	var avg float32
	for i := 0; i < size; i++ {
		sum+=arr[i]
	}
	avg= float32(sum)/float32(size)
	return avg
}