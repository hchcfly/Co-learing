package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	urlStr := "http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName?theCityName=北京"

	client := http.Client{}

	response, err := client.Get(urlStr)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		fmt.Println(string(body))
	}
	fmt.Printf("response: %v\n", response)

}