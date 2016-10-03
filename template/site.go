package template

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/thegrandpackard/wiki/database"
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
	User            *database.User
	Error           string
}

// SiteInit Helper
func SiteInit(r *http.Request) *Site {

	sess := session.GetSession(r)
	user := session.GetSessionUser(sess)

	return &Site{
		Title:           "Index",
		JsTopPage:       "",
		Content:         "",
		JsBotPage:       "",
		LastSearchQuery: "",
		Session:         sess,
		User:            user,
		Error:           "",
	}
}
