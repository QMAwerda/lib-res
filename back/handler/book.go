package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gitlab.com/idoko/bucketeer/db"
	"gitlab.com/idoko/bucketeer/models"
)

var bookIsbnKey = "bookIsbn"

func books(router chi.Router) {
	router.Get("/", getAllBooks)
	router.Post("/", createBook)

	router.Route("/{bookIsbn}", func(router chi.Router) {
		router.Use(BookContext)
		router.Get("/", getBook)
		router.Put("/", updateBook)
		router.Delete("/", deleteBook)
	})
}

func BookContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bookIsbn := chi.URLParam(r, "bookIsbn")
		if bookIsbn == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("book Isbn is required")))
			return
		}
		isbn, err := strconv.ParseUint(bookIsbn, 10, 64)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid book Isbn")))
		}
		ctx := context.WithValue(r.Context(), bookIsbnKey, isbn)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := dbInstance.GetAllBooks()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, books); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	if err := book.Bind(r); err != nil {
		fmt.Println("ошибка111", err)
		render.Render(w, r, ErrBadRequest)
		return
	}
	fmt.Println("пришел объект", book)
	if err := dbInstance.AddBook(book); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, book); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	bookIsbn := r.Context().Value(bookIsbnKey).(uint64)
	book, err := dbInstance.GetBookByISBN(bookIsbn)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}

	if book.Amount == 0 {
		msg := models.HasNoBook{
			MsgErr: "The book is out of stock",
		}
		if err = render.Render(w, r, &msg); err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
	} else {
		if err = render.Render(w, r, &book); err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	bookIsbn := r.Context().Value(bookIsbnKey).(uint64)
	err := dbInstance.DeleteBook(bookIsbn)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	bookIsbn := r.Context().Value(bookIsbnKey).(uint64)
	bookData := models.Book{}
	if err := render.Bind(r, &bookData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	book, err := dbInstance.UpdateBook(bookIsbn, bookData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &book); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
