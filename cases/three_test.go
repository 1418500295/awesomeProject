package main

import (
	"awesomeProject/utils"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)


func Test_login(t *testing.T) {
	var path, _ = os.Getwd()
	currentPath := strings.Split(path,"\\ca")
	projectPath := currentPath[0]
	var data  = utils.FileMain(projectPath,"getDemo.json",0)
	var testFile  = utils.GetUrl(projectPath,"url.json","getDemo")
	mapData := utils.SendGet(testFile,data)
	assert.Equal(t,"200",mapData["status"])
	fmt.Println(1)
	fmt.Println(os.Getwd())

}
func Test_post(t *testing.T) {
	var path, _ = os.Getwd()
	currentPath := strings.Split(path,"\\ca")
	projectPath := currentPath[0]
	var data  = utils.FileMain(projectPath,"getDemo.json",0)
	var testFile  = utils.GetUrl(projectPath,"url.json","getDemo")
	//rs := utils.SendGet(testFile,data)
	//fmt.Println(rs)
	resp, _ := HttpRequest.Get(testFile,data)
	body, _ := resp.Body()
	fmt.Println(string(body))


}
func Test_postParam(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "daine"
	data["sex"] = "male"
	resp, _ := HttpRequest.Post("http://localhost:8889/postSecond",data)
	body, _ := resp.Body()
	fmt.Println(string(body))

}

func Test_jsonPost(t *testing.T) {

	data := map[string]interface{}{
		"name":"james",
		"age":"23",
	}
	resp, _ := HttpRequest.JSON().Post("http://localhost:8889/postDemo",data)
	body, _ := resp.Body()
	fmt.Println(string(body))


}

func Test_withHeader(t *testing.T) {
	//header := make(map[string]string)
	//header["Cookie"] = "login=true"
	//resp, _ := HttpRequest.SetHeaders(header).Get("http://localhost:8889/get/with/cookies")
	cookie := make(map[string]string)
	cookie["login"] = "true"
	resp, _ := HttpRequest.SetCookies(cookie).Get("http://localhost:8889/get/with/cookies")
	body, _ := resp.Body()
	fmt.Println(string(body))

}

