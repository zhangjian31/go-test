package main

import(
	"fmt"
)

func  main() {
	//声明并初始化
	// var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
	//数组声明
	var arr [3][4]int
	//数组初始化
	//倒数第二行的 } 必须要有逗号，因为最后一行的 } 不能单独一行
	// 或者写成{8, 9, 10, 11}} 
	 arr =[3][4]int{
		 {0, 1, 2, 3} ,   /*  第一行索引为 0 */
		 {4, 5, 6, 7} ,   /*  第二行索引为 1 */
		 {8, 9, 10, 11}, /* 第三行索引为 2 */
		}
 for  i := 0; i < 3; i++ {
      for j := 0; j < 4; j++ {
         fmt.Printf("arr[%d][%d] = %d\n", i,j, arr[i][j] )
      }
   }

}