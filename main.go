package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) <  2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURl := os.Args[1]	

	getHTML(baseURl)

	fmt.Printf("starting crawl of: %s...\n", baseURl)

	htmlBody, err := getHTML(baseURl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(htmlBody)
}