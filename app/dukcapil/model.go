package dukcapil

type (
	DataRequest struct {
		NIK          string `json:"nik" binding:"required,min=16,max=16"`
		NamaLengkap  string `json:"nama_lengkap" binding:"required"`
		JenisKelamin string `json:"jenis_kelamin"  binding:"required"`
		TempatLahir  string `json:"tempat_lahir" binding:"required"`
	}

	Content struct {
		NOKK           string `json:"NO_KK"`
		NamaLengkap    string `json:"NAMA_LGKP"`
		TempatLahir    string `json:"TMPT_LHR"`
		TanggalLahir   string `json:"TGL_LHR"`
		PropName       string `json:"PROP_NAME"`
		KabName        string `json:"KAB_NAME"`
		KecName        string `json:"KEC_NAME"`
		KelName        string `json:"KEL_NAME"`
		NoRT           string `json:"NO_RT"`
		NoRW           string `json:"NO_RW"`
		Alamat         string `json:"ALAMAT"`
		NamaIbu        string `json:"NAMA_LGKP_IBU"`
		StatusKawin    string `json:"STATUS_KAWIN"`
		JenisPekerjaan string `json:"JENIS_PKRJN"`
		JenisKelamin   string `json:"JENIS_KLMIN"`
		NoProp         string `json:"NO_PROP"`
		NoKab          string `json:"NO_KAB"`
		NoKec          string `json:"NO_KEC"`
		NoKel          string `json:"NO_KEL"`
		NIK            string `json:"NIK"`
	}

	Response struct {
		Content          []Content `json:"content"`
		LastPage         bool      `json:"lastPage"`
		NumberOfElements int       `json:"numberOfElements"`
		Sort             string    `json:"sort"`
		TotalElements    int       `json:"totalElements"`
		FirstPage        bool      `json:"firstPage"`
		Number           int       `json:"number"`
		Size             int       `json:"size"`
	}
)
