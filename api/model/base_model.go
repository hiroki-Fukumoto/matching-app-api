package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string `gorm:"size:36;primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	now := time.Now()
	b.CreatedAt = &now
	return
}

func (b *Base) BeforeUpdate(db *gorm.DB) error {
	now := time.Now()
	b.UpdatedAt = &now
	return nil
}
