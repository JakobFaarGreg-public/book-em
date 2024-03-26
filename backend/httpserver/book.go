package main

import (
	"database/sql"
	"log"
)

type Book struct {
	id     int
	author Author
	title  string
}

type BookDAO struct {
	db *sql.DB
}

func (dao *BookDAO) GetBookByID(bookID int) (*Book, error) {
	var book Book
	err := dao.db.QueryRow(
		"SELECT book.id AS book_id, book.title AS book_title, "+
			"author.id as author_id, author.first_name as author_first_name, author.last_name as author_last_name "+
			"FROM book JOIN author on book.author_id = author.id WHERE book.id = $1",
		bookID).Scan(book.id, book.title, book.author.id, book.author.firstName, book.author.lastName)
	if err != nil {
		log.Fatal(err)
	}

	return &book, err
}
