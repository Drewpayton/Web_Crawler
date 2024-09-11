package main

import (
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	htmlNode, err := html.Parse(htmlReader)

	if err != nil {
		return []string{}, err
	}

	anchorTags := []string{}

	var TraverseHTML func(node *html.Node)
	TraverseHTML = func(node *html.Node) {
		if node == nil {
			return 
		}

		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					if strings.Contains(attr.Val, "https:") {
						anchorTags = append(anchorTags, attr.Val)
					}else {
						anchorTags = append(anchorTags, rawBaseURL + attr.Val)
					}
				}
			}
		}

		TraverseHTML(node.FirstChild)
		TraverseHTML(node.NextSibling)
	}

	TraverseHTML(htmlNode)

	
	return anchorTags, nil
}