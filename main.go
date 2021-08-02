package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)


type Driver struct {
	Name	 string	//`json:"position"`
	Position string //`json:"position"`
	Car      string //`json:"car"`
	Points   string //`json:"points"`
}

var headers = map[string]string{
	"authority":                 "www.formula1.com",
	"cache-control":             "max-age=0",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36",
	"sec-fetch-dest":            "document",
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"sec-fetch-site":            "none",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-user":            "?1",
	"accept-language":           "en-US,en;q=0.9",
}

var client = &http.Client{}

type Ranking []Driver

func ScrapeTable() Ranking {
	link := "https://www.formula1.com/en/results.html/2021/drivers.html"
	var list Ranking
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		var Item Driver
		
		Item.Position =  string(s.Find("tr  td:nth-child(2)").Text())
		
		Item.Name = string(s.Find("tr  td:nth-child(3)").Text())

		Item.Car = string(s.Find("tr  td:nth-child(5)").Text())

		Item.Points = string(s.Find("tr  td:nth-child(6)").Text())
		
		fmt.Println(i)

		list = append(list, Item)

	})
	return list
}

func main(){
	for i , v := range ScrapeTable(){
		fmt.Println(i, "--", v)
	}
}