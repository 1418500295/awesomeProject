package utils

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

func GetUrl(projectPath string,urlFile string, uriName string) string {
	//var projectPath,err  = os.Getwd()
	//projectPath = "C:\\Users\\ASUS\\go\\src\\awesomeProject"

	filePath := projectPath + "/config/" + urlFile
	fmt.Println(filePath)
	urlData, _ := ioutil.ReadFile(filePath)
	//将字符串转换为Map
	js := gjson.Parse(string(urlData)).Map()
	//将gjson.Result转换为string
	var host = js["host"].Str
	var uri = js["uri"].Map()[uriName].Str
	var testUrl = host + uri
	return testUrl





}


