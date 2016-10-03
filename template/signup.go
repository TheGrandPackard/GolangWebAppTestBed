package template

import (
	"html/template"
	"log"
	"net/http"

	"github.com/thegrandpackard/wiki/database"
	"github.com/thegrandpackard/wiki/session"
)

var signupTitle = "Sign Up"
var signupJSTop = template.HTML("<link href=\"../css/signup.css\" rel=\"stylesheet\">")

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if checkLoggedIn(w, r) == true {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		sess := session.GetSession(r)

		if username == "" {
			LoadTemplates()
			type Index struct {
				Site *Site
			}
			resp := Index{
				Site: SiteInit(r),
			}

			resp.Site.Title = signupTitle
			resp.Site.JsTopPage = signupJSTop
			resp.Site.Error = "No Username Specified"

			err := contentTemplate["signup"].Execute(w, resp)
			if err != nil {
				log.Printf("Error executing login: %s", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else if password == "" {
			LoadTemplates()
			type Index struct {
				Site *Site
			}
			resp := Index{
				Site: SiteInit(r),
			}

			resp.Site.Title = signupTitle
			resp.Site.JsTopPage = signupJSTop
			resp.Site.Error = "No Password Specified"

			err := contentTemplate["signup"].Execute(w, resp)
			if err != nil {
				log.Printf("Error executing login: %s", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else if len(password) < 8 {
			LoadTemplates()
			type Index struct {
				Site *Site
			}
			resp := Index{
				Site: SiteInit(r),
			}

			resp.Site.Title = signupTitle
			resp.Site.JsTopPage = signupJSTop
			resp.Site.Error = "Password must be at least 8 characters long"

			err := contentTemplate["signup"].Execute(w, resp)
			if err != nil {
				log.Printf("Error executing login: %s", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {
			user := &database.User{
				Username: username,
				Password: password,
			}

			if err := user.Save(); err != nil {
				LoadTemplates()
				type Index struct {
					Site *Site
				}
				resp := Index{
					Site: SiteInit(r),
				}

				resp.Site.Title = signupTitle
				resp.Site.JsTopPage = signupJSTop
				resp.Site.Error = err.Error()

				err = contentTemplate["signup"].Execute(w, resp)
				if err != nil {
					log.Printf("Error executing login: %s", err.Error())
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else {
				sess.Values["user"] = user
				if err = sess.Save(r, w); err != nil {
					log.Printf("Error saving Session Values: %s", err)
				}
				http.Redirect(w, r, "/view/home", http.StatusFound)
			}
		}
	} else /* Get */ {
		LoadTemplates()
		type Index struct {
			Site *Site
		}
		resp := Index{
			Site: SiteInit(r),
		}

		resp.Site.Title = signupTitle
		resp.Site.JsTopPage = signupJSTop

		err := contentTemplate["signup"].Execute(w, resp)
		if err != nil {
			log.Printf("Error executing login: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
