package bookshop_test

import (
	"github.com/adamluzsi/testcase"
	"github.com/mikejeuga/book_river_bookshop/bookshop"
	"github.com/mikejeuga/book_river_bookshop/models"
	"github.com/mikejeuga/book_river_bookshop/specifications"
	"testing"
)

func TestAddBookToShop(t *testing.T) {
	s := testcase.NewSpec(t)
	theBookShop := models.NewBookShop(make(models.Inventory))
	service := bookshop.NewService(theBookShop)
	spec := specifications.NewBookShopSpec(service)
	s.Describe("The BookShop Service", func(s *testcase.Spec) {
		s.Test("", func(t *testcase.T) {
			spec.GetBookInInventory(t)
		})
	})

}
