package osDo

import (
	"fmt"
	"github.com/fatih/color"
)

// 用来卡住程序
func Sc() {
	// 等待用户按下任意键
	color.Cyan("Press Enter to exit...")
	fmt.Scanln()
	//os.Exit(0)
}
