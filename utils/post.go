package utils

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

func SendPost(urlName string,testData map[string]interface{} ) string{
	var resp,err  = HttpRequest.Post(urlName,testData)
	if err != nil {
		fmt.Println(err)

	}
	body, _ := resp.Body()
	return string(body)



}
