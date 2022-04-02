package model

type ResponseNews struct {
	TotalItem int `json:"total_item"`
	Error bool `json:"error"`
	ErrorMessage string `json:"error_message"`
	Items []Data `json:"items"`
}

type Data struct {
	Title string `json:"title"`
	Poster string `json:"poster"`
	Link string `json:"link"`
	Date string `json:"date"`
	ShortDescription string `json:"short_description"`
}