package book

import "context"

type Reader interface {
	Select(ctx context.Context, id string) (*Book, error)
	SelectAll(ctx context.Context) ([]*Book, error)
}

type Writer interface {
	Insert(ctx context.Context, book *Book) (int64, error)
	Update(ctx context.Context, book *Book) error
	Delete(ctx context.Context, id string) error
}

type Repository interface {
	Reader
	Writer
	Close(ctx context.Context) error
}
