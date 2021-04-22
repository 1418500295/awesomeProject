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




func ReadFile(file string,caseIndex int) map[string]interface{} {

	path, _ := os.Getwd()
	dir := filepath.Dir(path)
	bytes, _ := ioutil.ReadFile(dir + file)
	var jsonData []map[string]interface{}
	err := json.Unmarshal(bytes, &jsonData)
	if err != nil {
		fmt.Println(err)
	}
	return jsonData[caseIndex]

}

//读取properties文件
func GetUrl(urlName string) string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获到异常: ",err)
		}
	}()
	path, err:= os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	dir := filepath.Dir(path)
	files,err1 := os.Open(dir + "/test/"+PropertiesFile);if err1 != nil{
		fmt.Println(err1)
	}
	bytesStr, _ := ioutil.ReadAll(files)
	configStr := strings.Split(string(bytesStr),"\n")
	var endUrl string
	for _,i := range configStr {
		iSlice := strings.Split(strings.ReplaceAll(i,"\r",""),"=")
		if urlName == strings.Trim(iSlice[0]," "){
			url := strings.Trim(iSlice[1]," ")
			if !strings.HasPrefix(url,"/") {
				endUrl = Host + "/" +fmt.Sprintf("%v",url)
			}else if strings.HasPrefix(url,"/") {
				endUrl = Host + fmt.Sprintf("%v",url)
			}else {
				panic("请求地址格式不正确")
			}
		}
	}
	return endUrl
}
