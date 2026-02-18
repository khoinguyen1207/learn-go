package utils

type Pagination struct {
	Page         int32 `form:"page" json:"page"`
	Limit        int32 `form:"limit" json:"limit"`
	TotalRecords int32 `json:"total_records"`
	TotalPages   int32 `json:"total_pages"`
	HasNext      bool  `json:"has_next"`
	HasPrev      bool  `json:"has_prev"`
}

func NewPagination(page, limit, totalRecords int32) *Pagination {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	totalPages := (totalRecords + limit - 1) / limit

	return &Pagination{
		Page:         page,
		Limit:        limit,
		TotalRecords: totalRecords,
		TotalPages:   totalPages,
		HasNext:      page < totalPages,
		HasPrev:      page > 1,
	}
}
