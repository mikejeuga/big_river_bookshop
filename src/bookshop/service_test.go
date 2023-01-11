package bookshop_test

import (
	"github.com/adamluzsi/testcase"
	"github.com/mikejeuga/book_river_bookshop/models"
	"github.com/mikejeuga/book_river_bookshop/specifications"
	"github.com/mikejeuga/book_river_bookshop/src/bookshop"
	"testing"
)

func TestAddBookToShop(t *testing.T) {
	s := testcase.NewSpec(t)
	theBookShop := models.NewBookShop(make(models.Stock))
	service := bookshop.NewService(theBookShop)

	spec := specifications.NewBookShopSpec(service)

	s.Describe("The Service", func(s *testcase.Spec) {
		s.Test("adds a book", func(t *testcase.T) {
			spec.AddBookForTheStock(t)
		})

		s.Test("updates a book", func(t *testcase.T) {
			spec.UpdateBookInTheStock(t)
		})

		s.Test("deletes a book", func(t *testcase.T) {
			spec.RemoveBookFromTheStock(t)
		})
	})
}
