package main

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

func main() {
	data  :=  make(map[string]interface{})
	data["name"] = "daine"
	data["age"] = "26"
	resp, _ := HttpRequest.Get("http://localhost:8889/v1/getDemo",data)
	body, _ := resp.Body()
	fmt.Println(string(body))
}