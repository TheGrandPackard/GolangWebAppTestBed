package main

import (
	"html/template"
)

// WikiPage Struct
type WikiPage struct {
	ID    int           `json:"id"`
	Title string        `json:"name"`
	Body  template.HTML `json:"body"`
}

// WikiPages Struct
type WikiPages []WikiPage

func getWikiPages() WikiPages {
	var pages WikiPages

	rows, err := db.Query("SELECT id, title, body FROM wiki.page")
	checkDBError(err)
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var body string
		err = rows.Scan(&id, &title, &body)

		page := WikiPage{
			ID:    id,
			Title: title,
			Body:  template.HTML(body),
		}

		pages = append(pages, page)
	}

	return pages
}

func getWikiPage(p string) (*WikiPage, error) {
	var id int
	var title string
	var body string

	err := db.QueryRow("SELECT id, title, body FROM wiki.page WHERE title LIKE '"+p+"'").Scan(&id, &title, &body)
	if err != nil {
		return nil, err
	}

	page := &WikiPage{
		ID:    id,
		Title: title,
		Body:  template.HTML(body),
	}

	return page, nil
}

func (p *WikiPage) savePage() error {

	if p.ID == 0 {
		// insert
		stmt, err := db.Prepare("INSERT INTO wiki.page set title=?, body=?")
		if err != nil {
			return err
		}

		res, err := stmt.Exec(p.Title, string(p.Body))
		if err != nil {
			return err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return err
		}

		p.ID = int(id)

	} else {
		// update
		stmt, err := db.Prepare("UPDATE wiki.page set title=?, body=? where id=?")
		if err != nil {
			return err
		}

		res, err := stmt.Exec(p.Title, string(p.Body), p.ID)
		if err != nil {
			return err
		}

		_, err = res.RowsAffected()
		if err != nil {
			return err
		}
	}

	return nil
}
