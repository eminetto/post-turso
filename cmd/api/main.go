package main

import (
	"context"
	"github.com/eminetto/post-turso/book"
	"github.com/eminetto/post-turso/book/turso"
	"github.com/eminetto/post-turso/internal/http/chi"
	"net/http"
	"os"
	"time"
)

func main() {

	//@todo colocar essas variáveis no pacote config usando .env e o viper
	dbName := "local.db"
	primaryUrl := "libsql://demo-post-eminetto.turso.io"
	authToken := "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3NDI1NjE1MjQsImlkIjoiMzUwMWEwZTAtZDMwMy00MTUzLWFjMjMtMTA1M2E0NjMwNmY2IiwicmlkIjoiYTViYmVjNDAtNzAyYi00MmI3LWE5NjEtN2QyM2Q5YTFlMTE3In0.5SWgVSIJrOMx8UFlBlkdUlYxAC6wSWXDGGRw0_f8TF7pWyULsJeB1mBCnhOVUq8tF7mhETgjGy55vuAzVEoyBA"
	ctx := context.Background()
	repo, err := turso.NewRepository(dbName, primaryUrl, authToken)
	if err != nil {
		panic(err)
	}
	repo.DisableAutoSync() //@todo: ver se é a melhor estratégia, pensando nos testes
	defer repo.Close(ctx)
	s := book.NewService(repo)
	r := chi.Handlers(ctx, s)
	http.Handle("/", r)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      http.DefaultServeMux,
	}
	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
