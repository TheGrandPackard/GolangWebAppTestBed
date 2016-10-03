package template

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thegrandpackard/wiki/database"
	"github.com/thegrandpackard/wiki/session"
)

// AddTemplateRoutes -- Handles HTML Template Routes
func AddTemplateRoutes(router *mux.Router) {

	// Register Structs
	gob.Register(database.User{})

	// Load Templates
	err := LoadTemplates()
	if err != nil {
		log.Printf("Error Loading Templates: %s", err)
	}

	// Wiki Pages
	router.HandleFunc("/", wikiHandler)
	router.HandleFunc("/view/{page}", wikiHandler)
	router.HandleFunc("/edit/{page}", wikiHandler)
	router.HandleFunc("/save/{page}", wikiHandler)
	router.HandleFunc("/pages", wikiPagesHandler)

	// User Management Pages
	router.HandleFunc("/login", loginHandler).Methods("GET", "POST")
	router.HandleFunc("/logout", logoutHandler).Methods("GET")
	router.HandleFunc("/signup", signupHandler).Methods("GET", "POST")
	router.HandleFunc("/users", usersHandler).Methods("GET")
	router.HandleFunc("/users", userEditHandler).Methods("POST")
	router.HandleFunc("/users/edit/{userID}", userEditPageHandler).Methods("GET", "POST")
}

func checkLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	session := session.GetSession(r)
	if session == nil || session.Values["user"] == nil {
		if r.URL.Path != "/login" {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
		return false
	}
	return true
}

func checkManageUsers(r *http.Request) bool {
	sess := session.GetSession(r)
	if sess == nil || sess.Values["user"] == nil {
		return false
	}
	user := session.GetSessionUser(sess)
	return user.CanManageUsers()
}

func handleError(w http.ResponseWriter, r *http.Request, message string, code int) {
	LoadTemplates()
	type Index struct {
		Site *Site
	}
	resp := Index{
		Site: SiteInit(r),
	}
	resp.Site.Title = "Error"
	resp.Site.Error = message

	w.WriteHeader(code)
	err := contentTemplate["error"].Execute(w, resp)
	if err != nil {
		log.Printf("Error rendering error page: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
