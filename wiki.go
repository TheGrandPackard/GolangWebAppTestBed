package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html", "templates/login.html", "templates/logout.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *WikiPage) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9_ ]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
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

		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if session.Values["username"] == nil {
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

func main() {

	//getDHCPLeases()

	// Initialize Database Connection
	initDB()

	// Wiki Pages
	http.HandleFunc("/view/", checkLogin(makeHandler(viewHandler)))
	http.HandleFunc("/edit/", checkLogin(makeHandler(editHandler)))
	http.HandleFunc("/save/", checkLogin(makeHandler(saveHandler)))

	// Session Management Pages
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)

	// Static Files
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	http.ListenAndServe(":8080", nil)
}
