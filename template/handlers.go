package template

import (
	"log"
	"net/http"
	"regexp"

	"github.com/thegrandpackard/wiki/session"
)

// MapTemplateHandlers -- For main
func MapTemplateHandlers() {

	// Load Templates
	err := LoadTemplates()
	if err != nil {
		log.Printf("Error: %s", err)
	}

	// Wiki Pages
	http.HandleFunc("/view/", checkLogin(wikiMakeHandler(wikiViewHandler)))
	http.HandleFunc("/edit/", checkLogin(wikiMakeHandler(wikiEditHandler)))
	http.HandleFunc("/save/", checkLogin(wikiMakeHandler(wikiSaveHandler)))
	http.HandleFunc("/pages", checkLogin(wikiPagesHandler))

	// Session Management Pages
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/logout", LogoutHandler)
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9_ ]+)$")

func wikiMakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func checkLogin(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session := session.GetSession(r)

		if session == nil || session.Values["username"] == nil {
			log.Printf("No Session for User. Redirecting to Login")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// TODO: Session Expiry
		// if time.Now().After(session.Values["expiry"].(time.Time)) {
		// 	log.Printf("Session for User %s expired. Redirecting to Login", session.Values["username"])
		// 	http.Redirect(w, r, "/login", http.StatusFound)
		// 	return
		// }

		fn(w, r)
	}
}
