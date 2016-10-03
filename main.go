package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/thegrandpackard/wiki/database"
	"github.com/thegrandpackard/wiki/rest"
	"github.com/thegrandpackard/wiki/template"
)

func main() {

	// Initialize Database Connection
	database.InitDB()

	// Mux Router
	router := mux.NewRouter().StrictSlash(true)

	// Template Routes
	template.AddTemplateRoutes(router)

	// REST API Routes
	rest.AddRESTRoutes(router)

	// Static File Routes
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	srv := &http.Server{
		Handler: router,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
