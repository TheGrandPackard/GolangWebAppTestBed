package template

import (
	"log"
	"net/http"

	"github.com/thegrandpackard/wiki/session"
)

// LogoutHandler For HTTP
func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	sess := session.GetSession(r)

	username := sess.Values["username"]
	// Invalidate the session
	sess.Options.MaxAge = -1
	// Save it before we write to the response/return from the handler.
	if err := sess.Save(r, w); err != nil {
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

	resp.Site.Title = "Logged Out"

	err := contentTemplate["logout"].Execute(w, resp)
	if err != nil {
		log.Printf("Error executing logout: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
