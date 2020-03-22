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

// Delete is delete func
func (d *DB) Delete(key string) {
	delete(d.B, key)
	return
}

// Get is read func
func (d *DB) Get(key string) (*models.Book, error) {
	if book, ok := d.B[key]; ok {
		return book, nil
	}
	err := errors.New("key not found")
	return nil, err

}

// Post is create func
func (d *DB) Post(key string, book *models.Book) models.Book {
	d.B[key] = book
	return *d.B[key]
}

// Update is update func
func (d *DB) Update(key string, book *models.Book) (v *models.Book) {
	d.B[key] = book
	v = d.B[key]
	return
}

// GetAll is for returning all the books
func (d *DB) GetAll() map[string]*models.Book {
	books := d.B
	return books
}
