package models

import (
	"fmt"
	"github.com/google/uuid"
)

type BookShop struct {
	Stock Stock
}

func NewBookShop(stock Stock) *BookShop {
	return &BookShop{Stock: stock}
}

type Stock map[uuid.UUID]Book

func (b BookShop) Add(book Book) (uuid.UUID, error) {
	newBookID, err := uuid.NewUUID()
	if err != nil {
		return [16]byte{}, err
	}
	book.ID = newBookID
	b.Stock[book.ID] = book

	return book.ID, nil
}

func (b BookShop) GetBookBy(bookID uuid.UUID) (Book, error) {
	book, ok := b.Stock[bookID]
	if !ok {
		return Book{}, fmt.Errorf("this book does not exist in the inventory. bookID: %v", bookID)
	}
	return book, nil
}

func (b BookShop) Update(book Book) (Book, error) {
	b.Stock[book.ID] = book
	return b.Stock[book.ID], nil
}

func (b BookShop) Delete(book Book) error {
	book, ok := b.Stock[book.ID]
	if ok {
		delete(b.Stock, book.ID)
		return nil
	}
	return fmt.Errorf("there is no unit of this book in the stock, %v", book)
}
