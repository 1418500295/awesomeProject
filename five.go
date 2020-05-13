package main

import "fmt"
/**
函数用法
 */
func max( num1, num2 int) int {
	var result int
	if (num1 < num2){
		result = num1
	}else {
		result = num2
	}
	return result

}
/**
go语言函数可以返回多个值
 */
func swap(x, y string) (string,string)  {
	return x, y

}
func main() {
	var res int
	res = max(2,4)
	fmt.Printf("结果是%d",res)

	a,c := swap("google","ibm")
	fmt.Println(a,c)
}