package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
)

func extractURLS(url string) []string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	urls := make([]string, 0)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		urls = append(urls, url)
	})
	return urls
}

/*
func main() {
	urls := extractURLS("http://www.indiabix.com/")
	for _, url := range urls {
		fmt.Printf("New URL : %s\n", url)
	}
}
*/
