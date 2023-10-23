package region

import (
	"context"
	"database/sql"
	"errors"
)

type (
	RegionRepository interface {
		GetList(c context.Context, level, parentCode string) ([]Wilayah, error)
		GetRegion(c context.Context, code string) (Wilayah, error)
	}

	repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) RegionRepository {
	return &repository{
		db: db,
	}
}

func (q *repository) GetList(c context.Context, level, parentCode string) ([]Wilayah, error) {
	var r Wilayah
	var query string
	args := make([]interface{}, 0)
	items := []Wilayah{}

	switch level {
	case "1":
		query = `
		SELECT kode, nama, bmkg FROM ` + table + `
		WHERE LENGTH(kode) = 2`
	case "2":
		query = `
		SELECT kode, nama, bmkg FROM ` + table + `
		WHERE LENGTH(kode) = 5
		AND kode LIKE $1`
		args = append(args, parentCode[:2]+"%")
	case "3":
		query = `
		SELECT kode, nama, bmkg FROM ` + table + `
		WHERE LENGTH(kode) = 8
		AND kode LIKE $1`
		args = append(args, parentCode[:5]+"%")
	case "4":
		query = `
		SELECT kode, nama, bmkg FROM ` + table + `
		WHERE LENGTH(kode) = 13
		AND kode LIKE $1`
		args = append(args, parentCode[:8]+"%")
	}

	rows, err := q.db.QueryContext(c, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&r.Kode,
			&r.Nama,
			&r.BMKG,
		); err != nil {
			return nil, err
		}
		items = append(items, r)
	}

	return items, err
}

func (q *repository) GetRegion(c context.Context, code string) (Wilayah, error) {
	var r Wilayah

	query := `
		SELECT kode, nama, bmkg FROM ` + table + `
		WHERE kode = $1`

	row := q.db.QueryRowContext(c, query, code)

	err := row.Scan(
		&r.Kode,
		&r.Nama,
		&r.BMKG,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return r, errors.New("record not found")
		}
		return r, err
	}

	return r, err
}
