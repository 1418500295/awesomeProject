package main

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"strings"
	"time"
)

var (
	num =  10000
	usedTime = 0
	ok_num = 0
	fal_num = 0
)
//接口压测实例
func main(){
	start_time := time.Now().Format(time.UnixDate)
	runTimes(num)
	end_time := time.Now().Format(time.UnixDate)
	//usedTime := (end_time - start_time)
	fmt.Println(start_time)
	fmt.Println("成功的数量：",ok_num)
	fmt.Println("失败的数量：",num - ok_num)
	fmt.Println("总耗时：",start_time,end_time)

}
func runTimes(num int)  {
	for i := 0;i <num; i++{
		go httpSend()
	}
	time.Sleep(time.Second*1)
}

func httpSend()  {

	 data  :=  make(map[string]interface{})
	 data["name"] = "daine"
	 data["age"] = "26"
	 resp, _ := HttpRequest.Get("http://localhost:8889/v1/getDemo",data)
	 body, _ := resp.Body()
	 fmt.Println(string(body))
	 if strings.Contains(string(body),"\"info\":\"success\"") == true{
	 	ok_num += 1
	 }
	 fal_num += 1


}

















