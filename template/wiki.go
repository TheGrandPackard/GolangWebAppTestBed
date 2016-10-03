package template

import (
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/thegrandpackard/wiki/database"
)

var wikiEditJsTop = template.HTML("<script src=\"//cdn.tinymce.com/4/tinymce.min.js\"></script>")
var wikiPagesJsTop = template.HTML("<link href=\"../css/pages.css\" rel=\"stylesheet\">")

func wikiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := vars["page"]

	if page == "" {
		wikiViewHandler(w, r, "home")
		return
	}

	var validPath = regexp.MustCompile("^/(edit|save|view)/")
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}

	switch m[1] {
	case "view":
		wikiViewHandler(w, r, page)
		return
	case "edit":
		wikiEditHandler(w, r, page)
		return
	case "save":
		wikiSaveHandler(w, r, page)
		return
	}
}

func wikiViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	r.ParseForm()
	version := r.FormValue("version")
	var p *database.WikiPage

	if version != "" {
		ver, _ := strconv.Atoi(version)
		p, _ = database.GetWikiPageVersion(title, ver)
	} else {
		p, _ = database.GetWikiPage(title)
	}

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

	if err := contentTemplate["view"].Execute(w, resp); err != nil {
		handleError(w, r, err.Error(), http.StatusInternalServerError)
	}
}

func wikiEditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := database.GetWikiPage(title)
	if err != nil {
		handleError(w, r, err.Error(), http.StatusInternalServerError)
		return
	} else if p == nil {
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
	resp.Site.JsTopPage = wikiEditJsTop

	if err = contentTemplate["edit"].Execute(w, resp); err != nil {
		handleError(w, r, err.Error(), http.StatusInternalServerError)
	}
}

func wikiSaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	id, _ := strconv.Atoi(r.FormValue("id"))
	p := &database.WikiPage{ID: id, Title: title, Body: body}

	err := p.Save()
	if err != nil {
		handleError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func wikiPagesHandler(w http.ResponseWriter, r *http.Request) {
	if checkLoggedIn(w, r) == false {
		handleError(w, r, "Page Not Found", http.StatusNotFound)
		return
	}

	p, err := database.GetWikiPages()
	if err != nil {
		handleError(w, r, err.Error(), http.StatusInternalServerError)
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
	resp.Site.JsTopPage = wikiPagesJsTop

	if err = contentTemplate["pages"].Execute(w, resp); err != nil {
		handleError(w, r, err.Error(), http.StatusInternalServerError)
	}
}
