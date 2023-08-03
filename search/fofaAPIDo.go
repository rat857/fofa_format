package search

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"fofa_format/osDo"
	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
	"time"
)

type Data struct {
	Errmsg          string   `json:"errmsg"`
	Error           bool     `json:"error"`
	ConsumedFpoint  int      `json:"consumed_fpoint"`
	RequiredFpoints int      `json:"required_fpoints"`
	Size            int      `json:"size"`
	Page            int      `json:"page"`
	Mode            string   `json:"mode"`
	Query           string   `json:"query"`
	Results         []string `json:"results"`
}

func FofaSearch(email, key string, qbase64List []string) {
	color.Red("正在调用api查找...")
	fields := "host" // 提取哪个字段
	size := "10000"  // 提取多少条数据
	for _, qbase64 := range qbase64List {
		time.Sleep(3 * time.Second)
		url := "https://fofa.info/api/v1/search/all?" + "email=" + email + "&key=" + key + "&qbase64=" + qbase64 + "&fields=" + fields + "&size=" + size
		//fmt.Println(url)
		// 创建一个新的 Resty 客户端
		client := resty.New()

		// 发送 GET 请求并获取响应
		resp, err := client.R().Get(url)
		if err != nil {
			fmt.Println("HTTP请求错误:", err)
			return
		}

		// 定义一个变量来存储解析后的 JSON 数据
		var data Data

		// 解析 JSON 数据
		err = json.Unmarshal(resp.Body(), &data)
		if err != nil {
			fmt.Println("解析JSON错误:", err)
			return
		}
		if data.Error {
			color.Red("sorry %s", data.Errmsg)
			return
		} else {
			//return
			//fmt.Println("good")
			GoodResultList := osDo.Format(data.Results)
			decoded, _ := base64.StdEncoding.DecodeString(qbase64)
			color.Green("正在保存%s的结果，去重后得到%d条", string(decoded), len(GoodResultList))
			osDo.WriteListTxt(GoodResultList)
		}

	}
}
