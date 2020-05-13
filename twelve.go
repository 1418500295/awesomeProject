package main

import "fmt"

/**
可以使用内建函数 make 也可以使用 map 关键字来定义 Map:
 */
func main() {
	//先声明map，再初始化map
	//var firstMap map[string]string
	//firstMap = make(map[string]string)

	//声明的同时并初始化map
	firstMap := make(map[string]string)
	firstMap["name"] = "diane"
	firstMap["age"] = "12"
	firstMap["sex"] = "male"

	//输出键，并根据键获取值
	for key := range firstMap{
		fmt.Println(key,"value是",firstMap[key])
	}

	//删除元素，根据key删除
	delete(firstMap,"name")
	for key := range firstMap{
		fmt.Println(key,"value是",firstMap[key])
	}



}