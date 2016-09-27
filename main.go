package main

import (
	"net/http"

	"github.com/thegrandpackard/wiki/database"
	"github.com/thegrandpackard/wiki/template"
)

func main() {

	// Initialize Database Connection
	database.InitDB()

	// Template Handlers
	template.MapTemplateHandlers()

	// Static File Handlers
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	http.ListenAndServe(":8080", nil)
}
