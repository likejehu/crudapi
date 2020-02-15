package models

import "time"

// Book  is mine book representation
type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publisher   string
	PublishDate time.Time
	Rating      uint8
	Status      string
}

// Books  is mine books storage representation
var Books = map[int]*Book{}

// V is for vendetta
var V = 1

func main() {

}
