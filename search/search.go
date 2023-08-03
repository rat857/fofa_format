package search

import (
	"encoding/base64"
	"fofa_format/osDo"
	"github.com/fatih/color"
	"github.com/go-rod/rod"
	"strings"
)

// 对输入进行base64编码，并且拼接为fofaUrl并返回
func SearchInfo() string {
	input := osDo.PromptInput("查询语法")
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	url := "https://fofa.info/result?qbase64=" + encoded
	return url
}

// 调用"github.com/go-rod/rod"获取所有url的qbase64位
func GetAllUrl() []string {
	page := rod.New().NoDefaultDevice().MustConnect().MustPage(SearchInfo())
	color.Red("这一步可能会比较久，请耐心等待。。。")
	page.MustWindowFullscreen()//最大化窗口
	//下面这个不接受任何返回的目的是等元素加载完毕 ：获取多个元素的方法的名称都以 MustElements 或 Elements 作为前缀。
	//单元素选择器和多元素选择器之间的一个关键区别是，单元素选择器会等待元素出现。 如果一个多元素选择器没有找到任何东西，他会立即返回一个空列表。
	color.Red("正在进行第一层，获取country信息")
	page.MustElementX("//*[@id=\"__layout\"]/div/div[2]/div/div[2]/div[1]/div[3]/div[5]/div/div/ul/li[1]/div/div[1]/div[1]/a").MustProperty("href")
	c_list := page.MustElementsX("//*[@id=\"__layout\"]/div/div[2]/div/div[2]/div[1]/div[3]/div[5]/div/div/ul/li/div/div[1]/div[1]/a")
	var countryList = make([]string, 0)
	for _, i2 := range c_list {
		//fmt.Println(reflect.TypeOf(i2.MustProperty("href").Str()))
		countryList = append(countryList, i2.MustProperty("href").Str())
		bas := strings.Split(i2.MustProperty("href").Str(), "qbase64=")
		decoded, _ := base64.StdEncoding.DecodeString(bas[1])
		color.Blue(string(decoded))
		color.Red(i2.MustProperty("href").Str())
	}
	//fmt.Println(countryList)
	page.Close()
	color.Red("正在进行第二层，获取不同country的Server排名信息")
	var urlList = make([]string, 0)
	var qbase64 = make([]string, 0)
	for _, countryUrl := range countryList {
		page := rod.New().NoDefaultDevice().MustConnect().MustPage(countryUrl)
		page.MustWindowFullscreen()
		page.MustElementX("//*[@id=\"__layout\"]/div/div[2]/div/div[2]/div[1]/div[3]/div[12]/div/div/ul/li[1]/div/a")
		r_list := page.MustElementsX("//*[@id=\"__layout\"]/div/div[2]/div/div[2]/div[1]/div[3]/div[12]/div/div/ul/li/div/a")
		for _, i2 := range r_list {
			urlList = append(urlList, i2.MustProperty("href").Str())
			bas := strings.Split(i2.MustProperty("href").Str(), "qbase64=")
			qbase64 = append(qbase64, bas[1])
			decoded, _ := base64.StdEncoding.DecodeString(bas[1])
			color.Green(string(decoded))
			color.Red(i2.MustProperty("href").Str())
		}
		page.Close()
	}
	//fmt.Println(urlList)
	color.Blue("共得到%d条语法", len(urlList))
	return qbase64
}
