package main

import (
	"fofa_format/ico"
	"fofa_format/osDo"
	"fofa_format/search"
	"github.com/fatih/color"
	"os"
)

func main() {

	ico.Ico()
	//初始化保存结果的end.txt
	osDo.InitializeFile("end.txt")
	osDo.InitializeFile("http.txt")
	osDo.InitializeFile("https.txt")
	//如果有config.yaml则直接读取，如果没有则调用写config的函数写入
	_, err := os.ReadFile("config.yaml")
	if err != nil {
		fofaURL := osDo.PromptInput("请输入 FOFA 站点 URL (直接回车默认 https://fofa.info)")
		email := osDo.PromptInput("请输入你的邮箱号")
		key := osDo.PromptInput("请输入你的Key")
		osDo.WriterInfo(email, key, fofaURL)
	}
	email, key, fofaURL, _, _ := osDo.ReadInfo()
	search.SetAPIBase(fofaURL)
	color.Red("[email:%s key:%s url:%s]", email, key, search.GetAPIBase())
	//qbase64 := search.GetAllUrl()
	base64AllList, _ := search.GetAllBase64(email, key)
	//fmt.Println(urls)
	//var a = []string{"YXBwPSLnlYXmjbfpgJotVFBsdXMiICYmIHJlZ2lvbj0iR3Vhbmdkb25nIg==", "YXBwPSLnlYXmjbfpgJotVFBsdXMiICYmIHJlZ2lvbj0iQmVpamluZyI="}
	search.FofaSearch(email, key, base64AllList)
	//color.Red("finish!!!")
}
