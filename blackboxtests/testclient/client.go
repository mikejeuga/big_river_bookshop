package testclient

import (
	"github.com/google/uuid"
	"github.com/mikejeuga/book_river_bookshop/models"
	"net/http"
	"time"
)

type TestLibrarian struct {
	baseURL string
	engine  *http.Client
}

func (t *TestLibrarian) Add(book models.Book) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TestLibrarian) GetBookBy(id uuid.UUID) (models.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TestLibrarian) Update(book models.Book) (models.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TestLibrarian) Delete(book models.Book) error {
	//TODO implement me
	panic("implement me")
}

func NewTestLibrarian(baseURL string) *TestLibrarian {
	client := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   10 * time.Second,
	}
	return &TestLibrarian{baseURL: baseURL, engine: client}
}
