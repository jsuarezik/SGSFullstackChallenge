package responses

type ProductResponse struct {
	Status     int                `json:"status"`
	Message    string             `json:"message"`
	Data       interface{}        `json:"data"`
	Pagination PaginationMetadata `json:"pagination"`
}

type PaginationMetadata struct {
	CurrentPage int `json:"current"`
	PageSize    int `json:"page_size"`
	TotalPages  int `json:"total_pages"`
	TotalCount  int `json:"total_count"`
}
