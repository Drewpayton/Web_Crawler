package main

import (
	"net/url"
)

func normalizeURL(fullURL string) (string, error){
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return "", err
	}

	normalizedString := parsedURL.Host + parsedURL.Path
	if string(normalizedString[len(normalizedString) - 1]) == "/" {
		normalizedString = normalizedString[:len(normalizedString) - 1]
	}


	return normalizedString, nil
}