package internal_api

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	"github.com/gigaflex-co/ppt_backend/util"
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

		CountRecords(c context.Context, arg util.DataFilter, id string) (int, error)

		SIPDPSTanamRead(c context.Context, id int) ([]SIPDPSTanam, error)
		SIPDPSProduktivitasRead(c context.Context, id int) ([]SIPDPSProduktivitas, error)
		SIPDPSPusoRead(c context.Context, id int) ([]SIPDPSPuso, error)
		SIPDPSPanenRead(c context.Context, id int) ([]SIPDPSPanen, error)

		PerbenihanProdusenRead(c context.Context, id int) ([]PerbenihanData4, error)
		PerbenihanRekNasRead(c context.Context, id int) ([]PerbenihanData1, error)
		PerbenihanRekBpsbRead(c context.Context, id int) ([]PerbenihanData1, error)
		PerbenihanRekLssmRead(c context.Context, id int) ([]PerbenihanData1, error)
		PerbenihanRekPenyaluranRead(c context.Context, id int) ([]PerbenihanData2, error)
		PerbenihanRekPenyebaranRead(c context.Context, id int) ([]PerbenihanData3, error)
		PerbenihanRekProdusenRead(c context.Context, id int) ([]PerbenihanData4, error)
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
		values[1] = item.NO
		values[2] = item.KODE_PROVINSI
		values[3] = item.PROVINSI
		values[4] = item.KABUPATENKOTA
		values[5] = item.KECAMATAN
		values[6] = item.KELURAHAN
		values[7] = item.USERNAME
		values[8] = item.IDSIMLUH
		values[9] = item.NOMOR_REGISTRASI
		values[10] = item.TIPE_PRODUSEN
		values[11] = item.NAMA
		values[12] = item.NAMA_PIMPINAN
		values[13] = item.ALAMAT_PIMPINAN
		values[14] = item.ALAMAT_PRODUSEN
		values[15] = item.TELEPON
		values[16] = item.EMAIL
		values[17] = item.BENIH
		values[18] = item.TOTAL_LUAS_LAHAN
		values[19] = item.LAT
		values[20] = item.LNG
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

