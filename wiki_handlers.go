package main

import (
	"errors"
	"html/template"
	"log"
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

	type Index struct {
		Site *Site
		Page *WikiPage
	}

	resp := Index{
		Site: SiteInit(),
		Page: p,
	}

	resp.Site.Title = p.Title

	err = contentTemplate["view"].Execute(w, resp)
	if err != nil {
		log.Printf("Error executing view: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := getWikiPage(title)
	if err != nil {
		p = &WikiPage{Title: strings.Replace(title, "_", " ", -1)}
	}

	type Index struct {
		Site *Site
		Page *WikiPage
	}

	resp := Index{
		Site: SiteInit(),
		Page: p,
	}

	resp.Site.Title = "Edit | " + p.Title
	resp.Site.JsTopPage = template.HTML("<script src=\"//cdn.tinymce.com/4/tinymce.min.js\"></script>")

	err = contentTemplate["edit"].Execute(w, resp)
	if err != nil {
		log.Printf("Error executing edit: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	id, _ := strconv.Atoi(r.FormValue("id"))
	p := &WikiPage{ID: id, Title: title, Body: template.HTML(body)}

	err := p.savePage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func pagesHandler(w http.ResponseWriter, r *http.Request) {
	p := getWikiPages()

	type Index struct {
		Site  *Site
		Pages WikiPages
	}

	resp := Index{
		Site:  SiteInit(),
		Pages: p,
	}

	resp.Site.Title = "Pages"
	resp.Site.JsTopPage = template.HTML("<link href=\"../css/pages.css\" rel=\"stylesheet\">")

	err := contentTemplate["pages"].Execute(w, resp)
	if err != nil {
		log.Printf("Error executing pages: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
