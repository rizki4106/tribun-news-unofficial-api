package router

import (
	"github.com/gorilla/mux"
	"github.com/rizki4106/tribun-news-unofficial-api/controller"
)

func Router() * mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Root)
	r.HandleFunc("/{type}", controller.ParseData)
	// r.HandleFunc("/search/{param}", controller.Search)
	return r
}