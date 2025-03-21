package book

import (
	"context"
	"fmt"
)

type UseCase interface {
	Create(ctx context.Context, title, author string, category Category) (*Book, error)
	List(ctx context.Context) ([]*Book, error)
	Get(ctx context.Context, id string) (*Book, error)
	Update(ctx context.Context, id int64, title, author string, category int64) error
	Delete(ctx context.Context, id string) error
}

type Service struct {
	Repo Repository
}

// Interface guard
var _ UseCase = (*Service)(nil)

func NewService(repo Repository) *Service {
	return &Service{
		Repo: repo,
	}
}
func (s *Service) Create(ctx context.Context, title, author string, category Category) (*Book, error) {
	b := Book{
		Title:    title,
		Author:   author,
		Category: category,
	}
	id, err := s.Repo.Insert(ctx, &b)
	if err != nil {
		return nil, fmt.Errorf("inserting book: %w", err)
	}
	b.ID = id
	return &b, nil
}

func (s *Service) List(ctx context.Context) ([]*Book, error) {
	all, err := s.Repo.SelectAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("selecting books: %w", err)
	}
	return all, nil
}
func (s *Service) Get(ctx context.Context, id string) (*Book, error) {
	b, err := s.Repo.Select(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("selecting book: %w", err)
	}
	return b, nil
}

func (s *Service) Update(ctx context.Context, id int64, title, author string, category int64) error {
	b := Book{
		ID:       id,
		Title:    title,
		Author:   author,
		Category: Category(category),
	}
	err := s.Repo.Update(ctx, &b)
	if err != nil {
		return fmt.Errorf("updating book: %w", err)
	}
	return nil
}
func (s *Service) Delete(ctx context.Context, id string) error {
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("deleting book: %w", err)
	}
	return nil
}
