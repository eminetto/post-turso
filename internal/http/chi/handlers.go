package chi

import (
	"context"
	"github.com/eminetto/post-turso/book"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"net/http"
)

func Handlers(ctx context.Context, bookService book.UseCase) *chi.Mux {
	// Logger
	logger := httplog.NewLogger("post-turso", httplog.Options{
		JSON: true,
	})
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Method(http.MethodGet, "/v1/books", getBooks(bookService))
	r.Method(http.MethodGet, "/v1/books/{id}", getBook(bookService))
	r.Method(http.MethodPost, "/v1/books", postBooks(bookService))
	r.Method(http.MethodPut, "/v1/books/{id}", putBook(bookService))
	r.Method(http.MethodDelete, "/v1/books/{id}", deleteBook(bookService))

	return r
}
