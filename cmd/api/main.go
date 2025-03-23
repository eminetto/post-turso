package main

import (
	"context"
	"fmt"
	"github.com/eminetto/post-turso/book"
	"github.com/eminetto/post-turso/book/turso"
	"github.com/eminetto/post-turso/config"
	"github.com/eminetto/post-turso/internal/http/chi"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

const TIMEOUT = 30 * time.Second

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT,
	)
	defer stop()
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

	errShutdown := make(chan error, 1)
	go shutdown(srv, ctx, errShutdown)
	fmt.Printf("Listening on port %s\n", cfg.Port)
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
		return
	}
	err = <-errShutdown
	if err != nil {
		fmt.Println(err)
		return
	}
}

func shutdown(server *http.Server, ctxShutdown context.Context, errShutdown chan error) {
	<-ctxShutdown.Done()

	ctxTimeout, stop := context.WithTimeout(context.Background(), TIMEOUT)
	defer stop()

	err := server.Shutdown(ctxTimeout)
	switch err {
	case nil:
		fmt.Printf("\nShutting down server...\n")
		errShutdown <- nil
	case context.DeadlineExceeded:
		errShutdown <- fmt.Errorf("Forcing closing the server")
	default:
		errShutdown <- fmt.Errorf("Forcing closing the server")
	}
}
