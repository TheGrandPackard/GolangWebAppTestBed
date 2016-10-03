package database

import (
	"errors"
	"strconv"
	"time"
)

// User Struct
type User struct {
	ID            int        `db:"id"`
	Username      string     `db:"username"`
	Password      string     `db:"password"`
	Enabled       bool       `db:"enabled"`
	DateAdded     time.Time  `db:"date_added"`
	DateLastLogin *time.Time `db:"date_last_login"`

	// Permissions
	ManageUsers bool `db:"manage_users"`
	ManagePages bool `db:"manage_pages"`
}

// Users Struct
type Users []User

// GetUser -- Get a user from database
func GetUser(username string) (*User, error) {
	user := &User{}
	err := db.QueryRowx("SELECT * FROM wiki.user WHERE username LIKE ?", username).StructScan(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

// GetUsers -- Get a user from database
func GetUsers(disabled bool) (Users, error) {
	var users Users
	var query = "SELECT * FROM wiki.user WHERE enabled = 1"
	if disabled {
		query = "SELECT * FROM wiki.user"
	}

	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.StructScan(&user)
		users = append(users, user)
	}

	return users, nil
}

// Save -- Insert or Update User
func (user *User) Save() error {
	if user.ID == 0 /* INSERT */ {
		stmt, err := db.NamedExec("INSERT INTO wiki.user set username=:username, password=:password", user)
		if err != nil {
			return err
		}
		id, err := stmt.LastInsertId()
		if err != nil {
			return err
		}
		user.ID = int(id)
	} else /* UPDATE */ {
		stmt, err := db.NamedExec("UPDATE wiki.user set username=:username, password=:password, enabled=:enabled, date_last_login=:date_last_login WHERE id=:id", user)
		if err != nil {
			return err
		}
		affected, err := stmt.RowsAffected()
		if err != nil {
			return err
		} else if affected != 1 {
			return errors.New("Rows Affected: " + strconv.Itoa(int(affected)))
		}
	}
	return nil
}

// CanManageUsers -- Check for ManageUsers
func (user *User) CanManageUsers() bool {
	return user.ManageUsers
}

// CanManagePages -- Check for ManagePages
func (user *User) CanManagePages() bool {
	return user.ManagePages
}
