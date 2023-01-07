package specifications

import (
	"github.com/adamluzsi/testcase"
	"github.com/google/uuid"
	"github.com/mikejeuga/book_river_bookshop/models"
)

type Shop interface {
	Add(book models.Book) (uuid.UUID, error)
	GetBook(id uuid.UUID) (models.Book, error)
}

type BookShopSpec struct {
	shopper Shop
}

func NewBookShopSpec(shopper Shop) *BookShopSpec {
	return &BookShopSpec{shopper: shopper}
}

func (b *BookShopSpec) GetBookInInventory(t *testcase.T) {
	s := testcase.NewSpec(t)
	s.Describe("The BookShop Behaviour", func(s *testcase.Spec) {

		s.When("The BookShop adds a Book to its stock", func(s *testcase.Spec) {
			var (
				book = testcase.Let(s, func(t *testcase.T) models.Book {
					return models.Book{
						Title: "The End of Everything",
						Author: models.Author{
							FirstName: "Darth",
							LastName:  "Vador",
						},
					}
				})
			)

			act := func(t *testcase.T) (uuid.UUID, error) {
				return b.shopper.Add(book.Get(t))
			}

			s.Then("The book is acquired and added to the stock", func(t *testcase.T) {
				bookID, err := act(t)
				t.Must.NoError(err)

				gotBook, err := b.shopper.GetBook(bookID)
				t.Must.NoError(err)

				t.Must.Equal(book.Get(t).Author, gotBook.Author)
				t.Must.Equal(book.Get(t).Title, gotBook.Title)
			})

		})

	})

}
