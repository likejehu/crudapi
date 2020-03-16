package db

import (
	"reflect"
	"testing"
)

func TestNewDB(t *testing.T) {
	exp := NewDB()
	if reflect.TypeOf(exp).String() != "*db.DB" {
		t.Errorf("NewDB() = %v; want *db.DB", exp)
	}
}
