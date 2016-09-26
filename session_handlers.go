package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

// User Struct
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Session max time (1800 seconds = 30 minutes)
var sessionMaxLifetime = 1800

var store = sessions.NewCookieStore([]byte("thereoncewasapersonwhoneershallbenamed"))

func loginHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {

		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		if session.Values["username"] == nil {
			//Database lookup of the user
			user, err := getDatabaseUser(username)

			//TODO: Database lookup and hash/salt the password before comparison

			if err != nil || password == user.Password {
				log.Printf("User %s logged in successfully", username)
				// Set some session values.
				session.Values["username"] = username

				// TODO: Session Expiry
				//session.Values["expiry"] = time.Now().Add(time.Duration(sessionMaxLifetime) * time.Second)

				// Save it before we write to the response/return from the handler.
				if err = session.Save(r, w); err != nil {
					log.Printf("Error saving Session Values: %s", err)
				}

				http.Redirect(w, r, "/view/home", http.StatusFound)
			} else {
				log.Printf("User %s login failed: %s", username, err)
				http.Redirect(w, r, "/login", http.StatusFound)
			}
		}
	} else /* Get */ {
		if username := session.Values["username"]; username != nil {
			log.Printf("User %s already logged in", username)
			http.Redirect(w, r, "/view/home", http.StatusFound)
		} else {
			type Index struct {
				Site *Site
			}

			resp := Index{
				Site: SiteInit(),
			}

			resp.Site.Title = "Login"
			resp.Site.JsTopPage = template.HTML("<link href=\"../css/login.css\" rel=\"stylesheet\">")

			err := contentTemplate["login"].Execute(w, resp)
			if err != nil {
				log.Printf("Error executing login: %s", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := session.Values["username"]
	// Clear the session values
	session.Values = make(map[interface{}]interface{})
	// Save it before we write to the response/return from the handler.
	if err = session.Save(r, w); err != nil {
		log.Printf("Error saving Session Values: %s", err)
	}
	log.Printf("User %s logged out", username)

	type Index struct {
		Site *Site
	}

	resp := Index{
		Site: SiteInit(),
	}

	resp.Site.Title = "Login"

	err = contentTemplate["logout"].Execute(w, resp)
	if err != nil {
		log.Printf("Error executing logout: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getDatabaseUser(u string) (User, error) {
	var id int
	var username string
	var password string

	row := db.QueryRow("SELECT id, username, password FROM wiki.user WHERE username LIKE '" + u + "'")
	if row == nil {
		return User{}, errors.New("No User: " + u)
	}

	if err := row.Scan(&id, &username, &password); err == nil {
		checkDBError(err)
	}

	user := User{
		ID:       id,
		Username: username,
		Password: password,
	}

	return user, nil
}
