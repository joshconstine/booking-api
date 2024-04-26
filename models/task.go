package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	AccountID         uint   `json:"accountId"`
	Message           string `json:"message"`
	MarkedAsCompleted bool   `json:"markedAsCompleted"`
	PriorityFlag      bool   `json:"priorityFlag"`
	DueDate           time.Time
}
