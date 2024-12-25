package db

import (
	"database/sql"
	"fmt"

	"gitlab.com/idoko/bucketeer/models"
)

func (db Database) GetAllBooks() (*models.BookList, error) {
	list := &models.BookList{}

	selectQuery :=
		`SELECT * 
		 FROM books
		 WHERE amount > 0
		 ORDER BY title ASC`

	rows, err := db.Conn.Query(selectQuery) // DESC - в порядке убывания
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.Isbn,
			&book.Title,
			&book.Author,
			&book.Publisher,
			&book.YearPublished,
			&book.Description,
			&book.Amount,
			&book.CreatedAt,
		)
		if err != nil {
			return list, err
		}
		list.Books = append(list.Books, book)
	}
	return list, nil
}

func (db Database) AddBook(book *models.Book) error {
	var isbn string
	var createdAt string
	query :=
		`INSERT INTO books (isbn, title, author, publisher, yearPublished, description, amount) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING isbn, created_at`
	err := db.Conn.QueryRow(
		query,
		book.Isbn,
		book.Title,
		book.Author,
		book.Publisher,
		book.YearPublished,
		book.Description,
		book.Amount,
	).Scan(&isbn, &createdAt)
	if err != nil {
		fmt.Println("err add book", err)
		return err
	}

	book.Isbn = isbn
	book.CreatedAt = createdAt
	return nil
}

func (db Database) GetBookByISBN(bookIsbn uint64) (models.Book, error) {
	book := models.Book{}

	query := `SELECT * FROM books WHERE isbn = $1;`
	row := db.Conn.QueryRow(query, bookIsbn)
	switch err := row.Scan(
		&book.Isbn,
		&book.Title,
		&book.Author,
		&book.Publisher,
		&book.YearPublished,
		&book.Description,
		&book.Amount,
		&book.CreatedAt,
	); err {
	case sql.ErrNoRows:
		return book, ErrNoMatch
	default:
		return book, err
	}
}

func (db Database) DeleteBook(bookIsbn uint64) error {
	query := `DELETE FROM books WHERE isbn = $1;`
	_, err := db.Conn.Exec(query, bookIsbn)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdateBook(bookIsbn uint64, bookData models.Book) (models.Book, error) {
	book := models.Book{}
	query :=
		`UPDATE books SET isbn=$1, title=$2, author=$3, publisher=$4, yearPublished=$5, description=$6,  amount=$7
		WHERE isbn=$8 
		RETURNING isbn, description, created_at;`
	err := db.Conn.QueryRow(
		query,
		bookData.Isbn,
		bookData.Title,
		bookData.Author,
		bookData.Publisher,
		bookData.YearPublished,
		bookData.Description,
		book.Amount,
		bookIsbn,
	).Scan(
		&book.Isbn,
		&book.Description,
		&book.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, ErrNoMatch
		}
		return book, err
	}

	return book, nil
}
