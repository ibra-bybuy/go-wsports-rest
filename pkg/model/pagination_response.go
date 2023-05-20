package model

type PaginationResponse struct {
	Pagination Pagination  `json:"pagination,omitempty"`
	Items      interface{} `json:"items"`
}
