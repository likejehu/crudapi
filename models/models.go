package models

import "time"

// Book  is mine book representation
type Book struct {
	Title       string
	Author      string
	Publisher   string
	PublishDate time.Time
	Rating      uint8
	Status      string
}
