package util

type (
	DataFilter struct {
		Search    string `json:"search"`
		SortBy    string `json:"sort_by"`
		SortOrder string `json:"sort_order"`
		Page      int    `json:"page"`
		PageSize  int    `json:"page_size"`
	}

	PaginationResponse struct {
		CurrentPage  int
		PageSize     int
		TotalPages   int
		TotalRecords int
	}

	TableIDs struct {
		IDs []string `json:"ids" binding:"required,dive,min=1"`
	}
)
