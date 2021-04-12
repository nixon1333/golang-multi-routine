package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func main() {
	parallelPtr := flag.Int("parallel", 10, "an int")
	flag.Parse()
	maxGoroutines := *parallelPtr
	urls := flag.Args()

	guard := make(chan struct{}, maxGoroutines)
	var wg sync.WaitGroup

	for i, element := range urls {
		wg.Add(1)
		element := element
		go func(n int) {
			guard <- struct{}{} // would block if guard channel is already filled
			defer wg.Done()
			ProcessUrl(element)
			<-guard
		}(i)
	}
	wg.Wait()
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
