package search

import (
	"encoding/base64"
	"fmt"
	"fofa_format/osDo"
	"fofa_format/template"
	"github.com/fatih/color"
	"github.com/go-the-way/exl"
	"os"
	"time"
)

func FofaSearch(email, key string, qbase64List []string, excelName string) {
	color.Yellow("开始查询具体资产...")
	fields := "host,protocol,country_name,region,domain,os,server,title,lastupdatetime,cname" // 提取哪个字段
	size := "10000"                                                                           // 提取多少条数据
	getSize := 0                                                                              //统计共有多少条数据
	var allResults = make([][]string, 0)
	//var forExcel = make([][]string,0)
	for _, qbase64 := range qbase64List {
		time.Sleep(3 * time.Second)
		url := "https://fofa.info/api/v1/search/all?" + "email=" + email + "&key=" + key + "&qbase64=" + qbase64 + "&fields=" + fields + "&size=" + size
		//fmt.Println(url)
		data := template.RestyStruct[template.Data](url)
		if data.Error {
			color.Red("sorry %s", data.Errmsg)
			osDo.Sc()
			return
		} else {
			GoodResultList := osDo.Format(data.Results)
			for _, result := range data.Results {
				allResults = append(allResults, result)
			}
			decoded, _ := base64.StdEncoding.DecodeString(qbase64)
			color.Green("正在保存%s的结果，去重后得到%d条", string(decoded), len(GoodResultList))
			getSize = getSize + len(GoodResultList)
			osDo.WriteListTxt(GoodResultList, "end.txt")
			httpUrls, httpsUrls := osDo.Fenkai(GoodResultList)
			osDo.WriteListTxt(httpUrls, "http.txt")
			osDo.WriteListTxt(httpsUrls, "https.txt")
		}
	}
	//fmt.Println(allResults)
	color.Red("去重后一共得到%d条资产", getSize)

	color.Red("Will be written to Excel")
	//for excel
	var tes = make([]template.ResultExcel, 0)
	for _, result := range allResults {
		var aa template.ResultExcel
		aa.Host = result[0]
		aa.Protocol = result[1]
		aa.CountryName = result[2]
		aa.Region = result[3]
		aa.Domain = result[4]
		aa.OS = result[5]
		aa.Server = result[6]
		aa.Title = result[7]
		aa.Lastupdatetime = result[8]
		aa.Cname = result[9]
		tes = append(tes, aa)
	}
	w := exl.NewWriter()
	if err := w.Write(excelName, tes); err != nil {
		fmt.Println(err)
		return
	}
	if err := w.SaveTo(excelName + ".xlsx"); err != nil {
		fmt.Println(err)
		return
	}
	color.HiBlue("Excel name is: %s.xlsx", excelName)
	//end
	color.Red("finish!!!")
	osDo.Sc()
	os.Exit(1)
}
