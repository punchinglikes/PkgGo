package main

import (
	//"fmt"
	"log"
	"net/http"
	//"os"
	"strings"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)
//returns the number of packages found in a search querry (max 50)
func PackageCount(packageArg string) int{


	packageName := packageArg


	res, err := http.Get("https://pkg.go.dev/search?limit=50&q=" + packageName)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	itemNumber := doc.Find("div.SearchResults-summary").Find("h1").Text()

	var itemNumberCount []string = strings.Split(itemNumber, " ")
	intParsed, err := strconv.Atoi(itemNumberCount[7])
	return intParsed

	
}

//returns link of the top package result
func PackageLink(packageArg string) []string {
	packageName := packageArg


	res, err := http.Get("https://pkg.go.dev/search?limit=5&q=" + packageName)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}


	correctDiv := doc.Find("div.SearchSnippet").Find("div.SearchSnippet-headerContainer").Find("span.SearchSnippet-header-path").Text()

	var correctDivLinks []string = strings.Split(correctDiv, ")")

	var top5Links []string
	var count int = 6
	for i := range correctDivLinks {
		if count > 1 {
			newLink := strings.Replace(correctDivLinks[i], "(", "", 1)
			newLink = "https://pkg.go.dev/" + newLink
			top5Links = append(top5Links, newLink)

			count = count - 1
		} else {
			break
		}

	}
	//fmt.Println(top5Links[0] + ", " + top5Links[1] + ", " + top5Links[2] + ", " + top5Links[3] + ", " + top5Links[4]) 
	return top5Links
}