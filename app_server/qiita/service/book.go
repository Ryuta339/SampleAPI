package service

import (
	"qiita/model"
)

type BookService struct{}

func (BookService) SetBook(book *model.Book) error {
	_, err := DbEngine.Insert(book)
	return err
}

func (BookService) GetBookList() []model.Book {
	tests := make([]model.Book, 0)
	if err := DbEngine.Distinct("id", "title", "content").Limit(10, 0).Find(&tests); err != nil {
		panic(err)
	}
	return tests
}

func (BookService) UpdateBook(newBook *model.Book) error {
	_, err := DbEngine.Id(newBook.Id).Update(newBook)
	return err
}

func (BookService) DeleteBook(id int) error {
	book := new(model.Book)
	_, err := DbEngine.Id(id).Delete(book)
	return err
}
