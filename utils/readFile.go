package utils

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"

)
//在外部被引用时，函数名必须大写
func FileMain(projectPath string,fileName string, index int) map[string]interface{} {
	//读取文件
	//projectPath, err := os.Getwd()
	//projectPath = "C:\\Users\\ASUS\\go\\src\\awesomeProject"

	filePath := projectPath + "/testdata/" +fileName
	result,err := os.Open(filePath)
	if err != nil{
		fmt.Println(err)
	}
	content, _ := ioutil.ReadAll(result)
	st := gjson.Parse(string(content)).Array()[index].String()
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(st),&data);err == nil{

	}
	return data





}