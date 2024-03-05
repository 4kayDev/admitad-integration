package models

import (
	"time"

	"github.com/google/uuid"
)

type Click struct {
	ID        uuid.UUID  `gorm:"column:id;primaryKey"`
	OfferID   uuid.UUID  `gorm:"column:offer_id"`
	Link      string     `gorm:"column:link"`
	RequestId uuid.UUID  `gorm:"cloumn:request_id;unique"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
}
