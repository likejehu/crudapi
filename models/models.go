package models

import (
	"time"
)

// Book  is mine book representation
type Book struct {
	ID          string    `json:"id" form:"id" query:"id"`
	Title       string    `json:"title" form:"title" query:"title"`
	Author      string    `json:"author" form:"author" query:"author"`
	Publisher   string    `json:"publisher" form:"publisher" query:"publisher"`
	PublishDate time.Time `json:"date" form:"date" query:"date"`
	Rating      uint8     `json:"rating" form:"rating" query:"rating"`
	Status      string    `json:"status" form:"status" query:"status"`
}
