package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

var cli *http.Client
var once sync.Once

func init() {
	once.Do(func() {
		cli = new(http.Client)
	})
}

func PostJson(req, body interface{}, headers map[string]string, host, path string) error {
	bs, err := json.Marshal(req)
	if err != nil {
		return err
	}
	fmt.Println("....req...", string(bs))
	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s%s", host, path), bytes.NewBuffer(bs))
	httpReq.Header.Add("Content-Type", "application/json")

	for k, v := range headers {
		httpReq.Header.Add(k, v)
	}

	if err != nil {
		return nil
	}
	resp, err := cli.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	fmt.Println("....resp...", string(bodyBytes))
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		return err
	}
	return nil
}

func GetJson(body interface{}, headers map[string]string, host, path string) error {
	url := fmt.Sprintf("%s/%s", host, path)
	fmt.Println("....req...", url)
	httpReq, err := http.NewRequest("GET", url, nil)
	httpReq.Header.Add("Content-Type", "application/json")

	for k, v := range headers {
		httpReq.Header.Add(k, v)
	}

	if err != nil {
		return nil
	}
	resp, err := cli.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	fmt.Println("....resp...", string(bodyBytes))
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		fmt.Println("err:------------------", err)
		return err
	}
	return nil
}

func PostFormData(url string) {
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("/Users/sunguangyong/Desktop/新疆env.txt")
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("", filepath.Base("/Users/sunguangyong/Desktop/新疆env.txt"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func Get() {
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
