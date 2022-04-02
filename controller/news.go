package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"github.com/rizki4106/tribun-news-unofficial-api/model"
)

const url = "https://www.tribunnews.com/bisnis"

/* This code is fetching data from the web page and storing it in a list of structs. */
func getData(endpoint string ) []model.Data {

	web_url := url + endpoint
	
	res, err := http.Get(web_url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	
	// 
	var response model.Data
	var listData = []model.Data{}
	
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	
	doc.Find(".pos_rel").Each(func(i int, s *goquery.Selection){
		title, _ := s.Find("a").Attr("title")
		link, _ := s.Find("a").Attr("href")
		poster, _ := s.Find("img").Attr("src")
		date := s.Find("time").Text()
		short_description := s.Find(".txt-oev-2").Text() 

		// remove /t and /n
		if len(short_description) > 0 {
			short_description = strings.ReplaceAll(short_description, "\t", "")
			short_description = strings.ReplaceAll(short_description, "\n", "")
		}

		// 
		response.Title =  title
		response.Link = link
		response.Poster  = poster
		response.Date = date
		response.ShortDescription = short_description

		if len(title) > 0 {
			listData = append(listData, response)
		}
	})

	return listData
}

// Method

/* This is the main function that will be called when the user access the endpoint /. */
func Root(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	data := getData("/")
	response := model.ResponseNews{len(data), false, "", data}
	json.NewEncoder(w).Encode(response)
}

/* This is the main function that will be called when the user access the endpoint /{type}. */
func ParseData(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	
	// get data
	data := getData("/" + param["type"])
	response := model.ResponseNews{len(data), false, "", data}
	json.NewEncoder(w).Encode(response)
}