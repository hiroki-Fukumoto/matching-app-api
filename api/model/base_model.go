package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}
