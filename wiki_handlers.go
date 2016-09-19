package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := getWikiPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := getWikiPage(title)
	if err != nil {
		p = &WikiPage{Title: strings.Replace(title, "_", " ", -1)}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	id, _ := strconv.Atoi(r.FormValue("id"))
	p := &WikiPage{ID: id, Title: title, Body: body}

	err := p.savePage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func pagesHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "pages.html", getWikiPages())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
