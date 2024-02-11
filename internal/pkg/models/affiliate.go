package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Affiliate struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey"`
	AdmitadID int       `gorm:"column:admitad_id;unique;index"`
	Share     int
}

func (a *Affiliate) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		a.ID = id
	}

	if a.AdmitadID == 0 {
		return gorm.ErrInvalidValue
	}

	if err := a.validateShare(); err != nil {
		return err
	}

	return nil
}

func (a *Affiliate) BeforeUpdate(tx *gorm.DB) error {
	if a.AdmitadID == 0 {
		return gorm.ErrInvalidValue
	}

	if err := a.validateShare(); err != nil {
		return err
	}

	return nil
}

func (a *Affiliate) validateShare() error {
	if a.Share <= 0 || a.Share > 100 {
		return gorm.ErrInvalidValue
	}

	return nil
}
