package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"test/logrus"
)

func get_request(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello"))
	fmt.Println(req.Body)
}

func httpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	code := resp.StatusCode
	if code != 200 {
		msg := fmt.Sprintf("Url request failure: %v %v", url, code)
		return "", errors.New(msg)
	}

	resp_byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res := string(resp_byte)
	return res, nil

}

func Get(url string, headers map[string]string) (string, error) {
	client := &http.Client{}
	requestGet, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	for index, value := range headers {
		requestGet.Header.Add(index, value)
	}

	resp, err := client.Do(requestGet)
	if err != nil {
		return "", err
	}

	code := resp.StatusCode
	if code != 200 {
		msg := fmt.Sprintf("Url request failure: %v %v", url, code)
		return "", errors.New(msg)
	}

	resp_byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res := string(resp_byte)
	return res, nil
}

func Post(url string, data map[string]string, headers map[string]string) (string, error) {
	client := &http.Client{}

	jsonData, _ := json.Marshal(data)

	requestPost, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return "", err
	}

	for index, value := range headers {
		requestPost.Header.Add(index, value)
	}

	resp, err := client.Do(requestPost)
	if err != nil {
		return "", err
	}

	code := resp.StatusCode
	if code != 200 {
		msg := fmt.Sprintf("Url request failure: %v %v %v", url, code, jsonData)
		return "", errors.New(msg)
	}

	resp_byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	res := string(resp_byte)
	return res, nil
}

func main() {
	//http.HandleFunc("/", get_request)
	//http.ListenAndServe(":8001", nil)
	log := logrus.New()
	res, err := HttpPost("https://www.baidu.com", map[string]string{"Host": "www.baidu.com"}, map[string]string{"name": "www.baidu.com"})
	if err != nil {
		log.Warn(err)
	} else {
		fmt.Println(res)
	}
}
