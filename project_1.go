package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>Keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {

	var s Sitemapindex
	var n News

	response, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(response.Body)
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		response, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(response.Body)
		xml.Unmarshal(bytes, &n)
	}

}
