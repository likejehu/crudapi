package db

import (
	"errors"

	"github.com/likejehu/crudapi/models"
)

// DB is mine representation of database
type DB struct {
	B map[string]*models.Book
}

//NewDB is for init new DB
func NewDB() *DB {
	var d DB
	d.B = make(map[string]*models.Book)
	return &d
}

// Library is mine library
var Library = NewDB()

// ErrorNotFound is 404 err for db, when key is not found
var ErrorNotFound = errors.New("key not found")

// Delete is delete func
func (d *DB) Delete(key string) error {
	if _, ok := d.B[key]; ok {
		delete(d.B, key)
		return nil
	}
	err := ErrorNotFound
	return err
}

// Get is read func
func (d *DB) Get(key string) (*models.Book, error) {
	if book, ok := d.B[key]; ok {
		return book, nil
	}
	err := ErrorNotFound
	return nil, err

}

// Post is create func
func (d *DB) Post(key string, book *models.Book) models.Book {
	d.B[key] = book
	return *d.B[key]

}

// Update is update func
func (d *DB) Update(key string, book *models.Book) (*models.Book, error) {
	if v, ok := d.B[key]; ok {
		d.B[key] = book
		v = d.B[key]
		return v, nil
	}
	err := ErrorNotFound
	return nil, err
}

// GetAll is for returning all the books
func (d *DB) GetAll() map[string]*models.Book {
	books := d.B
	return books
}
