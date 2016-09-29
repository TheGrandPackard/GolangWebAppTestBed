package template

import (
	"html/template"
	"net/http"

	"github.com/thegrandpackard/wiki/database"
)

var usersJsTop = template.HTML("<link href=\"../css/users.css\" rel=\"stylesheet\">")

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if checkLoggedIn(r) == false {
		handleError(w, r, "Page Not Found", http.StatusNotFound)
		return
	}

	u, err := database.GetUsers()
	if err != nil {
		handleError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	LoadTemplates()
	type Index struct {
		Site  *Site
		Users database.Users
	}
	resp := Index{
		Site:  SiteInit(r),
		Users: u,
	}
	resp.Site.Title = "Users"
	resp.Site.JsTopPage = usersJsTop

	if err = contentTemplate["users"].Execute(w, resp); err != nil {
		handleError(w, r, err.Error(), http.StatusInternalServerError)
	}
}
