package osDo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 引导用户输入 传入一个引导词，返回用户输入的内容
func PromptInput(key string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(key + ":")
	input, _ := reader.ReadString('\n')
	input = strings.ReplaceAll(input, "\n", "") //去除换行
	return input
}
