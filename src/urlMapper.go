package main

import (
	"fmt"
	"io/ioutil"
)

func Redirect(name string) (string, bool) {
	// Alpha Test: Hardcoded static redirect target lookup table
	if redirectUrl, ok := urlRedirectMap[name]; ok {
		return redirectUrl, true
	}
	return "", false
}

func LoadTXTFile(name string) (string, bool) {
	// Alpha Test: Hardcoded static filepath lookup table
	if responseTxt, ok := urlFilepathMap[name]; ok {
		filepath := fmt.Sprintf("res/%s", responseTxt)
		content, err := ioutil.ReadFile(filepath)
		if err != nil {
			fmt.Printf("Err: %s\n", err)
			return "", false
		}
		return string(content), true
	}
	return "", false
}
