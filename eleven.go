package main

import "fmt"
/**
Go 语言中 range 关键字用于 for
循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
 */
func main() {
	nums := []int{2,3,5,7}
	for i,v := range nums{
		if v == 3{
			fmt.Println("index:",i)
		}
	}

	kvs := map[string]string{
		"name":"diane",
		"age":"12",
		"sex":"male"}
	for k,v := range kvs{
		fmt.Println(k,v)
	}
}
