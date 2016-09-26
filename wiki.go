package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var isProduction = true
var isTemplateLoaded = false
var contentTemplate map[string]*template.Template

// LoadTemplates
func LoadTemplates() (err error) {
	if isTemplateLoaded && isProduction {
		return
	}

	isTemplateLoaded = true

	contentTemplate = make(map[string]*template.Template)
	log.Println("Loading Templates")

	bData, err := ioutil.ReadFile("templates/_template.tpl")
	if err != nil {
		return err
	}

	contentTemplate["_template"], err = template.New("_template").Parse(string(bData))
	if err != nil {
		return err
	}

	bData, err = ioutil.ReadFile("templates/_header.tpl")
	if err != nil {
		return err
	}

	_, err = contentTemplate["_template"].New("header").Parse(string(bData))
	if err != nil {
		return err
	}

	paths := []string{
		"view",
		"edit",
		"login",
		"logout",
		"pages",
	}

	for _, path := range paths {
		contentTemplate[path], err = contentTemplate["_template"].Clone()
		if err != nil {
			return err
		}

		bData, err = ioutil.ReadFile("templates/" + path + ".tpl")
		if err != nil {
			return err
		}
		_, err = contentTemplate[path].New("content").Parse(string(bData))
		if err != nil {
			return err
		}
	}

	return nil
}

// Site Struct
type Site struct {
	Title           string
	UseHeader       bool
	JsTopPage       template.HTML
	JsBotPage       template.HTML
	Content         string
	LastSearchQuery string
}

// Site Init
func SiteInit() *Site {
	return &Site{
		Title:           "Index",
		JsTopPage:       "",
		Content:         "",
		JsBotPage:       "",
		LastSearchQuery: "",
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
	err := LoadTemplates()
	if err != nil {
		log.Printf("Error: %s", err)
	}

	// Wiki Pages
	http.HandleFunc("/view/", checkLogin(makeHandler(viewHandler)))
	http.HandleFunc("/edit/", checkLogin(makeHandler(editHandler)))
	http.HandleFunc("/save/", checkLogin(makeHandler(saveHandler)))
	http.HandleFunc("/pages", checkLogin(pagesHandler))

	// Session Management Pages
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)

	// Static Files
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	http.ListenAndServe(":8080", nil)
}
