package template

import (
	"html/template"
	"log"
	"net/http"

	"github.com/thegrandpackard/wiki/database"
	"github.com/thegrandpackard/wiki/session"
)

// LoginHandler for HTTP
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	session := session.GetSession(r)

	if r.Method == http.MethodPost {

		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		if session.Values["username"] == nil {
			//Database lookup of the user
			user, err := database.GetUser(username)

			//TODO: Database lookup and hash/salt the password before comparison
			if err != nil && password == user.Password {
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

			LoadTemplates()
			type Index struct {
				Site *Site
			}

			resp := Index{
				Site: SiteInit(r),
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
