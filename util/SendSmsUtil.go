package util

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

//json数据解析
type Message struct {
	Message   string
	RequestId string
	BizId     string
	Code      string
}
var message Message //阿里云返回的json信息对应的类

func SendRemindSms(content string) {
	client, err := dysmsapi.NewClientWithAccessKey("ap-northeast-1", "<accessKeyId>", "<accessSecret>")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = "133986****"
	request.SignName = "你的signName"
	request.TemplateCode = "SMS_20*****"
	request.TemplateParam = "{\"content\":"+content+"}"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	//记得判断错误信息
	err = json.Unmarshal(response.GetHttpContentBytes(), &message)
	if message.Message != "OK" {
		//阿里云操作失败的错误处理
		return
	}
}

