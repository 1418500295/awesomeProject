package main

import (
	"fmt"
	"testing"
)


func Test_01(t *testing.T) {
	fmt.Println("3")

}


//age 为可选参数，可传可不传
func One(name string,age ...interface{})  {
	fmt.Println(name,age)
}


/*
匿名函数，函数返回值为一个函数，并且返回值的函数带一个int参数，并且有返回值为int类型
 */
//func One(name string) func(int) int  {
//	return func(age int) int {
//		fmt.Println(name,111)
//		return 111
//	}
//}
//
//func main() {
//	c := One("make")
//	c(12)
//	fmt.Println(c(12))



