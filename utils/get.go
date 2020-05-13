package utils

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

func SendGet(urlName string, testData map[string]interface{}) map[string]interface{} {
	var resp,err = HttpRequest.Get(urlName,testData)
	if err != nil{
		fmt.Println(err)
	}
	body,_ := resp.Body()
	var mapData  map[string]interface{}
	//将json字符串转换为map
	if json.Unmarshal([]byte(string(body)),&mapData) == err{
		fmt.Println(err)
	}
	return mapData

}
