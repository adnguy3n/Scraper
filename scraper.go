package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	const url = "http://httpbin.org/"

	collector := colly.NewCollector()

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting ", request.URL)
	})

	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("Got a response from ", response.Request.URL)
	})

	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("An Error occurred: ", err)
	})

	collector.OnHTML("a[href]", func(htmlElement *colly.HTMLElement) {
		link := htmlElement.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", htmlElement.Text, link)
	})

	collector.Visit(url)
}
