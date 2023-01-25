package bookshop

import (
	"github.com/google/uuid"
	"github.com/mikejeuga/book_river_bookshop/models"
	"github.com/mikejeuga/book_river_bookshop/specifications"
)

type Service struct {
	BookShop specifications.Inventory
}

func NewService(bookShop specifications.Inventory) *Service {
	return &Service{BookShop: bookShop}
}

func (s *Service) Add(b models.Book) (uuid.UUID, error) {
	return s.BookShop.Add(b)
}
func (s *Service) GetBookBy(id uuid.UUID) (models.Book, error) {
	return s.BookShop.GetBookBy(id)
}

func (s *Service) Update(book models.Book) (models.Book, error) {
	updatedBook, err := s.BookShop.Update(book)
	if err != nil {
		return models.Book{}, err
	}
	return updatedBook, nil
}
func (s *Service) Delete(book models.Book) error {
	return s.BookShop.Delete(book)
}
