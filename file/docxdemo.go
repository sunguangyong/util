package main

import (
	"baliance.com/gooxml/document"
	"fmt"
	"log"
)

func main() {
	path := "/Users/sunguangyong/Desktop/工作资料/数采资料/modbus.docx"
	doc, err := document.Open(path)
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	//doc.Paragraphs()得到包含文档所有的段落的切片

	for _, para := range doc.Paragraphs() {
		//run为每个段落相同格式的文字组成的片段
		//fmt.Println("-----------第", i, "段-------------")
		for _, run := range para.Runs() {
			//fmt.Print("\t-----------第", j, "格式片段-------------")
			fmt.Print(run.Text())
		}
		fmt.Println()
	}
}