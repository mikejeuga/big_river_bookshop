package bookshop

import (
	"github.com/google/uuid"
	"github.com/mikejeuga/book_river_bookshop/models"
)

type Service struct {
	BookShop *models.BookShop
}

func (s *Service) Delete(book models.Book) error {
	return s.BookShop.Stock.Delete(book)
}

func NewService(bookShop *models.BookShop) *Service {
	return &Service{BookShop: bookShop}
}

func (s *Service) GetBookBy(id uuid.UUID) (models.Book, error) {
	return s.BookShop.Stock.GetBookBy(id)
}

func (s *Service) Add(b models.Book) (uuid.UUID, error) {
	return s.BookShop.Stock.Add(b)
}
func (s *Service) Update(book models.Book) (models.Book, error) {
	updatedBook, err := s.BookShop.Stock.Update(book)
	if err != nil {
		return models.Book{}, err
	}
	return updatedBook, nil
}
