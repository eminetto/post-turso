package book_test

import (
	"context"
	"fmt"
	"github.com/eminetto/post-turso/book"
	"github.com/eminetto/post-turso/book/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		b := &book.Book{
			Title:    "test title",
			Author:   "test author",
			Category: book.WantToRead,
		}
		repo := mocks.NewRepository(t)
		repo.On("Insert", ctx, b).Return(int64(1), nil)
		s := book.NewService(repo)
		saved, err := s.Create(ctx, "test title", "test author", book.WantToRead)
		assert.Nil(t, err)
		assert.Equal(t, int64(1), saved.ID)
		assert.Equal(t, "test title", saved.Title)
		assert.Equal(t, "test author", saved.Author)
		assert.Equal(t, book.WantToRead, saved.Category)
	})
	t.Run("fail", func(t *testing.T) {
		b := &book.Book{
			Title:    "test title",
			Author:   "test author",
			Category: book.WantToRead,
		}
		repo := mocks.NewRepository(t)
		repo.On("Insert", ctx, b).Return(int64(0), fmt.Errorf("some error"))
		s := book.NewService(repo)
		saved, err := s.Create(ctx, "test title", "test author", book.WantToRead)
		assert.NotNil(t, err)
		assert.Nil(t, saved)
	})
}
