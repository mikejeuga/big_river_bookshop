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
		var (
			book = testcase.Let(s, func(t *testcase.T) models.Book {
				return models.Book{
					Title: "The End of Everything",
					Author: models.Author{
						FirstName: "Darth",
						LastName:  "Vador",
					},
					Edition: 2011,
				}
			})
		)
		s.Test("adds a book", func(t *testcase.T) {
			spec.AddBookForTheStock(t, book)
		})

		s.Test("updates a book", func(t *testcase.T) {
			spec.UpdateBookInTheStock(t, book, 2022)
		})

		s.Test("deletes a book", func(t *testcase.T) {
			spec.RemoveBookFromTheStock(t, book)
		})
	})
}
