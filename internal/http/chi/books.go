package chi

import (
	"encoding/json"
	"github.com/eminetto/post-turso/book"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type bookRequest struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

type bookResponse struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func getBooks(bookService book.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		all, err := bookService.List(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var result []bookResponse
		for _, b := range all {
			result = append(result, bookResponse{
				ID:       b.ID,
				Title:    b.Title,
				Author:   b.Author,
				Category: b.Category.String(),
			})
		}
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func getBook(bookService book.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		b, err := bookService.Get(r.Context(), int64(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		result := &bookResponse{
			ID:       b.ID,
			Title:    b.Title,
			Author:   b.Author,
			Category: b.Category.String(),
		}
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func postBooks(bookService book.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var br bookRequest
		err := json.NewDecoder(r.Body).Decode(&br)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//@todo validate data
		_, err = bookService.Create(r.Context(), br.Title, br.Author, book.NewCategory(br.Category))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})

}

func deleteBook(bookService book.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = bookService.Delete(r.Context(), int64(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func putBook(bookService book.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var br bookRequest
		err := json.NewDecoder(r.Body).Decode(&br)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//@todo validate data

		err = bookService.Update(r.Context(), int64(id), br.Title, br.Author, book.NewCategory(br.Category))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
