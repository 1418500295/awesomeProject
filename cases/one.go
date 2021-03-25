package main

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	//创建计数器
	wg         = sync.WaitGroup{}
	num        = 1000 //设置并发数量
	ok_num     = 0   //初始化请求成功的数量
	time_list  []int
)
type Request struct {
	url  string
	data map[string]interface{}
}

func main() {
	start_time := (time.Now().UnixNano())
	fmt.Printf("开始时间：%v \n",start_time)
	do(num)
	end_time := time.Now().UnixNano()
	fmt.Printf("结束时间：%v \n", end_time)
	fmt.Println("成功的数量：", ok_num)
	fmt.Printf("失败的数量：%v \n", num-ok_num)
	fmt.Printf("总耗时：%v 秒 \n", (end_time - start_time)/1e9)
	sum := 0
	for _,i := range time_list {
		sum = sum + i
	}
	fmt.Printf("平均响应时间是:%.2f 秒 \n",float64(sum)/float64(num)/1000)
	fmt.Printf("QPS：%.2f",float64(num)/(float64(sum)/float64(num)/1000))

}

func do(num int) {
	//设置多核CPU运行
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(num) //初始化计数器
	for i := 0; i < num; i++ {
		go httpSend()

	}
	//主协程阻塞子协程执行完毕
	wg.Wait()

}

func httpSend() string {
	//计数器减一
	defer wg.Done()
	var request Request
	request.url = "http:///api/v1//"
	request.data = make(map[string]interface{})
	request.data["TypeId"] = 
	request.data["AgentId"] = 
	sTime := time.Now().UnixNano()/1e6
	resp, _ := HttpRequest.Get(request.url, request.data)
	body, _ := resp.Body()
	eTime := time.Now().UnixNano()/1e6
	use_time := (eTime - sTime)
	time_list = append(time_list, int(use_time))
	fmt.Println(string(body))
	if strings.Contains(string(body), "0") {
		ok_num += 1
	}
	return string(body)

}
