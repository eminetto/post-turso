package turso

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/eminetto/post-turso/book"
	"github.com/tursodatabase/go-libsql"
	"os"
	"path/filepath"
	"time"
)

type Repository struct {
	DB        *sql.DB
	dir       string
	connector *libsql.Connector
	autoSync  bool
}

var ErrNotFound = errors.New("not found")

func NewRepository(dbName, url, authToken string) (*Repository, error) {
	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		return nil, fmt.Errorf("creating temporary directory: %w", err)
	}

	dbPath := filepath.Join(dir, dbName)

	syncInterval := time.Minute

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, url,
		libsql.WithAuthToken(authToken),
		libsql.WithSyncInterval(syncInterval),
	)

	if err != nil {
		return nil, fmt.Errorf("creating connector: %w", err)
	}
	db := sql.OpenDB(connector)
	return &Repository{
		DB:        db,
		dir:       dir,
		connector: connector,
		autoSync:  true,
	}, nil
}

func (r *Repository) DisableAutoSync() {
	r.autoSync = false
}

func (r *Repository) Select(ctx context.Context, id string) (*book.Book, error) {
	rows, err := r.DB.Query("SELECT * FROM books where id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("selecting book: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var b book.Book

		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Category); err != nil {
			return nil, fmt.Errorf("scanning book: %w", err)
		}

		return &b, nil
	}
	return nil, ErrNotFound
}

func (r *Repository) SelectAll(ctx context.Context) ([]*book.Book, error) {
	rows, err := r.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, fmt.Errorf("selecting books: %w", err)
	}
	defer rows.Close()

	var books []*book.Book

	for rows.Next() {
		var b book.Book

		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Category); err != nil {
			return nil, fmt.Errorf("scanning book: %w", err)
		}

		books = append(books, &b)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("interacting with books: %w", err)
	}
	if len(books) == 0 {
		return nil, ErrNotFound
	}

	return books, nil
}

func (r *Repository) Insert(ctx context.Context, book *book.Book) (int64, error) {
	stmt, err := r.DB.Prepare(`
		insert into books (title, author, category) 
		values(?,?,?)`)
	if err != nil {
		return 0, fmt.Errorf("preparing statement: %w", err)
	}
	result, err := stmt.Exec(
		book.Title,
		book.Author,
		book.Category,
	)
	if err != nil {
		return 0, fmt.Errorf("executing statement: %w", err)
	}
	err = stmt.Close()
	if err != nil {
		return 0, fmt.Errorf("closing statement: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("getting last insert ID: %w", err)
	}
	if r.autoSync {
		_, err = r.connector.Sync()
		if err != nil {
			return 0, fmt.Errorf("syncing connector: %w", err)
		}
	}

	return id, nil
}

func (r *Repository) Update(ctx context.Context, book *book.Book) error {
	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	return nil
}

func (r *Repository) Close(ctx context.Context) error {
	err := os.RemoveAll(r.dir)
	if err != nil {
		return fmt.Errorf("removing temporary directory: %w", err)
	}
	err = r.connector.Close()
	if err != nil {
		return fmt.Errorf("closing connector: %w", err)
	}
	err = r.DB.Close()
	if err != nil {
		return fmt.Errorf("closing repository: %w", err)
	}
	return nil
}
