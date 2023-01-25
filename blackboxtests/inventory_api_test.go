//+go:build acceptance

package blackboxtests

import (
	"fmt"
	"github.com/adamluzsi/testcase"
	"github.com/kelseyhightower/envconfig"
	"github.com/mikejeuga/book_river_bookshop/blackboxtests/testclient"
	"github.com/mikejeuga/book_river_bookshop/models"
	"github.com/mikejeuga/book_river_bookshop/specifications"
	"testing"
)

func TestBookShopInventory(t *testing.T) {
	s := testcase.NewSpec(t)

	var testConfig testclient.Config
	err := envconfig.Process("", &testConfig)
	if err != nil {
		t.Fatal("Could not load environment variables!")
	}

	fmt.Println(testConfig)

	s.Describe("A Book Seller", func(s *testcase.Spec) {
		testBookSeller := testclient.NewTestLibrarian(testConfig)
		spec := specifications.NewBookShopSpec(testBookSeller)

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

		s.Test("can add a book to the shop stock", func(t *testcase.T) {
			spec.AddBookForTheStock(t, book)
		})
	})
}
