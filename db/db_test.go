package db

import (
	"testing"
)

func TestNewDB(t *testing.T) {
	var i interface{} = NewDB()
	d, ok := i.(*DB)
	if ok != true {
		t.Errorf("NewDB() = %v; want type *DB", d)
	}
}
