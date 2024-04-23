package wallet

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 传入一个字符串，然后写入文件
func writeLog(content string) {
	directory := "~/.lotus"

	// 判断目录是否存在
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		// 目录不存在，创建目录
		_ = os.Mkdir(directory, 0755)
	}
	logPath := path.Join(directory, "wallet.log")
	// 打开文件，如果文件不存在则创建
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 将字符串写入文件
	currentTime := time.Now()
	s := fmt.Sprintf("%s%s\n", currentTime.String(), content)
	_, err = file.WriteString(s)
	if err != nil {
		fmt.Println("无法写入文件:", err)
		return
	}
	fmt.Println("字符串已写入文件")
}
