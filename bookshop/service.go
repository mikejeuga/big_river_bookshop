package bookshop

import (
	"github.com/google/uuid"
	"github.com/mikejeuga/book_river_bookshop/models"
)

type Service struct {
	BookShop *models.BookShop
}

func NewService(bookShop *models.BookShop) *Service {
	return &Service{BookShop: bookShop}
}

func (s *Service) GetBook(id uuid.UUID) (models.Book, error) {
	return s.BookShop.Stock.GetBookBy(id)
}
func (s *Service) Add(b models.Book) (uuid.UUID, error) {
	return s.BookShop.Stock.Add(b)
}
