package main

import (
	"context"
	"fmt"
	"github.com/eminetto/post-turso/book"
	"github.com/eminetto/post-turso/book/turso"
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
	repo.DisableAutoSync()
	defer repo.Close(ctx)
	s := book.NewService(repo)
	book, err := s.Create(ctx, "Neuromancer", "William Gibson", int64(book.Read))
	if err != nil {
		panic(err)
	}
	fmt.Println(book)
}
