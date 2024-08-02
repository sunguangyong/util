package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func PostFormData(url string, fileName string, fileContent []byte) {
	method := "POST"
	// Mock file content (replace with your actual file content)
	// Prepare the payload
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	// Create form file field
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}
	// Write file content to part
	_, err = part.Write(fileContent)
	if err != nil {
		fmt.Println("Error writing file content:", err)
		return
	}
	// Close the writer
	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing writer:", err)
		return
	}
	// Create HTTP request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	// Set Content-Type header
	req.Header.Set("Content-Type", writer.FormDataContentType())
	// Send the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()
	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response:", string(body))
}

func Get() (body []byte, err error) {
	url := "http://172.10.40.63:30085/v1/image/2fGZtimUnvlR1uVaYiPmMkI8zMU/d4ddd8172ea9bcdfcde623d5a99ca11d"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	return body, nil
}

func main() {
	url := "http://121.91.171.184:8989/upload/mintklub/meta"
	fmt.Println("url === ", url)
	fineName := "/Users/sunguangyong/Desktop/test.txt"
	data, err := ioutil.ReadFile(fineName)
	if err != nil {
		fmt.Println("err ==== ", err)
	}
	fmt.Println("lllll", string(data))
	PostFormData(url, data)
}
