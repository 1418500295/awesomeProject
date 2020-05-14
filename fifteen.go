package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	//string转换为map
	a := "{\"name\":\"daine\",\"age\":\"26\"}"
	fmt.Println(reflect.TypeOf(a))
	var data = make(map[string]interface{})
	if err := json.Unmarshal([]byte(a),&data);err == nil{
		fmt.Println(err)
	}
	fmt.Println(data)

	var b = make(map[string]interface{})
	b["weight"] = "185"
	b["height"] = "174"
	jsonStr, _ := json.Marshal(b)
	fmt.Println(string(jsonStr))
	fmt.Println(reflect.TypeOf(string(jsonStr)))


}