func (q *repository) StorePerbenihanRekNasFetch(c context.Context, res []PerbenihanData1, id int) error {
	deleteQuery := "DELETE FROM " + table + " WHERE identifier = $1"

	_, err := q.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	numFields := 11
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
		values[1] = item.NO
		values[2] = item.JENIS
		values[3] = item.PROVINSI
		values[4] = item.JENIS_BENIH
		values[5] = item.KELAS_BENIH
		values[6] = item.VARIETAS
		values[7] = item.REALISASI_LUAS
		values[8] = item.REALISASI_PRODUKSI
		values[9] = item.VOLUME
		values[10] = item.DICATAT
		values[11] = item.DIPERBARUI

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

	numFields := 11
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
		values[1] = item.NO
		values[2] = item.JENIS
		values[3] = item.PROVINSI
		values[4] = item.JENIS_BENIH
		values[5] = item.KELAS_BENIH
		values[6] = item.VARIETAS
		values[7] = item.REALISASI_LUAS
		values[8] = item.REALISASI_PRODUKSI
		values[9] = item.VOLUME
		values[10] = item.DICATAT
		values[11] = item.DIPERBARUI

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

	numFields := 11
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
		values[1] = item.NO
		values[2] = item.JENIS
		values[3] = item.PROVINSI
		values[4] = item.JENIS_BENIH
		values[5] = item.KELAS_BENIH
		values[6] = item.VARIETAS
		values[7] = item.REALISASI_LUAS
		values[8] = item.REALISASI_PRODUKSI
		values[9] = item.VOLUME
		values[10] = item.DICATAT
		values[11] = item.DIPERBARUI

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

	numFields := 23
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
		values[1] = item.NO
		values[2] = item.TAHUN
		values[3] = item.BULAN
		values[4] = item.PROVINSI
		values[5] = item.KABUPATENKOTA
		values[6] = item.KECAMATAN
		values[7] = item.PRODUSEN_BENIH
		values[8] = item.KELAS_BENIH
		values[9] = item.KOMODITI
		values[10] = item.VARIETAS
		values[11] = item.STOK_LALU
		values[12] = item.PRODUKSI_BENIH
		values[13] = item.PENGADAAN
		values[14] = item.JUMLAH_STOK
		values[15] = item.PENYALURAN
		values[16] = item.APBN
		values[17] = item.APBD
		values[18] = item.FREE_MARKET
		values[19] = item.JUMLAH_SALUR
		values[20] = item.TOTAL
		values[21] = item.SISA_STOK
		values[22] = item.DICATAT
		values[23] = item.DIPERBARUI

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
		values[1] = item.NO
		values[2] = item.TAHUN
		values[3] = item.BULAN
		values[4] = item.PROVINSI
		values[5] = item.KABUPATENKOTA
		values[6] = item.KECAMATAN
		values[7] = item.KELURAHAN
		values[8] = item.PETA
		values[9] = item.REALISASI_TANAM_LUAS
		values[10] = item.BENIH
		values[11] = item.JENIS_BENIH
		values[12] = item.VARIETAS
		values[13] = item.TOTAL_LUAS
		values[14] = item.DICATAT
		values[15] = item.DIPERBARUI

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
		values[1] = item.NO
		values[2] = item.KODE_PROVINSI
		values[3] = item.PROVINSI
		values[4] = item.KABUPATENKOTA
		values[5] = item.KECAMATAN
		values[6] = item.KELURAHAN
		values[7] = item.USERNAME
		values[8] = item.IDSIMLUH
		values[9] = item.NOMOR_REGISTRASI
		values[10] = item.TIPE_PRODUSEN
		values[11] = item.NAMA
		values[12] = item.NAMA_PIMPINAN
		values[13] = item.ALAMAT_PIMPINAN
		values[14] = item.ALAMAT_PRODUSEN
		values[15] = item.TELEPON
		values[16] = item.EMAIL
		values[17] = item.BENIH
		values[18] = item.TOTAL_LUAS_LAHAN
		values[19] = item.LAT
		values[20] = item.LNG
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

func (q *repository) CountRecords(c context.Context, arg util.DataFilter, id string) (int, error) {
	args := make([]interface{}, 0)
	query := `SELECT COUNT(*) FROM ` + table + ``

	if arg.Search != "" {
		query += ` WHERE lower(` + table + `.name) LIKE CONCAT('%%',$1::text,'%%')`
		args = append(args, arg.Search)
	}

	query += ` AND identifier = ` + id + ``

	var totalRecords int
	if err := q.db.QueryRowContext(c, query, args...).Scan(&totalRecords); err != nil {
		return 0, err
	}

	return totalRecords, nil
}

func (q *repository) SIPDPSTanamRead(c context.Context, id int) ([]SIPDPSTanam, error) {
	columnNames := make([]string, 0, 22)
	for i := 1; i <= 22; i++ {
		columnNames = append(columnNames, "f"+strconv.Itoa(i))
	}

	selectedColumns := strings.Join(columnNames, ", ")

	query := "SELECT " + selectedColumns + " FROM " + table + " WHERE identifier = $1 ORDER BY f3 DESC"

	rows, err := q.db.QueryContext(c, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []SIPDPSTanam{}

	for rows.Next() {
		var r SIPDPSTanam

		if err := rows.Scan(
			&r.NIPReporter,
			&r.NamaReporter,
			&r.TanggalLaporan,
			&r.TanggalKunjungan,
			&r.JenisKelompok,
			&r.NamaProvinsi,
			&r.NamaKabupaten,
			&r.NamaKecamatan,
			&r.NamaDesa,
			&r.KategoriLahan,
			&r.JenisLahan,
			&r.JenisTanamanPangan,
			&r.NamaVarietas,
			&r.JenisBantuan,
			&r.SumberBantuan,
			&r.TahunBantuan,
			&r.LuasArea,
			&r.HST,
			&r.Latitude,
			&r.Longitude,
			&r.Photos,
			&r.Status,
		); err != nil {
			return nil, err
		}

		items = append(items, r)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (q *repository) SIPDPSProduktivitasRead(c context.Context, id int) ([]SIPDPSProduktivitas, error) {
	columnNames := make([]string, 0, 18)
	for i := 1; i <= 18; i++ {
		columnNames = append(columnNames, "f"+strconv.Itoa(i))
	}

	selectedColumns := strings.Join(columnNames, ", ")

	query := "SELECT " + selectedColumns + " FROM " + table + " WHERE identifier = $1 ORDER BY f3 DESC"

	rows, err := q.db.QueryContext(c, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []SIPDPSProduktivitas{}

	for rows.Next() {
		var r SIPDPSProduktivitas

		if err := rows.Scan(
			&r.NIPReporter,
			&r.NamaReporter,
			&r.TanggalLaporan,
			&r.TanggalKunjungan,
			&r.NamaProvinsi,
			&r.NamaKabupaten,
			&r.NamaKecamatan,
			&r.NamaDesa,
			&r.KategoriLahan,
			&r.JenisLahan,
			&r.JenisTanamanPangan,
			&r.TeknikPengukuran,
			&r.Jumlah,
			&r.Latitude,
			&r.Longitude,
			&r.Photos,
			&r.NamaVerifikator,
			&r.Status,
		); err != nil {
			return nil, err
		}

		items = append(items, r)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (q *repository) SIPDPSPusoRead(c context.Context, id int) ([]SIPDPSPuso, error) {
	columnNames := make([]string, 0, 15)
	for i := 1; i <= 15; i++ {
		columnNames = append(columnNames, "f"+strconv.Itoa(i))
	}

	selectedColumns := strings.Join(columnNames, ", ")

	query := "SELECT " + selectedColumns + " FROM " + table + " WHERE identifier = $1 ORDER BY f3 DESC"

	rows, err := q.db.QueryContext(c, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []SIPDPSPuso{}

	for rows.Next() {
		var r SIPDPSPuso

		if err := rows.Scan(
			&r.NIPReporter,
			&r.NamaReporter,
			&r.TanggalLaporan,
			&r.TanggalKejadian,
			&r.NamaProvinsi,
			&r.NamaKabupaten,
			&r.NamaKecamatan,
			&r.NamaDesa,
			&r.JenisTanamanPangan,
			&r.PenyebabPuso,
			&r.Latitude,
			&r.Longitude,
			&r.Photos,
			&r.NamaVerifikator,
			&r.Status,
		); err != nil {
			return nil, err
		}

		items = append(items, r)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (q *repository) SIPDPSPanenRead(c context.Context, id int) ([]SIPDPSPanen, error) {
	columnNames := make([]string, 0, 19)
	for i := 1; i <= 19; i++ {
		columnNames = append(columnNames, "f"+strconv.Itoa(i))
	}

	selectedColumns := strings.Join(columnNames, ", ")

	query := "SELECT " + selectedColumns + " FROM " + table + " WHERE identifier = $1 ORDER BY f3 DESC"

	rows, err := q.db.QueryContext(c, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []SIPDPSPanen{}

	for rows.Next() {
		var r SIPDPSPanen

		if err := rows.Scan(
			&r.NIPReporter,
			&r.NamaReporter,
			&r.TanggalLaporan,
			&r.TanggalKunjungan,
			&r.NamaProvinsi,
			&r.NamaKabupaten,
			&r.NamaKecamatan,
			&r.NamaDesa,
			&r.JenisTanamanPangan,
			&r.NamaVarietas,
			&r.KategoriPengelola,
			&r.NamaPengelola,
			&r.Luas,
			&r.Perkiraan,
			&r.Latitude,
			&r.Longitude,
			&r.Photos,
			&r.NamaVerifikator,
			&r.Status,
		); err != nil {
			return nil, err
		}

		items = append(items, r)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (q *repository) PerbenihanProdusenRead(c context.Context, id int) ([]PerbenihanData4, error) {
	columnNames := make([]string, 0, 22)
	for i := 1; i <= 22; i++ {
		columnNames = append(columnNames, "f"+strconv.Itoa(i))
	}

	selectedColumns := strings.Join(columnNames, ", ")

	query := "SELECT " + selectedColumns + " FROM " + table + " WHERE identifier = $1 AND LENGTH(f15) > 2 AND LENGTH(f15) < 15 AND f19 IS NOT NULL AND f20 IS NOT NULL ORDER BY f22 DESC LIMIT 1000"

	rows, err := q.db.QueryContext(c, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []PerbenihanData4{}

	for rows.Next() {
		var r PerbenihanData4

		if err := rows.Scan(
			&r.NO,
			&r.KODE_PROVINSI,
			&r.PROVINSI,
			&r.KABUPATENKOTA,
			&r.KECAMATAN,
			&r.KELURAHAN,
			&r.USERNAME,
			&r.IDSIMLUH,
			&r.NOMOR_REGISTRASI,
			&r.TIPE_PRODUSEN,
			&r.NAMA,
			&r.NAMA_PIMPINAN,
			&r.ALAMAT_PIMPINAN,
			&r.ALAMAT_PRODUSEN,
			&r.TELEPON,
			&r.EMAIL,
			&r.BENIH,
			&r.TOTAL_LUAS_LAHAN,
			&r.LAT,
			&r.LNG,
			&r.DICATAT,
			&r.DIPERBARUI,
		); err != nil {
			return nil, err
		}

		items = append(items, r)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (q *repository) PerbenihanRekNasRead(c context.Context, id int) ([]PerbenihanData1, error) {
	columnNames := make([]string, 0, 11)
	for i := 1; i <= 11; i++ {
		columnNames = append(columnNames, "f"+strconv.Itoa(i))
	}

	selectedColumns := strings.Join(columnNames, ", ")

	query := "SELECT " + selectedColumns + " FROM " + table + " WHERE identifier = $1 ORDER BY f11 DESC"

	rows, err := q.db.QueryContext(c, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []PerbenihanData1{}

	for rows.Next() {
		var r PerbenihanData1

		if err := rows.Scan(
			&r.NO,
			&r.JENIS,
			&r.PROVINSI,
			&r.JENIS_BENIH,
			&r.KELAS_BENIH,
			&r.VARIETAS,
			&r.REALISASI_LUAS,
			&r.REALISASI_PRODUKSI,
			&r.VOLUME,
			&r.DICATAT,
			&r.DIPERBARUI,
		); err != nil {
			return nil, err
		}

		items = append(items, r)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (q *repository) PerbenihanRekBpsbRead(c context.Context, id int) ([]PerbenihanData1, error) {
	columnNames := make([]string, 0, 11)
	for i := 1; i <= 11; i++ {
		columnNames = append(columnNames, "f"+strconv.Itoa(i))
	}

	selectedColumns := strings.Join(columnNames, ", ")

	query := "SELECT " + selectedColumns + " FROM " + table + " WHERE identifier = $1 ORDER BY f11 DESC"

	rows, err := q.db.QueryContext(c, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []PerbenihanData1{}

	for rows.Next() {
		var r PerbenihanData1

		if err := rows.Scan(
			&r.NO,
			&r.JENIS,
			&r.PROVINSI,
			&r.JENIS_BENIH,
			&r.KELAS_BENIH,
			&r.VARIETAS,
			&r.REALISASI_LUAS,
			&r.REALISASI_PRODUKSI,
			&r.VOLUME,
			&r.DICATAT,
			&r.DIPERBARUI,
		); err != nil {
			return nil, err
		}

		items = append(items, r)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (q *repository) PerbenihanRekLssmRead(c context.Context, id int) ([]PerbenihanData1, error) {
	columnNames := make([]string, 0, 11)
	for i := 1; i <= 11; i++ {
		columnNames = append(columnNames, "f"+strconv.Itoa(i))
	}

	selectedColumns := strings.Join(columnNames, ", ")

	query := "SELECT " + selectedColumns + " FROM " + table + " WHERE identifier = $1 ORDER BY f11 DESC"

	rows, err := q.db.QueryContext(c, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []PerbenihanData1{}

	for rows.Next() {
		var r PerbenihanData1

		if err := rows.Scan(
			&r.NO,
			&r.JENIS,
			&r.PROVINSI,
			&r.JENIS_BENIH,
			&r.KELAS_BENIH,
			&r.VARIETAS,
			&r.REALISASI_LUAS,
			&r.REALISASI_PRODUKSI,
			&r.VOLUME,
			&r.DICATAT,
			&r.DIPERBARUI,
		); err != nil {
			return nil, err
		}

		items = append(items, r)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (q *repository) PerbenihanRekPenyaluranRead(c context.Context, id int) ([]PerbenihanData2, error) {
	columnNames := make([]string, 0, 23)
	for i := 1; i <= 23; i++ {
		columnNames = append(columnNames, "f"+strconv.Itoa(i))
	}

	selectedColumns := strings.Join(columnNames, ", ")

	query := "SELECT " + selectedColumns + " FROM " + table + " WHERE identifier = $1 ORDER BY f11 DESC"

	rows, err := q.db.QueryContext(c, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []PerbenihanData2{}

	for rows.Next() {
		var r PerbenihanData2

		if err := rows.Scan(
			&r.NO,
			&r.TAHUN,
			&r.BULAN,
			&r.PROVINSI,
			&r.KABUPATENKOTA,
			&r.KECAMATAN,
			&r.PRODUSEN_BENIH,
			&r.KELAS_BENIH,
			&r.KOMODITI,
			&r.VARIETAS,
			&r.STOK_LALU,
			&r.PRODUKSI_BENIH,
			&r.PENGADAAN,
			&r.JUMLAH_STOK,
			&r.PENYALURAN,
			&r.APBN,
			&r.APBD,
			&r.FREE_MARKET,
			&r.JUMLAH_SALUR,
			&r.TOTAL,
			&r.SISA_STOK,
			&r.DICATAT,
			&r.DIPERBARUI,
		); err != nil {
			return nil, err
		}

		items = append(items, r)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (q *repository) PerbenihanRekPenyebaranRead(c context.Context, id int) ([]PerbenihanData3, error) {
	columnNames := make([]string, 0, 15)
	for i := 1; i <= 15; i++ {
		columnNames = append(columnNames, "f"+strconv.Itoa(i))
	}

	selectedColumns := strings.Join(columnNames, ", ")

	query := "SELECT " + selectedColumns + " FROM " + table + " WHERE identifier = $1 ORDER BY f11 DESC"

	rows, err := q.db.QueryContext(c, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []PerbenihanData3{}

	for rows.Next() {
		var r PerbenihanData3

		if err := rows.Scan(
			&r.NO,
			&r.TAHUN,
			&r.BULAN,
			&r.PROVINSI,
			&r.KABUPATENKOTA,
			&r.KECAMATAN,
			&r.KELURAHAN,
			&r.PETA,
			&r.REALISASI_TANAM_LUAS,
			&r.BENIH,
			&r.JENIS_BENIH,
			&r.VARIETAS,
			&r.TOTAL_LUAS,
			&r.DICATAT,
			&r.DIPERBARUI,
		); err != nil {
			return nil, err
		}

		items = append(items, r)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (q *repository) PerbenihanRekProdusenRead(c context.Context, id int) ([]PerbenihanData4, error) {
	columnNames := make([]string, 0, 22)
	for i := 1; i <= 22; i++ {
		columnNames = append(columnNames, "f"+strconv.Itoa(i))
	}

	selectedColumns := strings.Join(columnNames, ", ")

	query := "SELECT " + selectedColumns + " FROM " + table + " WHERE identifier = $1 ORDER BY f11 DESC"

	rows, err := q.db.QueryContext(c, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []PerbenihanData4{}

	for rows.Next() {
		var r PerbenihanData4

		if err := rows.Scan(
			&r.NO,
			&r.KODE_PROVINSI,
			&r.PROVINSI,
			&r.KABUPATENKOTA,
			&r.KECAMATAN,
			&r.KELURAHAN,
			&r.USERNAME,
			&r.IDSIMLUH,
			&r.NOMOR_REGISTRASI,
			&r.TIPE_PRODUSEN,
			&r.NAMA,
			&r.NAMA_PIMPINAN,
			&r.ALAMAT_PIMPINAN,
			&r.ALAMAT_PRODUSEN,
			&r.TELEPON,
			&r.EMAIL,
			&r.BENIH,
			&r.TOTAL_LUAS_LAHAN,
			&r.LAT,
			&r.LNG,
			&r.DICATAT,
			&r.DIPERBARUI,
		); err != nil {
			return nil, err
		}

		items = append(items, r)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
