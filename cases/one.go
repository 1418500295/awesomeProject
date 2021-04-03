package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"github.com/tidwall/gjson"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	//创建计数器
	wg         = sync.WaitGroup{}
	num        = 100 //设置并发数量
	ok_num     = 0   //初始化请求成功的数量
	time_list  []int
	userId = 
	//agentId = 
	pwd = ""
)
type Request struct {
	url  string
	data map[string]interface{}
}

//MD5 加密
func Md5ToString(content string)  string{
	h:= md5.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))

}
//排序
func SortData(data map[string]interface{}) string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("异常信息：",err)
		}
	}()
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	sortStr := ""
	for _,i := range keys {
		sortStr += i+"="+fmt.Sprintf("%v",data[i])+"&"

	}
	return strings.TrimRight(sortStr,"&")
}
//获取时间戳
func SetTs() string {
	return strconv.FormatInt(time.Now().Unix()*1000,10)
}


func main() {
	start_time := time.Now().UnixNano()
	fmt.Printf("开始时间：%v \n",start_time)
	do(num)
	end_time := time.Now().UnixNano()
	fmt.Printf("结束时间：%v \n", end_time)
	fmt.Println("成功的数量：", ok_num)
	fmt.Printf("失败的数量：%v \n", num-ok_num)
	fmt.Printf("总耗时：%.2f 秒 \n", float64(end_time - start_time)/1e9)
	sum := 0
	for _,i := range time_list {
		sum = sum + i
	}
	fmt.Printf("平均响应时间是:%.2f 秒 \n",float64(sum)/float64(num)/1000)
	fmt.Printf("QPS：%.2f",float64(num)/(float64(sum)/float64(num)/1000))

}

func do(num int) {
	//设置多核CPU运行
	//runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(num) //初始化计数器
	for i := 0; i < num; i++ {
		go httpSend()

	}
	//主协程阻塞子协程执行完毕
	wg.Wait()

}
//发送请求
func httpSend() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误信息: ",err)
		}
	}()
	//计数器减一
	defer wg.Done()
	var request Request
	request.url = "http://"
	data := make(map[string]interface{})
	//request.data["TypeId"] = 89
	//request.data["AgentId"] = 14926056
	params := make(map[string]interface{})
	params["UserId"] = userId
	params["ts"] = SetTs()
	params["Key"] = Md5ToString(pwd)
	sign := strings.ToUpper(Md5ToString(SortData(params)))
	data["UserId"] = userId
	data["ts"] = SetTs()
	data["sign"] = sign
	sTime := time.Now().UnixNano()/1e6
	resp, _ := HttpRequest.Get(request.url, data)
	if resp.StatusCode() == 200 {
		body, _ := resp.Body()
		//fmt.Println(string(body))
		eTime := time.Now().UnixNano()/1e6
		use_time := (eTime - sTime)
		//fmt.Println(use_time)
		time_list = append(time_list, int(use_time))
		jsonData := gjson.Parse(string(body)).Map()
		if strings.Contains(jsonData["Data"].String(),"UserID") &&
			jsonData["Msg"].String() == "success"{
			ok_num += 1
		}
	}else {
		panic("请求异常")
	}


}
