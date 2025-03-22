package turso

import (
	"context"
	"database/sql"
	"github.com/eminetto/post-turso/book"
	_ "github.com/glebarez/go-sqlite"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestRepository(t *testing.T) {
	ctx := context.Background()
	dir, err := os.MkdirTemp("", "testlibsql-*")
	assert.Nil(t, err)
	dbName := "test.db"
	dbPath := filepath.Join(dir, dbName)
	db, err := sql.Open("sqlite", dbPath)
	defer db.Close()
	defer os.RemoveAll(dir)
	repo, err := NewTestRepository()
	assert.Nil(t, err)
	repo.SetDB(db)
	b := &book.Book{
		Title:    "Foundation",
		Author:   "Isaac Asimov",
		Category: book.Read,
	}
	err = repo.CreateTable(ctx)
	assert.Nil(t, err)
	id, err := repo.Insert(ctx, b)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), id)
	saved, err := repo.Select(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, b.Title, saved.Title)
	all, err := repo.SelectAll(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(all))
	assert.Equal(t, b.Title, all[0].Title)
	b.ID = id
	b.Title = "The Foundation"
	assert.Nil(t, repo.Update(ctx, b))
	saved, err = repo.Select(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, "The Foundation", saved.Title)
	assert.Nil(t, repo.Delete(ctx, id))
	all, err = repo.SelectAll(ctx)
	assert.Equal(t, ErrNotFound, err)

}
