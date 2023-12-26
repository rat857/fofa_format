package osDo

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"net/url"
	"os"
	"strings"
)

// 写入文件（把List按行写入）
func WriteListTxt(resList []string, filName string) { //第一个参数是写入的url的list,第二个参数生成的文件名
	//resList = RemoveDuplicateIPs(resList) // 给ip去重,因为一个ip上可能有几个端口，几个端口上可能都有漏洞
	//创建文件
	file, err := os.OpenFile(filName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		Sc()
		panic(err)
	}
	defer file.Close()

	//创建 bufio.Writer
	writer := bufio.NewWriter(file)

	//循环写入文件
	for _, line := range resList {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println(err)
			Sc()
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
func InitializeFile(filName string) {
	file, err := os.OpenFile(filName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		color.Red("%s文件初始化失败", filName)
		Sc()
	}
	defer file.Close()
}

// 生成http和https的切片
func Fenkai(urls []string) ([]string, []string) {

	var httpURLs []string
	var httpsURLs []string

	for _, url := range urls {
		if strings.HasPrefix(url, "http://") {
			httpURL := strings.TrimPrefix(url, "http://")
			httpURLs = append(httpURLs, httpURL)
		} else if strings.HasPrefix(url, "https://") {
			httpsURL := strings.TrimPrefix(url, "https://")
			httpsURLs = append(httpsURLs, httpsURL)
		}
	}
	return httpURLs, httpsURLs
	/*	//fmt.Println("HTTP URLs:")
		for _, url := range httpURLs {
			fmt.Println(url)
		}

		//fmt.Println("\nHTTPS URLs:
		for _, url := range httpsURLs {
			fmt.Println(url)
		}*/
}

// format生成的切片
func Format(result [][]string) []string {
	var newSlice []string

	for _, goodhost := range result {
		if strings.HasPrefix(goodhost[0], "https://") {
			newSlice = append(newSlice, goodhost[0])
		} else {
			newSlice = append(newSlice, "http://"+goodhost[0])
		}
	}
	newSlice = RemoveDuplicateIPs(newSlice) //给ip去重,因为一个ip上可能有几个端口，几个端口上可能都有漏洞
	return newSlice
}
