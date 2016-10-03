package template

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/thegrandpackard/wiki/database"
	"github.com/thegrandpackard/wiki/session"
)

var loginTitle = "Login"
var loginJSTop = template.HTML("<link href=\"../css/login.css\" rel=\"stylesheet\">")

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if checkLoggedIn(w, r) == true {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	sess := session.GetSession(r)

	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		if user := session.GetSessionUser(sess); user == nil /* User set in session  */ {
			//Database lookup of the user
			user, err := database.GetUser(username)

			//TODO: Database lookup and hash/salt the password before comparison
			if user != nil && password == user.Password && user.Enabled {
				log.Printf("User %s logged in successfully", username)
				now := time.Now()
				user.DateLastLogin = &now
				if err = user.Save(); err != nil {
					log.Printf("Error Saving User: %s", err.Error())
				}

				// Set username in session
				sess.Values["user"] = user
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

				if err := contentTemplate["login"].Execute(w, resp); err != nil {
					handleError(w, r, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	} else /* Get */ {
		if user := session.GetSessionUser(sess); user != nil {
			log.Printf("User %s already logged in", user.Username)
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

			if err := contentTemplate["login"].Execute(w, resp); err != nil {
				handleError(w, r, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
