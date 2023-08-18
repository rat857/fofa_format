package search

import (
	"encoding/base64"
	"fofa_format/osDo"
	"fofa_format/template"
	"github.com/fatih/color"
	"os"
	"time"
)

func FofaSearch(email, key string, qbase64List []string) {
	color.Yellow("开始查询具体资产...")
	fields := "host" // 提取哪个字段
	size := "10000"  // 提取多少条数据
	getSize := 0     //统计共有多少条数据
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
			decoded, _ := base64.StdEncoding.DecodeString(qbase64)
			color.Green("正在保存%s的结果，去重后得到%d条", string(decoded), len(GoodResultList))
			getSize = getSize + len(GoodResultList)
			osDo.WriteListTxt(GoodResultList, "end.txt")
			httpUrls, httpsUrls := osDo.Fenkai(GoodResultList)
			osDo.WriteListTxt(httpUrls, "http.txt")
			osDo.WriteListTxt(httpsUrls, "https.txt")
		}
	}
	color.Red("去重后一共得到%d条资产", getSize)
	color.Red("finish!!!")
	osDo.Sc()
	os.Exit(1)
}
