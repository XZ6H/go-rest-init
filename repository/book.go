package repository

import (
	"github.com/jinzhu/gorm"

	"myapp/model"
)

func ListBooks(db *gorm.DB) (model.Books, error) {
	books := make([]*model.Book, 0)
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}

	if len(books) == 0 {
		return nil, nil
	}

	return books, nil
}

func ReadBook(db *gorm.DB, id uint) (*model.Book, error) {
	book := &model.Book{}
	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func DeleteBook(db *gorm.DB, id uint) error {
	book := &model.Book{}
	if err := db.Where("id = ?", id).Delete(&book).Error; err != nil {
		return err
	}

	return nil
}