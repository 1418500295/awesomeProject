
package common

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"net/url"
)

func DoGet(urlName string,filename string,caseIndex int) string {

	jsonData := ReadFile(filename,caseIndex)
	for s := range jsonData {
		if s == "Pwd"{
			jsonData["Key"] = Md5ToString(jsonData[s].(string))
		}
	}
	jsonData["sign"] = GetSign(filename,caseIndex)
	jsonData["ts"] = SetTs()
	delete(jsonData,"Pwd")
	//判断map中的值类型，因为map解析会将int值自动转为float64，并将值(一般为interface类型)转为string或int
	for k,v := range jsonData {
		switch v.(type) {
		case float64:
			jsonData[k] = int(v.(float64))
		case string:
			jsonData[k] = v.(string)
		}
	}
	res,err := HttpRequest.Get(GetUrl(urlName),jsonData);if err != nil {
		fmt.Println(err)
	}
	body, _ := res.Body()
	resStr,err1 := url.QueryUnescape(string(body));if err1 != nil {
		fmt.Println(err1)
	}
	return resStr


}
