package main

import (
	"fmt"
	"github.com/unidoc/unipdf/v3/common/license"
	"os"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func init() {
	// https://cloud.unidoc.io/api-keys 获取 api key 地址
	apiKey := `6c5bfe7b03087eb12e2503472c7622577a974d5e783e2b22f3c0d9198e11ec2f`
	err := license.SetMeteredKey(apiKey)
	if err != nil {
		panic(err)
	}
}

func main() {

	inputPath := "/Users/sunguangyong/Desktop/IND231-IND236.pdf"

	err := outputPdfText(inputPath)
	if err != nil {
		fmt.Printf("1111Error: %v\n", err)
	}
}

// outputPdfText prints out contents of PDF file to stdout.
func outputPdfText(inputPath string) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	//defer f.Chmod()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	fmt.Printf("--------------------\n")
	fmt.Printf("PDF to text extraction:\n")
	fmt.Printf("--------------------\n")
	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return err
		}

		ex, err := extractor.New(page)
		if err != nil {
			return err
		}

		text, err := ex.ExtractText()
		if err != nil {
			return err
		}

		fmt.Println("------------------------------")
		fmt.Printf("Page %d:\n", pageNum)
		fmt.Printf("\"%s\"\n", text)
		fmt.Println("------------------------------")
	}

	return nil
}