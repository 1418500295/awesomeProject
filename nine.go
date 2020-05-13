package main

import "fmt"

func main() {
	var number   = [5]int{1,2,4,6,5}
	fmt.Println(number[3])

	//定义一个切片，为一个长度可变的数组
	var num  = make([]int,3,5)

	if(num == nil){
		fmt.Println("切片是空的")
	}

	printSlice(num)

}

func printSlice(x []int)  {
	fmt.Printf("len=%d,cap=%d,slice=%v",len(x),cap(x),x)
	
}
