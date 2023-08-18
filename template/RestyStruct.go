package template

import (
	"encoding/json"
	"fmt"
	"fofa_format/osDo"
	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
	"time"
)

func RestyStruct[T any](url string) T {
	// 创建一个新的 Resty 客户端
	client := resty.New()

	var data T
	var success bool
	for i := 0; i < 5; i++ {
		// 发送 GET 请求并获取响应
		resp, err := client.R().Get(url)
		if err != nil {
			fmt.Printf("第 %d 次重试\n", i+1)
			fmt.Println("HTTP请求错误:", err)
			color.Red("请确认你的网络正常，或者稍后重新尝试")
			//osDo.Sc()
			time.Sleep(5 * time.Second)
			continue // 继续下一次循环重试
		}

		// 解析 JSON 数据
		err = json.Unmarshal(resp.Body(), &data)
		if err != nil {
			fmt.Println("解析JSON错误:", err)
			osDo.Sc()
			panic(err) // 继续下一次循环重试
		}

		// 请求成功，标记为成功，并跳出循环
		success = true
		break
	}

	if !success {
		color.Red("重试多次仍未成功获取数据")
		osDo.Sc()
		panic("error")
	}

	return data
}
