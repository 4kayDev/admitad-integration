package models

type ApiError struct {
	Description string `json:"error_description"`
	Code        int    `json:"error_code"`
}
