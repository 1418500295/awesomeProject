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
	
	
	
	c := "我将受到公开"
	//加密
	a := base64.StdEncoding.EncodeToString([]byte(c))
	fmt.Println(a)
	//解密
	d, _ := base64.StdEncoding.DecodeString(a)
	fmt.Println(string(d))




}
