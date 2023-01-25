package testclient

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/mikejeuga/book_river_bookshop/models"
	"io"
	"net/http"
	"net/url"
	"time"
)

type TestLibrarian struct {
	config Config
	engine *http.Client
}

type Config struct {
	BaseURL   string `envconfig:"BASE_URL"`
	BigSecret string `envconfig:"BIG_SECRET"`
}

func (bs *TestLibrarian) Add(book models.Book) (uuid.UUID, error) {
	addBookUrl, err := url.JoinPath(bs.config.BaseURL, "/books")
	if err != nil {
		return [16]byte{}, err
	}

	byteBook, err := json.Marshal(book)
	if err != nil {
		return [16]byte{}, err
	}

	req, err := http.NewRequest(http.MethodPost, addBookUrl, bytes.NewReader(byteBook))
	if err != nil {
		return [16]byte{}, err
	}

	req.Header.Set("X-API-KEY", bs.config.BigSecret)
	req.Header.Set("Content-Type", "application/json")

	res, err := bs.engine.Do(req)
	if err != nil {
		return [16]byte{}, err
	}

	var bookID uuid.UUID
	theID, err := io.ReadAll(res.Body)
	if err != nil {
		return [16]byte{}, err
	}

	err = json.Unmarshal(theID, &bookID)
	if err != nil {
		return [16]byte{}, err
	}

	return bookID, nil

}

func (bs *TestLibrarian) GetBookBy(id uuid.UUID) (models.Book, error) {
	getBookUrl, err := url.JoinPath(bs.config.BaseURL, "/books", id.String())
	if err != nil {
		return models.Book{}, err
	}

	req, err := http.NewRequest(http.MethodGet, getBookUrl, nil)
	if err != nil {
		return models.Book{}, err
	}

	req.Header.Set("X-API-KEY", bs.config.BigSecret)
	req.Header.Set("Content-Type", "application/json")

	res, err := bs.engine.Do(req)
	if err != nil {
		return models.Book{}, err
	}

	var requestedBook models.Book
	theBookData, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Book{}, err
	}

	err = json.Unmarshal(theBookData, &requestedBook)
	if err != nil {
		return models.Book{}, err
	}

	return requestedBook, nil

}

func (bs *TestLibrarian) Update(book models.Book) (models.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (bs *TestLibrarian) Delete(book models.Book) error {
	//TODO implement me
	panic("implement me")
}

func NewTestLibrarian(config Config) *TestLibrarian {
	client := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   10 * time.Second,
	}
	return &TestLibrarian{config: config, engine: client}
}
