package main

import (
	"context"
	"fmt"
	"github.com/eminetto/post-turso/book"
	"github.com/eminetto/post-turso/book/turso"
	"github.com/eminetto/post-turso/config"
	"github.com/eminetto/post-turso/internal/http/chi"
	"net/http"
	"time"
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
	r := chi.Handlers(ctx, s)
	http.Handle("/", r)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + cfg.Port,
		Handler:      http.DefaultServeMux,
	}
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}

}
