package main

import (
	"strings"
	"fmt"
)

var(
	links map[string]bool = make(map[string]bool)
	home = "http://www.indiabix.com"
)

func filter(stringArray []string,fn func(string) bool) []string{
	var returnStringArray []string
	for _,stringValue := range stringArray {
		if fn(stringValue){
			returnStringArray = append(returnStringArray,stringValue)
		}
	}
	return returnStringArray
}
func filterbyPrefix(urls []string,prefix string) []string{
        prefixFilterFn := func() func(str string) bool {
		return func(str string) bool{
			return !strings.HasPrefix(str,prefix)
		}
	};
	return filter(urls,prefixFilterFn())
}

func filter_urls(urls []string) []string{
	urls=filterbyPrefix(urls,"/about")
	urls=filterbyPrefix(urls,"javascript")
	return urls
	
}

func getNonVisitedUrls(urls []string) []string {
	nonVisitedUrlFn := func() func(str string) bool {
		return func(str string) bool{
			return !links[str]
		}
	};
	return filter(urls, nonVisitedUrlFn())
}

func addUrlToVisitedUrls(url string){
	links[url]=true
	fmt.Println("Visited URL :",url)
}

func crawlLinks(urls []string) {
	// Filters url based on regex
	urls=filter_urls(urls)

	// Get non visted urls
	// This acts as atermination condition for the recursive function
	// When thre are no visited links, for loop does not make any further recursive calls and hence this function terminates
	urls = getNonVisitedUrls(urls)
	//fmt.Println("output of getNonVisitedUrls :")
	//fmt.Println(urls)	
	//fmt.Println("End op")
	// Loop over all the non visited urls and crawls the respective web pages as well to collect all the links
	for _,url := range urls{
		addUrlToVisitedUrls(url)
		// Construct url from uri
		url = home+url
		childUrls := extractURLS(url)
		crawlLinks(childUrls)			
	}
	return
}

func main(){
	// Make home url visited
        links[home] = true
	
	urls := extractURLS(home)
	crawlLinks(urls)
	
	fmt.Println("In memory visited links")
	for url,_ := range links{
		fmt.Println(url)
	}
	
}
