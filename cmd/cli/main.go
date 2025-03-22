package main

import (
	"context"
	"fmt"
	"github.com/eminetto/post-turso/book"
	"github.com/eminetto/post-turso/book/turso"
	"github.com/eminetto/post-turso/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx := context.Background()
	repo, err := turso.NewRepository(cfg.DBName, cfg.TursoDatabaseURL, cfg.TursoAuthToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer repo.Close(ctx)
	s := book.NewService(repo)
	book, err := s.Create(ctx, "Neuromancer", "William Gibson", book.Read)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(book)
}
