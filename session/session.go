package session

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/thegrandpackard/wiki/database"
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

// GetSessionUser -- Get the user from a session, or return nil
func GetSessionUser(session *sessions.Session) *database.User {
	switch user := session.Values["user"].(type) {
	case database.User:
		return &user
	default:
		return nil
	}
}
