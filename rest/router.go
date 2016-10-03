package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thegrandpackard/wiki/database"
)

// AddRESTRoutes -- Handles REST routes
func AddRESTRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/user", usersGet).Methods("GET")
	router.HandleFunc("/api/v1/user", userPost).Methods("POST")
	router.HandleFunc("/api/v1/user/{userID}", userGet).Methods("GET")
	router.HandleFunc("/api/v1/user/{userID}", userPut).Methods("PUT")
	router.HandleFunc("/api/v1/user/{userID}", userDelete).Methods("DELETE")
}

func usersGet(w http.ResponseWriter, r *http.Request) {
	userlist, err := database.GetUsers(false)

	for idx := range userlist {
		userlist[idx].Password = ""
	}

	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(userlist); err != nil {
		panic(err)
	}
}

func userPost(w http.ResponseWriter, r *http.Request) {
}

func userGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	user, err := database.GetUser(userID)

	if err != nil {
		panic(err)
	}

	user.Password = ""

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func userPut(w http.ResponseWriter, r *http.Request) {
}

func userShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	fmt.Fprintln(w, "User show:", userID)
}

func userDelete(w http.ResponseWriter, r *http.Request) {
}
