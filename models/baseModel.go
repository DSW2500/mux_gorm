package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//Base :
type Base struct {
	ID        uuid.UUID `gorm:"primary_key"`
	CreatedOn *time.Time
	UpdatedOn *time.Time
	DeletedAt *time.Time
}
