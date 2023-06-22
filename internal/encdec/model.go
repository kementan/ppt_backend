package encdec

type (
	StringRequest struct {
		Value string `json:"value" binding:"required"`
		Type  string `json:"type" binding:"required"`
		T     string `json:"t" binding:"required"`
	}

	StringResponse struct {
		Result string `json:"result"`
	}
)
