package search

import (
	"encoding/base64"
	"fmt"
	"fofa_format/osDo"
	"fofa_format/template"
	"github.com/fatih/color"
	"os"
	"strings"
	"time"
)

// 对输入进行base64编码，并返回base64
func SearchInfo() string {
	input := osDo.PromptInput("查询语法")
	//encoded := base64.StdEncoding.EncodeToString([]byte(input))
	//url := "https://fofa.info/result?qbase64=" + encoded
	return input
}

// 获取不同的server,返回不同的server的字符串切片
func GetServerBase64(email, key string) []string {
	input := SearchInfo()
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	var re = make([]string, 0)
	url := "https://fofa.info/api/v1/search/stats?fields=asset_type,country,server&qbase64=" + encoded + "&email=" + email + "&key=" + key
	//fmt.Println(url)
	color.Yellow("正在获取第一层----Server")
	data := template.RestyStruct[template.JuheInfo](url)
	if data.Error {
		color.Red(data.Errmsg)
		os.Exit(700)
	}
	for _, server := range data.Aggs.Server {
		//fmt.Println(input + " && " + server.Name)
		a := fmt.Sprintf(`%s && server=="%s"`, input, server.Name)
		color.Green(a)
		fmt.Println(base64.StdEncoding.EncodeToString([]byte(a)))
		re = append(re, a)
	}
	if len(re) == 0 {
		re = append(re, input)
	}
	color.Red("---------距离第二层结束预计还有%d秒,耐心等待---------", len(re)*10)
	return re
}
func GetAllBase64(email, key string) []string {
	getServerBase64 := GetServerBase64(email, key)
	color.Yellow("正在获取第二层----Country/Region")
	var base64AllList = make([]string, 0)
	for _, ServBase := range getServerBase64 {
		time.Sleep(10 * time.Second)

		url := "https://fofa.info/api/v1/search/stats?fields=asset_type,country,server&qbase64=" + base64.StdEncoding.EncodeToString([]byte(ServBase)) + "&email=" + email + "&key=" + key
		data := template.RestyStruct[template.JuheInfo](url)
		if data.Error {
			color.Red(data.Errmsg)
			os.Exit(700)
		}
		for _, country := range data.Aggs.Countries {
			if country.Count >= 10000 {
				for _, region := range country.Regions {
					switch region.Name {
					case "Unknown":
						//a := fmt.Sprintf(`%s && region="" && country="%s"`, ServBase, country.Name)

					default:
						decodedBytes, _ := base64.StdEncoding.DecodeString(region.Code)
						color.Blue(string(decodedBytes))
						base64AllList = append(base64AllList, region.Code)
					}
				}
			} else {
				if strings.Contains(country.Name, "香港") || strings.Contains(country.Name, "台湾") {
					a := fmt.Sprintf(`%s && region="%s"`, ServBase, country.NameCode)
					color.Red(a)
					base64AllList = append(base64AllList, base64.StdEncoding.EncodeToString([]byte(a)))
				} else {
					a := fmt.Sprintf(`%s && country="%s"`, ServBase, country.NameCode)
					color.Red(a)
					base64AllList = append(base64AllList, base64.StdEncoding.EncodeToString([]byte(a)))
				}
			}
		}
		/*un := fmt.Sprintf(`%s && region=""`, ServBase)
		color.Green(un)
		base64AllList = append(base64AllList, base64.StdEncoding.EncodeToString([]byte(un)))*/
	}
	color.Yellow("一共找到%d条语法", len(base64AllList))
	color.Red("---------距离结束预计还有%d秒---------", len(base64AllList)*3)
	return base64AllList
}
