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




