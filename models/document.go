package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	Name string
	URL  string
}

func (d *Document) TableName() string {
	return "documents"
}

func (d *Document) MapDocumentToResponse() response.DocumentResponse {

	documentResponse := response.DocumentResponse{
		ID:   d.ID,
		Name: d.Name,
		URL:  d.URL,
	}

	return documentResponse

}
