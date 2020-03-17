package db

import (
	"testing"

	"github.com/likejehu/crudapi/models"
)

var testDB = NewDB()

var mb = map[string]*models.Book{
	"01": &models.Book{Title: "Priklyucheniya Igorya", Author: "Igor", Publisher: "Superizdatel", PublishDate: "2020-02-02", Rating: 0x1, Status: "CheckedIn"},
	"02": &models.Book{Title: "Novie priklyucheniya Igorya", Author: "Igor", Publisher: "Superizdatel", PublishDate: "2020-02-02", Rating: 0x2, Status: "CheckedIn"},
	"03": &models.Book{Title: "Supernovie priklyucheniya Igorya", Author: "Igor", Publisher: "Superizdatel", PublishDate: "2020-02-02", Rating: 0x3, Status: "CheckedIn"},
}

var tb = &models.Book{Title: "Super kniga", Author: "Igor", Publisher: "Superizdatel", PublishDate: "2020-02-02", Rating: 0x3, Status: "CheckedIn"}

func TestNewDB(t *testing.T) {
	var i interface{} = NewDB()
	d, ok := i.(*DB)
	if ok != true {
		t.Errorf("NewDB() = %T; want type *DB", d)
	}
}
func TestDelete(t *testing.T) {
	testDB.B = mb
	testDB.Delete("02")
	sv := testDB.B["02"]
	if sv != nil {
		t.Errorf("second value = %v; want  zero value", sv)
	}
}

func TestGet(t *testing.T) {
	testDB.B = mb
	sv := testDB.Get("02")
	if sv != testDB.B["02"] {
		t.Errorf("second value = %v; want  Novie priklyucheniya Igorya", sv)
	}
}

func TestPost(t *testing.T) {
	testDB.B = mb
	fv := testDB.Post("04", tb)
	if fv != *tb {
		t.Errorf("4th value = %v; want  Super kniga", fv)
	}

}
