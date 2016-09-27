package template

import (
	"html/template"
	"log"
	"net/http"

	"github.com/thegrandpackard/wiki/database"
	"github.com/thegrandpackard/wiki/session"
)

var loginTitle = "Login"
var loginJSTop = template.HTML("<link href=\"../css/login.css\" rel=\"stylesheet\">")

// LoginHandler for HTTP
func loginHandler(w http.ResponseWriter, r *http.Request) {
	sess := session.GetSession(r)

	if r.Method == http.MethodPost {

		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		if sess.Values["username"] == nil /* Username set in session  */ {
			//Database lookup of the user
			user, err := database.GetUser(username)

			//TODO: Database lookup and hash/salt the password before comparison
			if user != nil && password == user.Password {
				log.Printf("User %s logged in successfully", username)

				// Set username in session
				sess.Values["username"] = username
				// Session Expiry
				sess.Options.MaxAge = session.MaxLifetime
				// Save it before we write to the response/return from the handler.
				if err = sess.Save(r, w); err != nil {
					log.Printf("Error saving Session Values: %s", err)
				}
				http.Redirect(w, r, "/view/home", http.StatusFound)
			} else /* Invalid Username or Password */ {
				log.Printf("User %s login failed: %s", username, err)

				LoadTemplates()
				type Index struct {
					Site *Site
				}
				resp := Index{
					Site: SiteInit(r),
				}

				resp.Site.Title = loginTitle
				resp.Site.JsTopPage = loginJSTop
				resp.Site.Error = "Invalid Username or Password. Please Try again."

				err := contentTemplate["login"].Execute(w, resp)
				if err != nil {
					log.Printf("Error executing login: %s", err.Error())
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	} else /* Get */ {
		if username := sess.Values["username"]; username != nil {
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

			resp.Site.Title = loginTitle
			resp.Site.JsTopPage = loginJSTop

			err := contentTemplate["login"].Execute(w, resp)
			if err != nil {
				log.Printf("Error executing login: %s", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
