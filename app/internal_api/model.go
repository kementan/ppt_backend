package internal_api

import "time"

var table = "ppt_api_data_storages"

type (
	PptApiDataStorage struct {
		ID         interface{} `json:"id"`
		Identifier interface{} `json:"identifier"`
		F1         interface{} `json:"f1"`
		F2         interface{} `json:"f2"`
		F3         interface{} `json:"f3"`
		F4         interface{} `json:"f4"`
		F5         interface{} `json:"f5"`
		F6         interface{} `json:"f6"`
		F7         interface{} `json:"f7"`
		F8         interface{} `json:"f8"`
		F9         interface{} `json:"f9"`
		F10        interface{} `json:"f10"`
		F11        interface{} `json:"f11"`
		F12        interface{} `json:"f12"`
		F13        interface{} `json:"f13"`
		F14        interface{} `json:"f14"`
		F15        interface{} `json:"f15"`
		F16        interface{} `json:"f16"`
		F17        interface{} `json:"f17"`
		F18        interface{} `json:"f18"`
		F19        interface{} `json:"f19"`
		F20        interface{} `json:"f20"`
		F21        interface{} `json:"f21"`
		F22        interface{} `json:"f22"`
		F23        interface{} `json:"f23"`
		F24        interface{} `json:"f24"`
		F25        interface{} `json:"f25"`
		LongText   interface{} `json:"longtext"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
	}

	Pelatihan struct {
		StsIkut      interface{} `json:"sts_ikut"`
		IdPel        interface{} `json:"id_pel"`
		KodePel      interface{} `json:"kode_pel"`
		JudulPel     interface{} `json:"judul_pel"`
		LinkMateri   interface{} `json:"link_materi"`
		TanggalMulai interface{} `json:"tanggal_mulai"`
		JamMulai     interface{} `json:"jam_mulai"`
		TanggalAkhir interface{} `json:"tanggal_akhir"`
		JamAkhir     interface{} `json:"jam_akhir"`
		LinkFlyer    interface{} `json:"link_flyer"`
	}

	// untuk nasional, bpsb, dan lssm karena struktur mirip
	PerbenihanData1 struct {
		NO                 interface{} `json:"NO"`
		JENIS              interface{} `json:"JENIS"`
		PROVINSI           interface{} `json:"PROVINSI"`
		JENIS_BENIH        interface{} `json:"JENIS_BENIH"`
		KELAS_BENIH        interface{} `json:"KELAS_BENIH"`
		VARIETAS           interface{} `json:"VARIETAS"`
		REALISASI_LUAS     interface{} `json:"REALISASI_LUAS"`
		REALISASI_PRODUKSI interface{} `json:"REALISASI_PRODUKSI"`
		VOLUME             interface{} `json:"VOLUME"`
		DICATAT            interface{} `json:"DICATAT"`
		DIPERBARUI         interface{} `json:"DIPERBARUI"`
	}

	// penyaluran
	PerbenihanData2 struct {
		NO             interface{} `json:"NO"`
		TAHUN          interface{} `json:"TAHUN"`
		BULAN          interface{} `json:"BULAN"`
		PROVINSI       interface{} `json:"PROVINSI"`
		KABUPATENKOTA  interface{} `json:"KABUPATENKOTA"`
		KECAMATAN      interface{} `json:"KECAMATAN"`
		PRODUSEN_BENIH interface{} `json:"PRODUSEN_BENIH"`
		KELAS_BENIH    interface{} `json:"KELAS_BENIH"`
		KOMODITI       interface{} `json:"KOMODITI"`
		VARIETAS       interface{} `json:"VARIETAS"`
		STOK_LALU      interface{} `json:"STOK_LALU"`
		PRODUKSI_BENIH interface{} `json:"PRODUKSI_BENIH"`
		PENGADAAN      interface{} `json:"PENGADAAN"`
		JUMLAH_STOK    interface{} `json:"JUMLAH_STOK"`
		PENYALURAN     interface{} `json:"PENYALURAN"`
		APBN           interface{} `json:"APBN"`
		APBD           interface{} `json:"APBD"`
		FREE_MARKET    interface{} `json:"FREE_MARKET"`
		JUMLAH_SALUR   interface{} `json:"JUMLAH_SALUR"`
		TOTAL          interface{} `json:"TOTAL"`
		SISA_STOK      interface{} `json:"SISA_STOK"`
		DICATAT        interface{} `json:"DICATAT"`
		DIPERBARUI     interface{} `json:"DIPERBARUI"`
	}

	// penyebaran
	PerbenihanData3 struct {
		NO                   interface{} `json:"NO"`
		TAHUN                interface{} `json:"TAHUN"`
		BULAN                interface{} `json:"BULAN"`
		PROVINSI             interface{} `json:"PROVINSI"`
		KABUPATENKOTA        interface{} `json:"KABUPATENKOTA"`
		KECAMATAN            interface{} `json:"KECAMATAN"`
		KELURAHAN            interface{} `json:"KELURAHAN"`
		PETA                 interface{} `json:"PETA"`
		REALISASI_TANAM_LUAS interface{} `json:"REALISASI_TANAM_LUAS"`
		BENIH                interface{} `json:"BENIH"`
		JENIS_BENIH          interface{} `json:"JENIS_BENIH"`
		VARIETAS             interface{} `json:"VARIETAS"`
		TOTAL_LUAS           interface{} `json:"TOTAL_LUAS"`
		DICATAT              interface{} `json:"DICATAT"`
		DIPERBARUI           interface{} `json:"DIPERBARUI"`
	}

	// produsen
	PerbenihanData4 struct {
		NO               interface{} `json:"NO"`
		KODE_PROVINSI    interface{} `json:"KODE_PROVINSI"`
		PROVINSI         interface{} `json:"PROVINSI"`
		KABUPATENKOTA    interface{} `json:"KABUPATENKOTA"`
		KECAMATAN        interface{} `json:"KECAMATAN"`
		KELURAHAN        interface{} `json:"KELURAHAN"`
		USERNAME         interface{} `json:"USERNAME"`
		IDSIMLUH         interface{} `json:"IDSIMLUH"`
		NOMOR_REGISTRASI interface{} `json:"NOMOR_REGISTRASI"`
		TIPE_PRODUSEN    interface{} `json:"TIPE_PRODUSEN"`
		NAMA             interface{} `json:"NAMA"`
		NAMA_PIMPINAN    interface{} `json:"NAMA_PIMPINAN"`
		ALAMAT_PIMPINAN  interface{} `json:"ALAMAT_PIMPINAN"`
		ALAMAT_PRODUSEN  interface{} `json:"ALAMAT_PRODUSEN"`
		TELEPON          interface{} `json:"TELEPON"`
		EMAIL            interface{} `json:"EMAIL"`
		BENIH            interface{} `json:"BENIH"`
		TOTAL_LUAS_LAHAN interface{} `json:"TOTAL_LUAS_LAHAN"`
		LAT              interface{} `json:"LAT"`
		LNG              interface{} `json:"LNG"`
		DICATAT          interface{} `json:"DICATAT"`
		DIPERBARUI       interface{} `json:"DIPERBARUI"`
	}

	// Untuk Nas, BPSB, LSSM
	Perbenihan1 struct {
		Note []string          `json:"note"`
		Data []PerbenihanData1 `json:"data"`
	}

	//Untuk Penyaluran
	Perbenihan2 struct {
		Note []string          `json:"note"`
		Data []PerbenihanData2 `json:"data"`
	}

	// Untuk Penyebaran
	Perbenihan3 struct {
		Note []string          `json:"note"`
		Data []PerbenihanData3 `json:"data"`
	}

	// Untuk Produsen
	Perbenihan4 struct {
		Note []string          `json:"note"`
		Data []PerbenihanData4 `json:"data"`
	}

	// SIPDPS Tanam
	SIPDPSTanam struct {
		NIPReporter        interface{} `json:"nip_reporter"`
		NamaReporter       interface{} `json:"nm_reporter"`
		TanggalLaporan     interface{} `json:"tgl_lapor"`
		TanggalKunjungan   interface{} `json:"tgl_kunjungan"`
		JenisKelompok      interface{} `json:"jenis_kelompok"`
		NamaProvinsi       interface{} `json:"nm_prov"`
		NamaKabupaten      interface{} `json:"nm_kab"`
		NamaKecamatan      interface{} `json:"nm_kec"`
		NamaDesa           interface{} `json:"nm_desa"`
		KategoriLahan      interface{} `json:"kategori_lahan"`
		JenisLahan         interface{} `json:"jenis_lahan"`
		JenisTanamanPangan interface{} `json:"jenis_tanaman_pangan"`
		NamaVarietas       interface{} `json:"nm_varietas"`
		JenisBantuan       interface{} `json:"jenis_bantuan"`
		SumberBantuan      interface{} `json:"sumber_bantuan"`
		TahunBantuan       interface{} `json:"tahun_bantuan"`
		LuasArea           interface{} `json:"luas_area"`
		HST                interface{} `json:"hst"`
		Latitude           interface{} `json:"lat"`
		Longitude          interface{} `json:"lng"`
		Photos             interface{} `json:"photos"`
		Status             interface{} `json:"status"`
	}

	// SIPDPS Produktivitas
	SIPDPSProduktivitas struct {
		NIPReporter        interface{} `json:"nip_reporter"`
		NamaReporter       interface{} `json:"nm_reporter"`
		TanggalLaporan     interface{} `json:"tgl_lapor"`
		TanggalKunjungan   interface{} `json:"tgl_kunjungan"`
		NamaProvinsi       interface{} `json:"nm_prov"`
		NamaKabupaten      interface{} `json:"nm_kab"`
		NamaKecamatan      interface{} `json:"nm_kec"`
		NamaDesa           interface{} `json:"nm_desa"`
		KategoriLahan      interface{} `json:"kategori_lahan"`
		JenisLahan         interface{} `json:"jenis_lahan"`
		JenisTanamanPangan interface{} `json:"jenis_tanaman_pangan"`
		TeknikPengukuran   interface{} `json:"teknik_pengukuran"`
		Jumlah             interface{} `json:"jumlah"`
		Latitude           interface{} `json:"lat"`
		Longitude          interface{} `json:"lng"`
		Photos             interface{} `json:"photos"`
		NamaVerifikator    interface{} `json:"nm_verifikator"`
		Status             interface{} `json:"status"`
	}

	// SIPDPS Puso
	SIPDPSPuso struct {
		NIPReporter        interface{} `json:"nip_reporter"`
		NamaReporter       interface{} `json:"nm_reporter"`
		TanggalLaporan     interface{} `json:"tgl_lapor"`
		TanggalKejadian    interface{} `json:"tgl_kejadian"`
		NamaProvinsi       interface{} `json:"nm_prov"`
		NamaKabupaten      interface{} `json:"nm_kab"`
		NamaKecamatan      interface{} `json:"nm_kec"`
		NamaDesa           interface{} `json:"nm_desa"`
		JenisTanamanPangan interface{} `json:"jenis_tanaman_pangan"`
		PenyebabPuso       interface{} `json:"penyebab_puso"`
		Latitude           interface{} `json:"lat"`
		Longitude          interface{} `json:"lng"`
		Photos             interface{} `json:"photos"`
		NamaVerifikator    interface{} `json:"nm_verifikator"`
		Status             interface{} `json:"status"`
	}

	// SIPDPS Panen
	SIPDPSPanen struct {
		NIPReporter        interface{} `json:"nip_reporter"`
		NamaReporter       interface{} `json:"nm_reporter"`
		TanggalLaporan     interface{} `json:"tgl_lapor"`
		TanggalKunjungan   interface{} `json:"tgl_kunjungan"`
		NamaProvinsi       interface{} `json:"nm_prov"`
		NamaKabupaten      interface{} `json:"nm_kab"`
		NamaKecamatan      interface{} `json:"nm_kec"`
		NamaDesa           interface{} `json:"nm_desa"`
		JenisTanamanPangan interface{} `json:"jenis_tanaman_pangan"`
		NamaVarietas       interface{} `json:"nm_varietas"`
		KategoriPengelola  interface{} `json:"kategori_pengelola"`
		NamaPengelola      interface{} `json:"nama_pengelola"`
		Luas               interface{} `json:"luas"`
		Perkiraan          interface{} `json:"perkiraan"`
		Latitude           interface{} `json:"lat"`
		Longitude          interface{} `json:"lng"`
		Photos             interface{} `json:"photos"`
		NamaVerifikator    interface{} `json:"nm_verifikator"`
		Status             interface{} `json:"status"`
	}

	SIPDPS1 struct {
		CurrentPage  int           `json:"current_page"`
		Data         []SIPDPSTanam `json:"data"`
		FirstPageURL string        `json:"first_page_url"`
		From         int           `json:"from"`
		LastPage     int           `json:"last_page"`
		LastPageURL  string        `json:"last_page_url"`
		NextPageURL  string        `json:"next_page_url"`
		Path         string        `json:"path"`
		PerPage      int           `json:"per_page"`
		PrevPageURL  string        `json:"prev_page_url"`
		To           int           `json:"to"`
		Total        int           `json:"total"`
	}

	SIPDPS2 struct {
		CurrentPage  int                   `json:"current_page"`
		Data         []SIPDPSProduktivitas `json:"data"`
		FirstPageURL string                `json:"first_page_url"`
		From         int                   `json:"from"`
		LastPage     int                   `json:"last_page"`
		LastPageURL  string                `json:"last_page_url"`
		NextPageURL  string                `json:"next_page_url"`
		Path         string                `json:"path"`
		PerPage      int                   `json:"per_page"`
		PrevPageURL  string                `json:"prev_page_url"`
		To           int                   `json:"to"`
		Total        int                   `json:"total"`
	}

	SIPDPS3 struct {
		CurrentPage  int          `json:"current_page"`
		Data         []SIPDPSPuso `json:"data"`
		FirstPageURL string       `json:"first_page_url"`
		From         int          `json:"from"`
		LastPage     int          `json:"last_page"`
		LastPageURL  string       `json:"last_page_url"`
		NextPageURL  string       `json:"next_page_url"`
		Path         string       `json:"path"`
		PerPage      int          `json:"per_page"`
		PrevPageURL  string       `json:"prev_page_url"`
		To           int          `json:"to"`
		Total        int          `json:"total"`
	}

	SIPDPS4 struct {
		CurrentPage  int           `json:"current_page"`
		Data         []SIPDPSPanen `json:"data"`
		FirstPageURL string        `json:"first_page_url"`
		From         int           `json:"from"`
		LastPage     int           `json:"last_page"`
		LastPageURL  string        `json:"last_page_url"`
		NextPageURL  string        `json:"next_page_url"`
		Path         string        `json:"path"`
		PerPage      int           `json:"per_page"`
		PrevPageURL  string        `json:"prev_page_url"`
		To           int           `json:"to"`
		Total        int           `json:"total"`
	}
)
