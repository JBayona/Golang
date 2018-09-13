package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Second line is an error - _ -> error
	response, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(response.Body)
	stringBody := string(bytes)
	fmt.Println(stringBody)
	response.Body.Close()
}
