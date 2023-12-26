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
func GetServerBase64(email, key string) ([]string, []string, string) {
	var base64AllList = make([]string, 0) //这个切片就是最终要传给查询用的base64的切片
	input := SearchInfo()
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	var re = make([]string, 0) //这个切片是带着server的查询语法，需要再次获得county信息的
	url := "https://fofa.info/api/v1/search/stats?fields=asset_type,country,server&qbase64=" + encoded + "&email=" + email + "&key=" + key
	//fmt.Println(url)
	color.Yellow("Start!!!")
	//请求速度过快时重试5次
	var data template.JuheInfo
	for i := 0; i < 5; i++ {
		data = template.RestyStruct[template.JuheInfo](url)
		if data.Error {
			if strings.Contains(data.Errmsg, "速度") {
				color.Red(data.Errmsg)
				fmt.Printf("第 %d 次重试\n", i+1)
				time.Sleep(10 * time.Second)
				continue
			} else {
				color.Red(data.Errmsg)
				osDo.Sc()
				os.Exit(700)
			}
		} else {
			break
		}
	}
	if data.Error {
		color.Red("重试多次仍显示请求速过快，请稍后重试")
		osDo.Sc()
		os.Exit(700)
	}
	//到此
	//如果总资产小于10000直接开始查询
	if data.Size < 10000 {
		color.Yellow("总资产数量%d,资产数量小于10000，开始查询", data.Size)
		color.Green(input)
		bbinput := base64.StdEncoding.EncodeToString([]byte(input))
		fmt.Println(bbinput)
		base64AllList = append(base64AllList, bbinput)
		FofaSearch(email, key, base64AllList, input)
	}
	//到此
	color.Yellow("总资产数量%d,资产数量大于10000，开始分批查询", data.Size)
	color.Yellow("正在步入第一层----Server")
	for _, server := range data.Aggs.Server {
		//判断每一个server的总资产，如果大于10000再进行County查询
		if server.Count > 10000 {
			a := fmt.Sprintf(`%s && server=="%s"`, input, server.Name)
			color.Green("%s,有%d条资产", a, server.Count)
			fmt.Println(base64.StdEncoding.EncodeToString([]byte(a)))
			re = append(re, a)
		} else {
			b := fmt.Sprintf(`%s && server=="%s"`, input, server.Name)
			color.Red("%s有%d条", b, server.Count)
			bbas := base64.StdEncoding.EncodeToString([]byte(b))
			fmt.Println(bbas)
			base64AllList = append(base64AllList, bbas)
		}
	}
	//可能没有server
	if len(re) == 0 {
		re = append(re, input)
	}
	color.Red("---------有%d条需要步入第二层，有%d条不用步入第二层,距离第二层结束预计还有%d秒,耐心等待---------", len(re), len(base64AllList), len(re)*10)
	return re, base64AllList, input
}

// 返回所有查询语法base64的切片
func GetAllBase64(email, key string) ([]string, string) {
	getServerBase64, base64AllList, input := GetServerBase64(email, key)
	color.Yellow("正在步入第二层----Country/Region")

	for _, ServBase := range getServerBase64 {
		time.Sleep(10 * time.Second)

		url := "https://fofa.info/api/v1/search/stats?fields=asset_type,country,server&qbase64=" + base64.StdEncoding.EncodeToString([]byte(ServBase)) + "&email=" + email + "&key=" + key
		//请求速度过快时重试5次
		var data template.JuheInfo
		for i := 0; i < 5; i++ {
			data = template.RestyStruct[template.JuheInfo](url)
			if data.Error {
				if strings.Contains(data.Errmsg, "速度") {
					color.Red(data.Errmsg)
					fmt.Printf("第 %d 次重试\n", i+1)
					time.Sleep(5 * time.Second)
					continue
				} else {
					color.Red(data.Errmsg)
					osDo.Sc()
					os.Exit(700)
				}
			} else {
				break
			}
		}
		if data.Error {
			color.Red("重试多次仍显示请求速过快，请稍后重试")
			osDo.Sc()
			os.Exit(700)
		}
		//到此
		for _, country := range data.Aggs.Countries {
			//如果国家里的资产数大于10000，再解析到城市
			if country.Count > 10000 {
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
			} else { //如果国家里的资产数字小于10000直接添加到base64AllList
				//这里的if判断和程序的运行关系不大
				if strings.Contains(country.Name, "香港") || strings.Contains(country.Name, "台湾") {
					a := fmt.Sprintf(`%s && region="%s"`, ServBase, country.NameCode)
					//a:=ServBase+" && region="+country.NameCode
					color.Red(a)
					base64AllList = append(base64AllList, base64.StdEncoding.EncodeToString([]byte(a)))
				} else {
					a := fmt.Sprintf(`%s && country="%s"`, ServBase, country.NameCode)
					//a:=ServBase+" && country="
					color.Red(a)
					base64AllList = append(base64AllList, base64.StdEncoding.EncodeToString([]byte(a)))
				}
				//为止
			}
		}
		/*un := fmt.Sprintf(`%s && region=""`, ServBase)
		color.Green(un)
		base64AllList = append(base64AllList, base64.StdEncoding.EncodeToString([]byte(un)))*/
	}
	//color.Yellow("一共找到%d条语法", len(base64AllList))
	color.Red("---------一共需要查询%d条语法,距离结束预计还有%d秒---------", len(base64AllList), len(base64AllList)*3)
	return base64AllList, input
}
