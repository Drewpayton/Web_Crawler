package main

import (
	"net/url"
	"strings"
)

func normalizeURL(fullURL string) (string, error){
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return "", err
	}

	normalizedString := parsedURL.Host + parsedURL.Path
	
	normalizedString = strings.TrimSuffix(normalizedString, "/")


	return normalizedString, nil
}

