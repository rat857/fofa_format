package template

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
)

func RestyStruct[T any](url string) T {
	// 创建一个新的 Resty 客户端
	client := resty.New()

	// 发送 GET 请求并获取响应
	resp, err := client.R().Get(url)
	if err != nil {
		fmt.Println("HTTP请求错误:", err)
		color.Red("请确认你的网络正常")
		panic(err)
	}
	// 定义一个变量来存储解析后的 JSON 数据
	var data T
	// 解析 JSON 数据
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		fmt.Println("解析JSON错误:", err)
		panic(err)
	}
	return data
}
