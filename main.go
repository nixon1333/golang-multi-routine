package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// take input from cli
	// break those inputs to list of strings
	// support parallel keywords in cli
	// print those urls
	// make the md5 hash of the request > maybe redirected request
	// add unit tests
}

func ProcessUrl(urlText string) (string, string) {
	urlFullText := MakeHttpCall(urlText)
	hashedUrl := GetMD5Hash(urlFullText)
	fmt.Println(urlFullText, hashedUrl)
	return urlFullText, hashedUrl
}

func MakeHttpCall(urlText string) string {
    //validating the http in the url
	httpCheck := strings.Index(urlText, "http://")
	httpsCheck := strings.Index(urlText, "https://")
	requestUrl := urlText

	if httpCheck != 0 && httpsCheck != 0 {
		requestUrl = fmt.Sprintf("http://%s", urlText)
	}

	// making the request
	resp, err := http.Get(requestUrl)
	if err == nil {
		finalURL := resp.Request.URL.String()
		// replacing with the requested url
		urlText = finalURL
	}
	return urlText
}

func GetMD5Hash(text string) string {
	hashMaker := md5.New()
    hashMaker.Write([]byte(text))
	return hex.EncodeToString(hashMaker.Sum(nil))
}
