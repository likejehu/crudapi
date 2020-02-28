package models

type (
	// Book  is mine book representation
	Book struct {
		Title       string `json:"title" form:"title" query:"title" validate:"required"`
		Author      string `json:"author" form:"author" query:"author" validate:"required"`
		Publisher   string `json:"publisher" form:"publisher" query:"publisher" validate:"required"`
		PublishDate string `json:"date" form:"date" query:"date" validate:"required"`
		Rating      uint8  `json:"rating" form:"rating" query:"rating" validate:"min=1,max=3"`
		Status      string `json:"status" form:"status" query:"status" validate:"oneof=CheckedIn CheckedOut"`
	}
)
