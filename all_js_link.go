package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var links []string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://google.com"),
		chromedp.Sleep(2*time.Second), // Wait for JS to execute
		chromedp.Evaluate(`Array.from(document.querySelectorAll('a')).map(a => a.href)`, &links),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, link := range links {
		fmt.Println(link)
	}
}
