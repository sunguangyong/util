package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// 读取文本文件
	filePath := "path/to/your/file.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("无法读取文本文件：%v", err)
	}

	// 将字节内容转换为字符串
	text := string(content)
	fmt.Println("文本内容：", text)
}