package session

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

//MaxLifetime max time (1800 seconds = 30 minutes)
var MaxLifetime = 1800
var store = sessions.NewCookieStore([]byte("thereoncewasapersonwhoneershallbenamed"))

// GetSession for access in Handlers
func GetSession(r *http.Request) *sessions.Session {

	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Printf("Error getting session for request")
	}

	return session
}
