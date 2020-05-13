package main

import (
	"fmt"
)

func main() {
	var a  = 2
	var c  = 3
	fmt.Print(a +c )
	fmt.Println(a * 3)
	fmt.Println(a - c)
	fmt.Printf("a加c的结果是%d",a+c)

	if ( a > c){
		fmt.Println("a大于c")
	}else {
		fmt.Println("a小于c")
	}




}
