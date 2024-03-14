package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Offer struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey"`
	AdmitadID   int       `gorm:"column:admitad_id;unique;index"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"description"`
	ImageURL    string    `gorm:"column:image_url"`
	ShareValue  int       `gorm:"column:shared_value"`
	UserValue   string    `gorm:"column:user_value"`
	Data        string    `gorm:"column:data"`
	Link        string    `gorm:"column:link"`
	IsHidden    *bool     `gorm:"column:is_hidden"`

	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (a *Offer) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		a.ID = id
	}

	if err := a.validateShareValue(); err != nil {
		return err
	}

	return nil
}

func (a *Offer) BeforeUpdate(tx *gorm.DB) error {
	if err := a.validateShareValue(); err != nil {
		return err
	}

	return nil
}

func (a *Offer) validateShareValue() error {
	if a.ShareValue < 0 || a.ShareValue > 100 {
		return gorm.ErrInvalidValue
	}

	return nil
}
