package main

import (
	"fmt"
	"net/http"

	"github.com/rizki4106/tribun-news-unofficial-api/router"
)


func main(){
	r := router.Router()
	fmt.Println("Running at port 4000")
	http.ListenAndServe(":4000", r)
}