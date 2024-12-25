package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Book struct {
	Isbn          string `json:"isbn"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Publisher     string `json:"publisher"`
	YearPublished string `json:"year_published"`
	Description   string `json:"description"`
	Amount        int    `json:"amount"`
	CreatedAt     string `json:"created_at"`
}

type HasNoBook struct {
	MsgErr string `json:"error"`
}

type BookList struct {
	Books []Book `json:"books"`
}

type BookTMP struct {
	Isbn          string `json:"isbn"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Publisher     string `json:"publisher"`
	YearPublished string `json:"year_published"`
	Description   string `json:"description"`
	Amount        string `json:"amount"`
	CreatedAt     string `json:"created_at"`
}

func (i *Book) Bind(r *http.Request) error {
	bookTmp := &BookTMP{}
	decoder := json.NewDecoder(r.Body)

	// Десериализуем JSON в структуру Book
	if err := decoder.Decode(bookTmp); err != nil {
		return fmt.Errorf("log1 %v", err)
	}
	defer r.Body.Close()

	amountInt, err := strconv.ParseInt(bookTmp.Amount, 10, 64)
	if err != nil {
		return fmt.Errorf("log2 %v", err)
	}

	i.Isbn = bookTmp.Isbn
	i.Amount = int(amountInt)
	i.Title = bookTmp.Title
	i.Author = bookTmp.Author
	i.Publisher = bookTmp.Publisher
	i.YearPublished = bookTmp.YearPublished
	i.Description = bookTmp.Description
	i.CreatedAt = bookTmp.CreatedAt

	return nil
}

func (*BookList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Book) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*HasNoBook) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
