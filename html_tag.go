package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	url := "https://aparat.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}
	traverse(doc)
}

// Recursively traverse the HTML nodes and print their tag names
func traverse(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Println("Tag:", n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c)
	}
}
