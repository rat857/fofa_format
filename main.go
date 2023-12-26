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
		//fmt.Println("")
		email := osDo.PromptInput("请输入你的邮箱号")
		key := osDo.PromptInput("请输入你的Key")
		osDo.WriterInfo(email, key)
	}
	email, key, _, _ := osDo.ReadInfo()
	//
	color.Red("[email:%s key:%s]", email, key)
	//qbase64 := search.GetAllUrl()
	base64AllList, input := search.GetAllBase64(email, key)
	//fmt.Println(urls)
	//var a = []string{"YXBwPSLnlYXmjbfpgJotVFBsdXMiICYmIHJlZ2lvbj0iR3Vhbmdkb25nIg==", "YXBwPSLnlYXmjbfpgJotVFBsdXMiICYmIHJlZ2lvbj0iQmVpamluZyI="}
	search.FofaSearch(email, key, base64AllList, input)
	//color.Red("finish!!!")
}
