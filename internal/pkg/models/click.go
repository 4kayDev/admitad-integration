package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Click struct {
	ID        uuid.UUID  `gorm:"column:id;primaryKey"`
	OfferID   uuid.UUID  `gorm:"column:offer_id;"`
	RequestId uuid.UUID  `gorm:"cloumn:request_id;unique"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (c *Click) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	return nil
}
