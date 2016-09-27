package database

import "errors"

// User Struct
type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

// Users Struct
type Users []User

// GetUser -- Get a user from database
func GetUser(username string) (*User, error) {
	user := &User{}
	err := db.QueryRowx("SELECT id, username, password FROM wiki.user WHERE username LIKE ?", username).StructScan(user)
	if err != nil {
		return nil, errors.New("No User: " + username)
	}

	return user, nil
}

// SaveUser -- Insert or Update User
func (user *User) SaveUser() error {
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
		stmt, err := db.NamedExec("UPDATE wiki.user set username=:username, password=:password WHERE id=:id", user)
		if err != nil {
			return err
		}
		affected, err := stmt.RowsAffected()
		if err != nil {
			return err
		} else if affected != 1 {
			return errors.New("Rows Affected: " + string(affected))
		}
	}
	return nil
}
