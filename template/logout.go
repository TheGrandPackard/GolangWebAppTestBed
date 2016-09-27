package template

import (
	"log"
	"net/http"

	"github.com/thegrandpackard/wiki/session"
)

// LogoutHandler For HTTP
func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	session := session.GetSession(r)

	username := session.Values["username"]
	// Clear the session values
	session.Values = make(map[interface{}]interface{})
	// Save it before we write to the response/return from the handler.
	if err := session.Save(r, w); err != nil {
		log.Printf("Error saving Session Values: %s", err)
	}
	log.Printf("User %s logged out", username)

	LoadTemplates()
	type Index struct {
		Site *Site
	}

	resp := Index{
		Site: SiteInit(r),
	}

	resp.Site.Title = "Login"

	err := contentTemplate["logout"].Execute(w, resp)
	if err != nil {
		log.Printf("Error executing logout: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
