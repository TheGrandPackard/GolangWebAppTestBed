package template

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/thegrandpackard/wiki/session"
)

// Site Struct
type Site struct {
	Title           string
	UseHeader       bool
	JsTopPage       template.HTML
	JsBotPage       template.HTML
	Content         string
	LastSearchQuery string
	Session         *sessions.Session
	Error           string
}

// SiteInit Helper
func SiteInit(r *http.Request) *Site {

	session := session.GetSession(r)

	return &Site{
		Title:           "Index",
		JsTopPage:       "",
		Content:         "",
		JsBotPage:       "",
		LastSearchQuery: "",
		Session:         session,
		Error:           "",
	}
}
