package main

import (
	"net/http"
	"io/ioutil"
	"flag"
	"strings"
	"fmt"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getContentFromResp(resp http.Response) string {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleErr(err)
	return string(body)
}

func HttpGet(urlStr string) string {
	resp, err := http.Get(urlStr)
	handleErr(err)
	return getContentFromResp(*resp)
}

func HttpPost(urlStr string) string {
	resp, err := http.Post(
		urlStr,
		"application/x-www-form-urlencoded",
		strings.NewReader("name=instant"))
	handleErr(err)
	return getContentFromResp(*resp)
}

func HttpPut(urlStr string) string {
	client := &http.Client{}
	request, err := http.NewRequest("PUT", urlStr,
		strings.NewReader("name=instant"))
	handleErr(err)
	resp, err := client.Do(request)
	handleErr(err)
	return getContentFromResp(*resp)
}

var HttpFuncMap = make(map[string]func(string) string)

func main() {
	HttpFuncMap["GET"] = HttpGet
	HttpFuncMap["POST"] = HttpPost
	HttpFuncMap["PUT"] = HttpPut

	fMethod := flag.String("method", "get", "method for request")
	fTargetUrl := flag.String("url", "", "url for request")
	flag.Parse()

	*fMethod = strings.ToUpper(*fMethod)

	// check
	if HttpFuncMap[*fMethod] == nil {
		panic("method should in GET, POST, PUT")
	}
	if !strings.HasPrefix(*fTargetUrl, "http") {
		panic("url should starts with http")
	}

	var result = HttpFuncMap[*fMethod](*fTargetUrl)
	fmt.Println(result)
}
