package chi

import (
	"encoding/json"
	"github.com/eminetto/post-turso/book"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func getBooks(bookService book.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := bookService.List(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(b); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func getBook(bookService book.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		//@todo: validate the data
		b, err := bookService.Get(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(b); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

type bookRequest struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func postBooks(bookService book.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var br bookRequest
		err := json.NewDecoder(r.Body).Decode(&br)
		if err != nil {
			//w.WriteHeader(http.StatusBadRequest)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//@todo validate data
		_, err = bookService.Create(r.Context(), br.Title, br.Author, book.NewCategory(br.Category))
		if err != nil {
			//w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})

}

func deleteBook(bookService book.UseCase) http.Handler {
	return nil
}

func putBook(bookService book.UseCase) http.Handler {
	return nil
}
