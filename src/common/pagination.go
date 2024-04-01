package common

type Pagination struct {
	TotalItems   int64 `json:"total_items"`
	ItemsPerPage int   `json:"items_per_page"`
	CurrentPage  int64 `json:"current_page"`
	TotalPages   int64 `json:"total_pages"`
	HasNextPage  bool  `json:"has_next_page"`
	HasPrevPage  bool  `json:"has_prev_page"`
}

func GeneratePagination(totalItems int64, size int, from int64) Pagination {
	totalPages := (totalItems + int64(size) - 1) / int64(size)
	currentPage := from/int64(size) + 1
	hasNextPage := currentPage < totalPages
	hasPrevPage := currentPage > 1
	return Pagination{
		TotalItems:   totalItems,
		ItemsPerPage: size,
		CurrentPage:  currentPage,
		TotalPages:   totalPages,
		HasNextPage:  hasNextPage,
		HasPrevPage:  hasPrevPage,
	}
}
