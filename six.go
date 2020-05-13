package main

import "fmt"

/**
数组使用
 */
func main() {
	var n = [5]int{1,2,7}
	fmt.Println(n[0])



	var threeim = [3][3]int{
		{2,4,6},
		{5,8,5},
		{2,6,4},
	}
	var a = threeim[2][1]

	fmt.Println(a)
	var i,j int
	for i =0; i<3;i++ {
		for j =0; j<3;j++{
			fmt.Printf("threeim[%d][%d] = %d\n",i,j, threeim[i],threeim[j])
		}
	}



}

