package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	
	if resp.StatusCode >= 399 {
		return "", errors.New(string(resp.StatusCode))
	}

	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		return "", errors.New("not html")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return "" , err
	}
	
	return string(body), nil
}