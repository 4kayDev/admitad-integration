package service

type PaginationInput struct {
	Limit  int
	Offset int
}

type GetOffersInput struct {
	PaginationInput
}
