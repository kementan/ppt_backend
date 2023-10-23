package internal_api

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
)

type (
	InternalApiRepository interface {
		GetAll(c context.Context) ([]PptApiDataStorage, error)
		GetToken(c context.Context, key string) (string, error)
		StoreSIPDPSTanamFetch(c context.Context, res []SIPDPSTanam, id int) error
		StoreSIPDPSProduktivitasFetch(c context.Context, res []SIPDPSProduktivitas, id int) error
		StoreSIPDPSPusoFetch(c context.Context, res []SIPDPSPuso, id int) error
		StoreSIPDPSPanenFetch(c context.Context, res []SIPDPSPanen, id int) error

		StorePerbenihanProdusenFetch(c context.Context, res []PerbenihanData4, id int) error
		StorePerbenihanRekNasFetch(c context.Context, res []PerbenihanData1, id int) error
		StorePerbenihanRekBpsbFetch(c context.Context, res []PerbenihanData1, id int) error
		StorePerbenihanRekLssmFetch(c context.Context, res []PerbenihanData1, id int) error
		StorePerbenihanRekPenyaluranFetch(c context.Context, res []PerbenihanData2, id int) error
		StorePerbenihanRekPenyebaranFetch(c context.Context, res []PerbenihanData3, id int) error
		StorePerbenihanRekProdusenFetch(c context.Context, res []PerbenihanData4, id int) error
	}

	repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) InternalApiRepository {
	return &repository{
		db: db,
	}
}

