package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main(){
	var a  = make(map[string]interface{})
	a["name"] = "daine"
	a["sex"] = "male"
	a["age"] = 12
	data, _ := json.Marshal(a)
	fmt.Println(string(data))
	fmt.Println(reflect.TypeOf(string(data)))



}

















