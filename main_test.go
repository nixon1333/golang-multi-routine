package main

import "testing"

type urlTestStr struct {
	arg, expectedUrl, expectedHash string
}

var addTests = []urlTestStr{
	urlTestStr{"twitter.com", "https://twitter.com/", "4fd9c9a2de4fbab8494b2f27d37f3118"},
	urlTestStr{"baroquemusiclibrary.com", "http://baroquemusiclibrary.com", "369dcc7299cc08bc6bed46ae321da030"},
	urlTestStr{"adjust.com", "https://www.adjust.com/", "e5d912ebcf54d8fb2775e386fc935d15"},
	urlTestStr{"yahoo.com", "https://www.yahoo.com/", "63d1fa3d64b889ba4fbf960134eafa2c"},
	urlTestStr{"yandex.com", "https://yandex.com/", "2e2222f93c0a668cc0907fb41518bda4"},
	urlTestStr{"google.com", "http://www.google.com/", "ff90821feeb2b02a33a6f9fc8e5f3fcd"},
	urlTestStr{"facebook.com", "https://www.facebook.com/", "e203e98e4c606735cf56db84a002fd22"},
}

func TestMD5(t *testing.T) {
	for _, test := range addTests {
		if output := GetMD5Hash(test.expectedUrl); output != test.expectedHash {
			t.Errorf("Output %q not equal to expected %q", output, test.expectedHash)
		}
	}
}

func TestMakeHttpCall(t *testing.T) {
	for _, test := range addTests {
		if output := MakeHttpCall(test.arg); output != test.expectedUrl {
			t.Errorf("Output %q not equal to expected %q", output, test.expectedUrl)
		}
	}
}

func TestProcessUrl(t *testing.T) {
	for _, test := range addTests {
		outputUrl, outputHash := ProcessUrl(test.arg)
		if outputUrl != test.expectedUrl {
			t.Errorf("Output %q not equal to expected %q", outputUrl, test.expectedUrl)
		}
		if outputHash != test.expectedHash {
			t.Errorf("Output %q not equal to expected %q", outputHash, test.expectedHash)
		}
	}
}