func (q *repository) GetAll(c context.Context) ([]PptApiDataStorage, error) {
	query := "SELECT * FROM ppt_api_data_storages"
	rows, err := q.db.QueryContext(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []PptApiDataStorage
	for rows.Next() {
		var row PptApiDataStorage
		err := rows.Scan(
			&row.ID,
			&row.Identifier,
			&row.F1,
			&row.F2,
			&row.F3,
			&row.F4,
			&row.F5,
			&row.F6,
			&row.F7,
			&row.F8,
			&row.F9,
			&row.F10,
			&row.F11,
			&row.F12,
			&row.F13,
			&row.F14,
			&row.F15,
			&row.F16,
			&row.F17,
			&row.F18,
			&row.F19,
			&row.F20,
			&row.F21,
			&row.F22,
			&row.F23,
			&row.F24,
			&row.F25,
			&row.LongText,
			&row.CreatedAt,
			&row.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, row)
	}
	return data, nil
}

func (q *repository) GetToken(c context.Context, key string) (string, error) {
	query := "SELECT value FROM ppt_configurations WHERE name = $1"

	// Execute the query and scan the result.
	var value string
	err := q.db.QueryRowContext(c, query, key).Scan(&value)
	if err != nil {
		return "", err
	}

	return value, nil
}

func (q *repository) StoreSIPDPSTanamFetch(c context.Context, res []SIPDPSTanam, id int) error {
	deleteQuery := "DELETE FROM " + table + " WHERE identifier = $1"

	_, err := q.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	numFields := 22
	placeholders := make([]string, numFields)
	columns := make([]string, numFields)
	values := make([]interface{}, numFields+1)

	for i := 0; i < numFields; i++ {
		placeholders[i] = "$" + strconv.Itoa(i+2)
		columns[i] = "f" + strconv.Itoa(i+1)
	}

	query := "INSERT INTO " + table + " (identifier, " + strings.Join(columns, ", ") + ") " +
		"VALUES ($1, " + strings.Join(placeholders, ", ") + ")"

	stmt, err := q.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	tx, err := q.db.Begin()
	if err != nil {
		return err
	}

	for _, item := range res {
		values[0] = id
		values[1] = item.NIPReporter
		values[2] = item.NamaReporter
		values[3] = item.TanggalLaporan
		values[4] = item.TanggalKunjungan
		values[5] = item.JenisKelompok
		values[6] = item.NamaProvinsi
		values[7] = item.NamaKabupaten
		values[8] = item.NamaKecamatan
		values[9] = item.NamaDesa
		values[10] = item.KategoriLahan
		values[11] = item.JenisLahan
		values[12] = item.JenisTanamanPangan
		values[13] = item.NamaVarietas
		values[14] = item.JenisBantuan
		values[15] = item.SumberBantuan
		values[16] = item.TahunBantuan
		values[17] = item.LuasArea
		values[18] = item.HST
		values[19] = item.Latitude
		values[20] = item.Longitude
		values[21] = item.Photos
		values[22] = item.Status

		_, err := tx.Stmt(stmt).Exec(values...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (q *repository) StoreSIPDPSProduktivitasFetch(c context.Context, res []SIPDPSProduktivitas, id int) error {
	deleteQuery := "DELETE FROM " + table + " WHERE identifier = $1"

	_, err := q.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	numFields := 18
	placeholders := make([]string, numFields)
	columns := make([]string, numFields)
	values := make([]interface{}, numFields+1)

	for i := 0; i < numFields; i++ {
		placeholders[i] = "$" + strconv.Itoa(i+2)
		columns[i] = "f" + strconv.Itoa(i+1)
	}

	query := "INSERT INTO " + table + " (identifier, " + strings.Join(columns, ", ") + ") " +
		"VALUES ($1, " + strings.Join(placeholders, ", ") + ")"

	stmt, err := q.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	tx, err := q.db.Begin()
	if err != nil {
		return err
	}

	for _, item := range res {
		values[0] = id
		values[1] = item.NIPReporter
		values[2] = item.NamaReporter
		values[3] = item.TanggalLaporan
		values[4] = item.TanggalKunjungan
		values[5] = item.NamaProvinsi
		values[6] = item.NamaKabupaten
		values[7] = item.NamaKecamatan
		values[8] = item.NamaDesa
		values[9] = item.KategoriLahan
		values[10] = item.JenisLahan
		values[11] = item.JenisTanamanPangan
		values[12] = item.TeknikPengukuran
		values[13] = item.Jumlah
		values[14] = item.Latitude
		values[15] = item.Longitude
		values[16] = item.Photos
		values[17] = item.NamaVerifikator
		values[18] = item.Status

		_, err := tx.Stmt(stmt).Exec(values...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (q *repository) StoreSIPDPSPusoFetch(c context.Context, res []SIPDPSPuso, id int) error {
	deleteQuery := "DELETE FROM " + table + " WHERE identifier = $1"

	_, err := q.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	numFields := 15
	placeholders := make([]string, numFields)
	columns := make([]string, numFields)
	values := make([]interface{}, numFields+1)

	for i := 0; i < numFields; i++ {
		placeholders[i] = "$" + strconv.Itoa(i+2)
		columns[i] = "f" + strconv.Itoa(i+1)
	}

	query := "INSERT INTO " + table + " (identifier, " + strings.Join(columns, ", ") + ") " +
		"VALUES ($1, " + strings.Join(placeholders, ", ") + ")"

	stmt, err := q.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	tx, err := q.db.Begin()
	if err != nil {
		return err
	}

	for _, item := range res {
		values[0] = id
		values[1] = item.NIPReporter
		values[2] = item.NamaReporter
		values[3] = item.TanggalLaporan
		values[4] = item.TanggalKejadian
		values[5] = item.NamaProvinsi
		values[6] = item.NamaKabupaten
		values[7] = item.NamaKecamatan
		values[8] = item.NamaDesa
		values[9] = item.JenisTanamanPangan
		values[10] = item.PenyebabPuso
		values[11] = item.Latitude
		values[12] = item.Longitude
		values[13] = item.Photos
		values[14] = item.NamaVerifikator
		values[15] = item.Status

		_, err := tx.Stmt(stmt).Exec(values...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (q *repository) StoreSIPDPSPanenFetch(c context.Context, res []SIPDPSPanen, id int) error {
	deleteQuery := "DELETE FROM " + table + " WHERE identifier = $1"

	_, err := q.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	numFields := 19
	placeholders := make([]string, numFields)
	columns := make([]string, numFields)
	values := make([]interface{}, numFields+1)

	for i := 0; i < numFields; i++ {
		placeholders[i] = "$" + strconv.Itoa(i+2)
		columns[i] = "f" + strconv.Itoa(i+1)
	}

	query := "INSERT INTO " + table + " (identifier, " + strings.Join(columns, ", ") + ") " +
		"VALUES ($1, " + strings.Join(placeholders, ", ") + ")"

	stmt, err := q.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	tx, err := q.db.Begin()
	if err != nil {
		return err
	}

	for _, item := range res {
		values[0] = id
		values[1] = item.NIPReporter
		values[2] = item.NamaReporter
		values[3] = item.TanggalLaporan
		values[4] = item.TanggalKunjungan
		values[5] = item.NamaProvinsi
		values[6] = item.NamaKabupaten
		values[7] = item.NamaKecamatan
		values[8] = item.NamaDesa
		values[9] = item.JenisTanamanPangan
		values[10] = item.NamaVarietas
		values[11] = item.KategoriPengelola
		values[12] = item.NamaPengelola
		values[13] = item.Luas
		values[14] = item.Perkiraan
		values[15] = item.Latitude
		values[16] = item.Longitude
		values[17] = item.Photos
		values[18] = item.NamaVerifikator
		values[19] = item.Status

		_, err := tx.Stmt(stmt).Exec(values...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (q *repository) StorePerbenihanProdusenFetch(c context.Context, res []PerbenihanData4, id int) error {
	deleteQuery := "DELETE FROM " + table + " WHERE identifier = $1"

	_, err := q.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	numFields := 21
	placeholders := make([]string, numFields)
	columns := make([]string, numFields)
	values := make([]interface{}, numFields+1)

	for i := 0; i < numFields; i++ {
		placeholders[i] = "$" + strconv.Itoa(i+2)
		columns[i] = "f" + strconv.Itoa(i+1)
	}

	query := "INSERT INTO " + table + " (identifier, " + strings.Join(columns, ", ") + ") " +
		"VALUES ($1, " + strings.Join(placeholders, ", ") + ")"

	stmt, err := q.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	tx, err := q.db.Begin()
	if err != nil {
		return err
	}

	for _, item := range res {
		values[0] = id
		values[1] = item.KODE_PROVINSI
		values[2] = item.PROVINSI
		values[3] = item.KABUPATENKOTA
		values[4] = item.KECAMATAN
		values[5] = item.KELURAHAN
		values[6] = item.USERNAME
		values[7] = item.IDSIMLUH
		values[8] = item.NOMOR_REGISTRASI
		values[9] = item.TIPE_PRODUSEN
		values[10] = item.NAMA
		values[11] = item.NAMA_PIMPINAN
		values[12] = item.ALAMAT_PIMPINAN
		values[13] = item.ALAMAT_PRODUSEN
		values[14] = item.TELEPON
		values[15] = item.EMAIL
		values[16] = item.BENIH
		values[17] = item.TOTAL_LUAS_LAHAN
		values[18] = item.LAT
		values[19] = item.LNG
		values[20] = item.DICATAT
		values[21] = item.DIPERBARUI

		_, err := tx.Stmt(stmt).Exec(values...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (q *repository) StorePerbenihanRekNasFetch(c context.Context, res []PerbenihanData1, id int) error {
	deleteQuery := "DELETE FROM " + table + " WHERE identifier = $1"

	_, err := q.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	numFields := 10
	placeholders := make([]string, numFields)
	columns := make([]string, numFields)
	values := make([]interface{}, numFields+1)

	for i := 0; i < numFields; i++ {
		placeholders[i] = "$" + strconv.Itoa(i+2)
		columns[i] = "f" + strconv.Itoa(i+1)
	}

	query := "INSERT INTO " + table + " (identifier, " + strings.Join(columns, ", ") + ") " +
		"VALUES ($1, " + strings.Join(placeholders, ", ") + ")"

	stmt, err := q.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	tx, err := q.db.Begin()
	if err != nil {
		return err
	}

	for _, item := range res {
		values[0] = id
		values[1] = item.JENIS
		values[2] = item.PROVINSI
		values[3] = item.JENIS_BENIH
		values[4] = item.KELAS_BENIH
		values[5] = item.VARIETAS
		values[6] = item.REALISASI_LUAS
		values[7] = item.REALISASI_PRODUKSI
		values[8] = item.VOLUME
		values[9] = item.DICATAT
		values[10] = item.DIPERBARUI

		_, err := tx.Stmt(stmt).Exec(values...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (q *repository) StorePerbenihanRekBpsbFetch(c context.Context, res []PerbenihanData1, id int) error {
	deleteQuery := "DELETE FROM " + table + " WHERE identifier = $1"

	_, err := q.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	numFields := 10
	placeholders := make([]string, numFields)
	columns := make([]string, numFields)
	values := make([]interface{}, numFields+1)

	for i := 0; i < numFields; i++ {
		placeholders[i] = "$" + strconv.Itoa(i+2)
		columns[i] = "f" + strconv.Itoa(i+1)
	}

	query := "INSERT INTO " + table + " (identifier, " + strings.Join(columns, ", ") + ") " +
		"VALUES ($1, " + strings.Join(placeholders, ", ") + ")"

	stmt, err := q.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	tx, err := q.db.Begin()
	if err != nil {
		return err
	}

	for _, item := range res {
		values[0] = id
		values[1] = item.JENIS
		values[2] = item.PROVINSI
		values[3] = item.JENIS_BENIH
		values[4] = item.KELAS_BENIH
		values[5] = item.VARIETAS
		values[6] = item.REALISASI_LUAS
		values[7] = item.REALISASI_PRODUKSI
		values[8] = item.VOLUME
		values[9] = item.DICATAT
		values[10] = item.DIPERBARUI

		_, err := tx.Stmt(stmt).Exec(values...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (q *repository) StorePerbenihanRekLssmFetch(c context.Context, res []PerbenihanData1, id int) error {
	deleteQuery := "DELETE FROM " + table + " WHERE identifier = $1"

	_, err := q.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	numFields := 10
	placeholders := make([]string, numFields)
	columns := make([]string, numFields)
	values := make([]interface{}, numFields+1)

	for i := 0; i < numFields; i++ {
		placeholders[i] = "$" + strconv.Itoa(i+2)
		columns[i] = "f" + strconv.Itoa(i+1)
	}

	query := "INSERT INTO " + table + " (identifier, " + strings.Join(columns, ", ") + ") " +
		"VALUES ($1, " + strings.Join(placeholders, ", ") + ")"

	stmt, err := q.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	tx, err := q.db.Begin()
	if err != nil {
		return err
	}

	for _, item := range res {
		values[0] = id
		values[1] = item.JENIS
		values[2] = item.PROVINSI
		values[3] = item.JENIS_BENIH
		values[4] = item.KELAS_BENIH
		values[5] = item.VARIETAS
		values[6] = item.REALISASI_LUAS
		values[7] = item.REALISASI_PRODUKSI
		values[8] = item.VOLUME
		values[9] = item.DICATAT
		values[10] = item.DIPERBARUI

		_, err := tx.Stmt(stmt).Exec(values...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (q *repository) StorePerbenihanRekPenyaluranFetch(c context.Context, res []PerbenihanData2, id int) error {
	deleteQuery := "DELETE FROM " + table + " WHERE identifier = $1"

	_, err := q.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	numFields := 22
	placeholders := make([]string, numFields)
	columns := make([]string, numFields)
	values := make([]interface{}, numFields+1)

	for i := 0; i < numFields; i++ {
		placeholders[i] = "$" + strconv.Itoa(i+2)
		columns[i] = "f" + strconv.Itoa(i+1)
	}

	query := "INSERT INTO " + table + " (identifier, " + strings.Join(columns, ", ") + ") " +
		"VALUES ($1, " + strings.Join(placeholders, ", ") + ")"

	stmt, err := q.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	tx, err := q.db.Begin()
	if err != nil {
		return err
	}

	for _, item := range res {
		values[0] = id
		values[1] = item.TAHUN
		values[2] = item.BULAN
		values[3] = item.PROVINSI
		values[4] = item.KABUPATENKOTA
		values[5] = item.KECAMATAN
		values[6] = item.PRODUSEN_BENIH
		values[7] = item.KELAS_BENIH
		values[8] = item.KOMODITI
		values[9] = item.VARIETAS
		values[10] = item.STOK_LALU
		values[11] = item.PRODUKSI_BENIH
		values[12] = item.PENGADAAN
		values[13] = item.JUMLAH_STOK
		values[14] = item.PENYALURAN
		values[15] = item.APBN
		values[16] = item.APBD
		values[17] = item.FREE_MARKET
		values[18] = item.JUMLAH_SALUR
		values[19] = item.TOTAL
		values[20] = item.SISA_STOK
		values[21] = item.DICATAT
		values[22] = item.DIPERBARUI

		_, err := tx.Stmt(stmt).Exec(values...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (q *repository) StorePerbenihanRekPenyebaranFetch(c context.Context, res []PerbenihanData3, id int) error {
	deleteQuery := "DELETE FROM " + table + " WHERE identifier = $1"

	_, err := q.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	numFields := 14
	placeholders := make([]string, numFields)
	columns := make([]string, numFields)
	values := make([]interface{}, numFields+1)

	for i := 0; i < numFields; i++ {
		placeholders[i] = "$" + strconv.Itoa(i+2)
		columns[i] = "f" + strconv.Itoa(i+1)
	}

	query := "INSERT INTO " + table + " (identifier, " + strings.Join(columns, ", ") + ") " +
		"VALUES ($1, " + strings.Join(placeholders, ", ") + ")"

	stmt, err := q.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	tx, err := q.db.Begin()
	if err != nil {
		return err
	}

	for _, item := range res {
		values[0] = id
		values[1] = item.TAHUN
		values[2] = item.BULAN
		values[3] = item.PROVINSI
		values[4] = item.KABUPATENKOTA
		values[5] = item.KECAMATAN
		values[6] = item.KELURAHAN
		values[7] = item.PETA
		values[8] = item.REALISASI_TANAM_LUAS
		values[9] = item.BENIH
		values[10] = item.JENIS_BENIH
		values[11] = item.VARIETAS
		values[12] = item.TOTAL_LUAS
		values[13] = item.DICATAT
		values[14] = item.DIPERBARUI

		_, err := tx.Stmt(stmt).Exec(values...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (q *repository) StorePerbenihanRekProdusenFetch(c context.Context, res []PerbenihanData4, id int) error {
	deleteQuery := "DELETE FROM " + table + " WHERE identifier = $1"

	_, err := q.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	numFields := 21
	placeholders := make([]string, numFields)
	columns := make([]string, numFields)
	values := make([]interface{}, numFields+1)

	for i := 0; i < numFields; i++ {
		placeholders[i] = "$" + strconv.Itoa(i+2)
		columns[i] = "f" + strconv.Itoa(i+1)
	}

	query := "INSERT INTO " + table + " (identifier, " + strings.Join(columns, ", ") + ") " +
		"VALUES ($1, " + strings.Join(placeholders, ", ") + ")"

	stmt, err := q.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	tx, err := q.db.Begin()
	if err != nil {
		return err
	}

	for _, item := range res {
		values[0] = id
		values[1] = item.KODE_PROVINSI
		values[2] = item.PROVINSI
		values[3] = item.KABUPATENKOTA
		values[4] = item.KECAMATAN
		values[5] = item.KELURAHAN
		values[6] = item.USERNAME
		values[7] = item.IDSIMLUH
		values[8] = item.NOMOR_REGISTRASI
		values[9] = item.TIPE_PRODUSEN
		values[10] = item.NAMA
		values[11] = item.NAMA_PIMPINAN
		values[12] = item.ALAMAT_PIMPINAN
		values[13] = item.ALAMAT_PRODUSEN
		values[14] = item.TELEPON
		values[15] = item.EMAIL
		values[16] = item.BENIH
		values[17] = item.TOTAL_LUAS_LAHAN
		values[18] = item.LAT
		values[19] = item.LNG
		values[20] = item.DICATAT
		values[21] = item.DIPERBARUI

		_, err := tx.Stmt(stmt).Exec(values...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
