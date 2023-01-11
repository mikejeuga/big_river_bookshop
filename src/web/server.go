package web

import (
	"github.com/gorilla/mux"
	"github.com/mikejeuga/book_river_bookshop/specifications"
	"net/http"
)

type Server struct {
	BookShop specifications.Inventory
}

func NewServer(bookShop specifications.Inventory) *http.Server {
	r := mux.NewRouter()

	s := &Server{
		BookShop: bookShop,
	}

	r.HandleFunc("/", s.Home).Methods(http.MethodGet)

	return &http.Server{
		Addr:    ":8004",
		Handler: r,
	}
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {

}
