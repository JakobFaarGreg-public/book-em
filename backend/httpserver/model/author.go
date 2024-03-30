package model

import (
	"database/sql"
	"log"
)

type Author struct {
	id        int
	firstName string
	lastName  string
}

type AuthorDAO struct {
	db *sql.DB
}

func (dao *AuthorDAO) GetAuthorByID(userID int) (*Author, error) {
	var author Author
	err := dao.db.QueryRow(
		"SELECT id, first_name, last_name FROM author WHERE id = $1", userID).Scan(
		author.id, author.firstName, author.lastName)
	if err != nil {
		log.Fatal(err)
	}

	return &author, err
}
