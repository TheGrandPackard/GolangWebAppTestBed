package database

import (
	"errors"
	"html/template"
)

// WikiPage Struct
type WikiPage struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
	Body  string `db:"body"`
}

// WikiPages Struct
type WikiPages []WikiPage

//GetWikiPage -- Get  page by name
func GetWikiPage(title string) (*WikiPage, error) {
	var page *WikiPage
	err := db.QueryRowx("SELECT id, title, body FROM wiki.page WHERE title LIKE ?", title).StructScan(page)
	return page, err
}

//GetWikiPages -- Get all  pages
func GetWikiPages() (WikiPages, error) {
	var pages WikiPages

	rows, err := db.Queryx("SELECT id, title, body FROM wiki.page")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		page := WikiPage{}
		err = rows.StructScan(page)
		pages = append(pages, page)
	}

	return pages, nil
}

//SavePage -- Insert or Update page
func (page *WikiPage) SavePage() error {
	if page.ID == 0 /* INSERT */ {
		stmt, err := db.NamedExec("INSERT INTO wiki.page set title=:title, body=:body", page)
		if err != nil {
			return err
		}
		id, err := stmt.LastInsertId()
		if err != nil {
			return err
		}
		page.ID = int(id)
	} else /* UPDATE */ {
		stmt, err := db.NamedExec("UPDATE wiki.page set title=:title, body=:body WHERE id=:id", page)
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

//GetBody -- For rendering pages
func (page *WikiPage) GetBody() template.HTML {
	return template.HTML(page.Body)
}
