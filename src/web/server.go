package web

import (
	"encoding/json"
	"fmt"
	"github.com/adamluzsi/testcase/pp"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mikejeuga/book_river_bookshop/models"
	"github.com/mikejeuga/book_river_bookshop/specifications"
	"github.com/mikejeuga/book_river_bookshop/src/web/auth"
	"io"
	"net/http"
)

type Server struct {
	Config   auth.Config
	BookShop specifications.Inventory
}

func NewServer(config auth.Config, bookShop specifications.Inventory) *http.Server {
	r := mux.NewRouter()

	s := &Server{
		Config:   config,
		BookShop: bookShop,
	}

	r.Use(auth.FOMW(s.Config.BigSecret))

	r.HandleFunc("/", s.Home).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", s.GetBook).Methods(http.MethodGet)
	r.HandleFunc("/books", s.AddBook).Methods(http.MethodPost)

	return &http.Server{
		Addr:    ":8004",
		Handler: r,
	}
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I AM UP AND HEALTHY!")
}

func (s *Server) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	theBookID, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	book, err := s.BookShop.GetBookBy(theBookID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) AddBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book

	bookdata, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(bookdata, &newBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, err := s.BookShop.Add(newBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pp.PP(id)

	err = json.NewEncoder(w).Encode(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
