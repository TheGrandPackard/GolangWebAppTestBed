package template

import (
	"log"
	"net/http"

	"github.com/thegrandpackard/wiki/session"
)

var logoutTitle = "Logged Out"

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if checkLoggedIn(w, r) == false {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	sess := session.GetSession(r)
	user := session.GetSessionUser(sess)

	// Invalidate the session
	sess.Options.MaxAge = -1
	// Save it before we write to the response/return from the handler.
	if err := sess.Save(r, w); err != nil {
		log.Printf("Error saving Session Values: %s", err)
	}
	log.Printf("User %s logged out", user.Username)

	LoadTemplates()
	type Index struct {
		Site *Site
	}

	resp := Index{
		Site: SiteInit(r),
	}

	resp.Site.Title = logoutTitle

	err := contentTemplate["logout"].Execute(w, resp)
	if err != nil {
		log.Printf("Error executing logout: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
