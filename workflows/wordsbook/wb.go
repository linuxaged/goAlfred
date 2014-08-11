package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	// "os"
)

func main() {
	var url_yd *url.URL
	url_yd, err := url.Parse("http://fanyi.youdao.com")
	if err != nil {
		panic(err.Error())
	}
	url_yd.Path += "/openapi.do?"
	param := url.Values{}
	param.Add("keyfrom", "rediffuse")
	param.Add("key", "1698275791")
	param.Add("type", "data")
	param.Add("doctype", "json")
	param.Add("version", "1.1")
	param.Add("q", "开水")
	url_yd.RawQuery = param.Encode()
	// if len(os.Args) == 2 {
	result, err := http.Get(url_yd.String())
	if err != nil {
		panic(err.Error())
	} else {
		defer result.Body.Close()
		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			panic(err.Error())
		} else {
			fmt.Println(body)
		}
	}
	// }

}
