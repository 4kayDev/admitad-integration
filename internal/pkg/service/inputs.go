package service

import "github.com/google/uuid"

type PaginationInput struct {
	Limit  int
	Offset int
}

type GetOffersInput struct {
	PaginationInput
}

type GetSavedOffersByHiddenInput struct {
	PaginationInput
	IsHidden bool
}

type GetSavedOffersInput struct {
	PaginationInput
}

type SaveOfferInput struct {
	AdmitadId int
}

type UpdateOfferInput struct {
	ID          uuid.UUID
	Name        string
	Description string
	ImageURL    string
	IsHidden    *bool
	SharedValue int
	UserValue   string
}

type DeleteOfferInput struct {
	ID uuid.UUID
}

type InitLinkInput struct {
	RequestId uuid.UUID
	ID        uuid.UUID
}
