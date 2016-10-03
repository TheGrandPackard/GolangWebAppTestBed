package template

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thegrandpackard/wiki/database"
)

var usersTitle = "Users"
var usersJsTop = template.HTML("<link href=\"/css/users.css\" rel=\"stylesheet\">")

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if checkLoggedIn(w, r) == false || checkManageUsers(r) == false {
		handleError(w, r, "Page Not Found", http.StatusNotFound)
		return
	}

	r.ParseForm()
	disabled, _ := strconv.ParseBool(r.FormValue("disabled"))

	u, err := database.GetUsers(disabled)
	if err != nil {
		handleError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	LoadTemplates()
	type Index struct {
		Site         *Site
		Users        database.Users
		ShowDisabled bool
	}
	resp := Index{
		Site:         SiteInit(r),
		Users:        u,
		ShowDisabled: disabled,
	}
	resp.Site.Title = usersTitle
	resp.Site.JsTopPage = usersJsTop

	if err = contentTemplate["users"].Execute(w, resp); err != nil {
		handleError(w, r, err.Error(), http.StatusInternalServerError)
	}
}

func userEditHandler(w http.ResponseWriter, r *http.Request) {
	if checkLoggedIn(w, r) == false || checkManageUsers(r) == false {
		handleError(w, r, "Page Not Found", http.StatusNotFound)
		return
	}

	r.ParseForm()
	action := r.FormValue("action")
	username := r.FormValue("username")

	u, err := database.GetUser(username)
	if err != nil {
		handleError(w, r, "Error getting User "+username+". "+err.Error(), http.StatusInternalServerError)
		return
	}

	if action == "disable" {
		u.Enabled = false
		if err = u.Save(); err != nil {
			handleError(w, r, "Error Disabling user: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else if action == "enable" {
		u.Enabled = true
		if err = u.Save(); err != nil {
			handleError(w, r, "Error enabling user: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else if action == "edit" {
		u.Enabled = r.FormValue("enabled") == "on"
		u.ManageUsers = r.FormValue("manageUsers") == "on"
		u.ManagePages = r.FormValue("managePages") == "on"
		if err = u.Save(); err != nil {
			handleError(w, r, "Error enabling user: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	users, err := database.GetUsers(false)
	if err != nil {
		handleError(w, r, "Error getting Users. "+err.Error(), http.StatusInternalServerError)
		return
	}

	LoadTemplates()
	type Index struct {
		Site         *Site
		Users        database.Users
		ShowDisabled bool
	}
	resp := Index{
		Site:         SiteInit(r),
		Users:        users,
		ShowDisabled: false,
	}
	resp.Site.Title = usersTitle
	resp.Site.JsTopPage = usersJsTop

	if err = contentTemplate["users"].Execute(w, resp); err != nil {
		handleError(w, r, err.Error(), http.StatusInternalServerError)
	}
}

func userEditPageHandler(w http.ResponseWriter, r *http.Request) {
	if checkLoggedIn(w, r) == false || checkManageUsers(r) == false {
		handleError(w, r, "Page Not Found", http.StatusNotFound)
		return
	}

	vars := mux.Vars(r)
	username := vars["userID"]

	u, err := database.GetUser(username)
	if err != nil {
		handleError(w, r, "Error getting User "+username+". "+err.Error(), http.StatusInternalServerError)
		return
	}

	LoadTemplates()
	type Index struct {
		Site *Site
		User *database.User
	}
	resp := Index{
		Site: SiteInit(r),
		User: u,
	}
	resp.Site.Title = usersTitle
	resp.Site.JsTopPage = usersJsTop

	if err = contentTemplate["user_edit"].Execute(w, resp); err != nil {
		handleError(w, r, err.Error(), http.StatusInternalServerError)
	}
}
