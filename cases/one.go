package main

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"strings"
	"sync"
	"time"
)

var (
	//创建计数器
	wg = sync.WaitGroup{}
	num =  500
	usedTime = 0
	ok_num = 0
	fal_num = 0
	start_time = 0
	end_time = 0
)
//接口压测实例
func main(){
	runTimes(num)
	fmt.Println("成功的数量：",ok_num)
	fmt.Printf("失败的数量：%v \n",num - ok_num)
	fmt.Printf("总耗时：%v毫秒",end_time - start_time )

}

func runTimes(num int)  {
	start_time = int(time.Now().UnixNano() / 1e6)
	//设置计数器
	wg.Add(num)
	for i := 0;i <num; i++{
		go send()
	}
	end_time = int(time.Now().UnixNano()/1e6)
	//主协程等待子协程执行完毕
	wg.Wait()

	//time.Sleep(time.Second*3)

}

func httpSend()  {

	data  :=  make(map[string]interface{})
	data["coinid"] = "68"

	headers := make(map[string]string)
	headers["Cookie"] = ""
	resp, _ := HttpRequest.SetHeaders(headers).Post("",data)
	body, _ := resp.Body()
	fmt.Println(string(body))
	//if strings.Contains(string(body),"\"info\":\"success\"") == true{
	//	ok_num += 1
	//}
	//fal_num += 1


}

func send()  {
	//第一次调用，计数器减一
	defer wg.Done()
	req := HttpRequest.NewRequest()
	data := make(map[string]string)
	data["name"]= "make"
	resp, _ := req.Get("http://127.0.0.1:1234/api/test?name=make")
	body, _ := resp.Body()
	fmt.Println(string(body))
	if strings.Contains(string(body),"make")==true {
		ok_num += 1
	}



}
