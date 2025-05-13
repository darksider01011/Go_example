package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
)

func js(url string) []string {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var links []string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(2*time.Second), // Wait for JS to execute
		chromedp.Evaluate(`Array.from(document.querySelectorAll('a')).map(a => a.href)`, &links),
	)
	if err != nil {
		log.Fatal(err)
	}
	keys := make(map[string]bool)
	link := []string{}
	for _, li := range links {
		if !keys[li] {
			keys[li] = true
			link = append(link, li)
		}
	}

	for _, a := range link {
		fmt.Println(a)
	}

	return link
}

func parse(url string) []string {
	var parse_link []string
	c := colly.NewCollector()
	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasPrefix(link, "/") {
			link = url + link
		}
		//fmt.Println(link)
		parse_link = append(parse_link, link)
	})
	c.Visit(url)

	keys := make(map[string]bool)
	link := []string{}
	for _, li := range parse_link {
		if !keys[li] {
			keys[li] = true
			link = append(link, li)
		}
	}
	for _, a := range link {
		fmt.Println(a)
	}
	return link
}

func main() {
	uri := "https://yahoo.com/"
	a := parse(uri)
	for _, b := range a {
		parse(b)
	}
}
