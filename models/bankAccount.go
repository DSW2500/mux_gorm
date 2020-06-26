package models

import (
	uuid "github.com/satori/go.uuid"
) // indirect

//Bank : Model for table
type Bank struct {
	Base
	Name   string
	Amount float64
	Type   string
	UserID uuid.UUID
}
