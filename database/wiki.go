package database

import (
	"errors"
	"html/template"
)

// WikiPage Struct
type WikiPage struct {
	ID      int    `db:"id"`
	Title   string `db:"title"`
	Body    string `db:"body"`
	Version int    `db:"version"`
}

// WikiPages Struct
type WikiPages []WikiPage

//GetWikiPage -- Get page by name
func GetWikiPage(title string) (*WikiPage, error) {
	page := &WikiPage{}
	if err := db.QueryRowx("SELECT * FROM wiki.page WHERE title LIKE ? AND version = (SELECT MAX(version) FROM wiki.page WHERE title LIKE ?)", title, title).StructScan(page); err != nil {
		return nil, err
	}
	return page, nil
}

// GetWikiPageVersion -- Get page by name and version
func GetWikiPageVersion(title string, version int) (*WikiPage, error) {
	page := &WikiPage{}
	if err := db.QueryRowx("SELECT * FROM wiki.page WHERE title LIKE ? AND version = ?", title, version).StructScan(page); err != nil {
		return nil, err
	}
	return page, nil
}

//GetWikiPages -- Get all  pages
func GetWikiPages() (WikiPages, error) {
	var pages WikiPages

	rows, err := db.Queryx("SELECT * FROM wiki.page WHERE wiki.page.version = (SELECT MAX(page_version.version) FROM wiki.page AS page_version WHERE page_version.id = wiki.page.id) GROUP BY id")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		page := WikiPage{}
		err = rows.StructScan(&page)
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)
	}

	return pages, nil
}

//Save -- Insert or Update page
func (page *WikiPage) Save() error {

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
		var newVersion int
		err := db.QueryRowx("SELECT MAX(version) FROM wiki.page WHERE id=?", page.ID).Scan(&newVersion)
		if err != nil {
			return err
		}
		page.Version = newVersion + 1

		stmt, err := db.NamedExec("INSERT INTO wiki.page set id=:id, title=:title, body=:body, version=:version", page)
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
