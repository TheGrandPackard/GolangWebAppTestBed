package database

import "errors"

// User Struct
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Users Struct
type Users []User

// GetUser from database
func GetUser(u string) (*User, error) {
	user := &User{}

	row := db.QueryRow("SELECT id, username, password FROM wiki.user WHERE username LIKE '" + u + "'")
	if row == nil {
		return &User{}, errors.New("No User: " + u)
	}

	if err := row.Scan(&user.ID, &user.Username, &user.Password); err == nil {
		return nil, err
	}

	return user, nil
}
