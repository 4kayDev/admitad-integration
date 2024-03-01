package service

import "github.com/google/uuid"

type PaginationInput struct {
	Limit  int
	Offset int
}

type GetOffersInput struct {
	PaginationInput
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
	SharedValue int
}

type DeleteOfferInput struct {
	ID uuid.UUID
}

type InitLinkInput struct {
	RequestId uuid.UUID
	ID        uuid.UUID
}
