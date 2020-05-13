package main

import "fmt"

func main() {
	var number =  []int{4,2,7,5,34,89}
	fmt.Println(number[5])
	fmt.Println(number[2:6])
	fmt.Println(append(number,23,45))
	fmt.Println(number)

	number = append(number,11)
	fmt.Println(number)

}
