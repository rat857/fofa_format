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
	osDo.InitializeFile()
	//如果有config.yaml则直接读取，如果没有则调用写config的函数写入
	_, err := os.ReadFile("config.yaml")
	if err != nil {
		//fmt.Println("")
		email := osDo.PromptInput("请输入你的邮箱号")
		key := osDo.PromptInput("请输入你的Key")
		osDo.WriterInfo(email, key)
	}
	email, key := osDo.ReadInfo()
	//
	color.Red("[email:%s key:%s]", email, key)
	qbase64 := search.GetAllUrl()
	//fmt.Println(qbase64)
	search.FofaSearch(email, key, qbase64)
	color.Red("finish!!!")
}
