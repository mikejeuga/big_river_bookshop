package blackboxtests

import (
	"github.com/adamluzsi/testcase"
	"github.com/mikejeuga/book_river_bookshop/blackboxtests/testclient"
	"github.com/mikejeuga/book_river_bookshop/specifications"
	"testing"
)

func TestBookShopInventory(t *testing.T) {
	s := testcase.NewSpec(t)

	testBookSeller := testclient.NewTestLibrarian("http://localhost:8004")
	spec := specifications.NewBookShopSpec(testBookSeller)

	s.Describe("A Book Seller", func(s *testcase.Spec) {
		s.Test("can add a book to the shop stock", func(t *testcase.T) {
			t.Skip("Until server implemented")
			spec.AddBookForTheStock(t)
		})
	})
}
