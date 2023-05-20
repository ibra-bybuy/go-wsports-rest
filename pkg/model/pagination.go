package model

type Pagination struct {
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
}
