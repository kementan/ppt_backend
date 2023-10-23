package region

var table = "ppt_wilayah"

type (
	Wilayah struct {
		Kode string  `json:"kode"`
		Nama string  `json:"nama"`
		BMKG *string `json:"bmkg"`
	}
)
