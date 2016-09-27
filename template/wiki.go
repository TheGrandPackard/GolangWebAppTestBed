package template

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/thegrandpackard/wiki/database"
)

func wikiViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := database.GetWikiPage(title)
	if p == nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	LoadTemplates()
	type Index struct {
		Site *Site
		Page *database.WikiPage
	}
	resp := Index{
		Site: SiteInit(r),
		Page: p,
	}

	resp.Site.Title = p.Title

	err = contentTemplate["view"].Execute(w, resp)
	if err != nil {
		log.Printf("Error executing view: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func wikiEditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := database.GetWikiPage(title)
	if err != nil {
		p = &database.WikiPage{Title: strings.Replace(title, "_", " ", -1)}
	}

	LoadTemplates()
	type Index struct {
		Site *Site
		Page *database.WikiPage
	}
	resp := Index{
		Site: SiteInit(r),
		Page: p,
	}

	resp.Site.Title = "Edit | " + p.Title
	resp.Site.JsTopPage = template.HTML("<script src=\"//cdn.tinymce.com/4/tinymce.min.js\"></script>")

	err = contentTemplate["edit"].Execute(w, resp)
	if err != nil {
		log.Printf("Error executing edit: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func wikiSaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	id, _ := strconv.Atoi(r.FormValue("id"))
	p := &database.WikiPage{ID: id, Title: title, Body: body}

	err := p.SavePage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
	return
}

func wikiPagesHandler(w http.ResponseWriter, r *http.Request) {
	p, err := database.GetWikiPages()
	if err != nil {
		log.Printf("Error executing pages: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	LoadTemplates()
	type Index struct {
		Site  *Site
		Pages database.WikiPages
	}
	resp := Index{
		Site:  SiteInit(r),
		Pages: p,
	}

	resp.Site.Title = "Pages"
	resp.Site.JsTopPage = template.HTML("<link href=\"../css/pages.css\" rel=\"stylesheet\">")

	err = contentTemplate["pages"].Execute(w, resp)
	if err != nil {
		log.Printf("Error executing pages: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
