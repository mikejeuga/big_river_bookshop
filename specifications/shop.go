package specifications

import (
	"github.com/adamluzsi/testcase"
	"github.com/google/uuid"
	"github.com/mikejeuga/book_river_bookshop/models"
)

type Inventory interface {
	Add(book models.Book) (uuid.UUID, error)
	GetBookBy(id uuid.UUID) (models.Book, error)
	Update(book models.Book) (models.Book, error)
	Delete(book models.Book) error
}

type BookShopSpec struct {
	shopper Inventory
}

func NewBookShopSpec(shopper Inventory) *BookShopSpec {
	return &BookShopSpec{shopper: shopper}
}

func (b *BookShopSpec) AddBookForTheStock(t *testcase.T) {
	s := testcase.NewSpec(t)
	s.When("it adds a book to the bookshop inventory,", func(s *testcase.Spec) {
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

		act := func(t *testcase.T) (uuid.UUID, error) {
			return b.shopper.Add(book.Get(t))
		}

		s.Then("the book is acquired and added to the inventory.", func(t *testcase.T) {
			bookID, err := act(t)
			t.Must.NoError(err)

			gotBook, err := b.shopper.GetBookBy(bookID)
			t.Must.NoError(err)

			t.Must.Equal(book.Get(t).Author, gotBook.Author)
			t.Must.Equal(book.Get(t).Title, gotBook.Title)
		})

	})
}

func (b *BookShopSpec) UpdateBookInTheStock(t *testcase.T) {
	s := testcase.NewSpec(t)
	s.Describe("updating a book in the inventory", func(s *testcase.Spec) {
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

		act := func(t *testcase.T) (models.Book, error) {
			return b.shopper.Update(book.Get(t))
		}

		s.When("the book is in the shop and we need to update it,", func(s *testcase.Spec) {
			s.Before(func(t *testcase.T) {
				id, err := b.shopper.Add(book.Get(t))
				t.Must.NoError(err)

				theBook, err := b.shopper.GetBookBy(id)
				t.Must.NoError(err)

				theBook.Edition = 2020

				book.Set(t, theBook)
			})

			s.Then("it is updated in the inventory.", func(t *testcase.T) {
				updateBook, err := act(t)
				t.Must.NoError(err)

				t.Must.Equal(book.Get(t).Author, updateBook.Author)
				t.Must.Equal(book.Get(t).Title, updateBook.Title)
				t.Must.Equal(2020, updateBook.Edition)
			})

		})
	})
}

func (b *BookShopSpec) RemoveBookFromTheStock(t *testcase.T) {
	s := testcase.NewSpec(t)
	s.Describe("deleting a book from the inventory", func(s *testcase.Spec) {
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

		act := func(t *testcase.T) error {
			return b.shopper.Delete(book.Get(t))
		}

		s.When("the book is in the shop and is about to be sold out,", func(s *testcase.Spec) {
			s.Before(func(t *testcase.T) {
				id, err := b.shopper.Add(book.Get(t))
				t.Must.NoError(err)

				theBook, err := b.shopper.GetBookBy(id)
				t.Must.NoError(err)

				book.Set(t, theBook)
			})

			s.Then("it is deleted from the inventory.", func(t *testcase.T) {
				err := act(t)
				t.Must.NoError(err)

				_, err = b.shopper.GetBookBy(book.Get(t).ID)
				t.Must.NotNil(err)
			})

		})
	})
}
