package chi

import (
	"context"
	"encoding/json"
	"github.com/eminetto/post-turso/book"
	"github.com/eminetto/post-turso/book/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooks(t *testing.T) {
	ctx := context.Background()
	s := mocks.NewUseCase(t)
	books := []*book.Book{
		{
			ID:       1,
			Title:    "Title 1",
			Author:   "Author 1",
			Category: book.WantToRead,
		},
		{
			ID:       2,
			Title:    "Title 2",
			Author:   "Author 2",
			Category: book.Read,
		},
	}
	s.On("List", mock.AnythingOfType("*context.valueCtx")).Return(books, nil)
	h := Handlers(ctx, s)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/v1/books", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var results []*bookResponse
	err = json.Unmarshal(w.Body.Bytes(), &results)
	assert.NoError(t, err)
	assert.Equal(t, len(books), len(results))
}
