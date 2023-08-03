package osDo

import (
	"bufio"
	"github.com/fatih/color"
	"net/url"
	"os"
	"strings"
)

// 写入文件（把List按行写入）
func WriteListTxt(resList []string) { //第一个参数是写入的url的list,第二个参数生成的文件名
	//resList = RemoveDuplicateIPs(resList) // 给ip去重,因为一个ip上可能有几个端口，几个端口上可能都有漏洞
	//创建文件
	file, err := os.OpenFile("end.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//创建 bufio.Writer
	writer := bufio.NewWriter(file)

	//循环写入文件
	for _, line := range resList {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}

	//刷新缓存
	writer.Flush()
}

// 给IP去重http //128.0.0.1 890
func RemoveDuplicateIPs(urls []string) []string {
	encountered := make(map[string]bool)
	result := []string{}

	for _, urlStr := range urls {
		u, err := url.Parse(urlStr)
		if err != nil {
			continue
		}

		host := u.Hostname()
		if !encountered[host] {
			encountered[host] = true
			result = append(result, urlStr)
		}
	}

	return result
}

// 初始化存放结果的文件
func InitializeFile() {
	file, err := os.OpenFile("end.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		color.Red("文件初始化失败")
	}
	defer file.Close()
}

// format生成的切片
func Format(urls []string) []string {
	var newSlice []string

	for _, goodhost := range urls {
		if strings.HasPrefix(goodhost, "https://") {
			newSlice = append(newSlice, goodhost)
		} else {
			newSlice = append(newSlice, "http://"+goodhost)
		}
	}
	newSlice = RemoveDuplicateIPs(newSlice) //给ip去重,因为一个ip上可能有几个端口，几个端口上可能都有漏洞
	return newSlice
}
