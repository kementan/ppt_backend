package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gigaflex-co/ppt_backend/util"
)

type (
	ServiceRepository interface {
		GetDataBy(c context.Context, field string, value string) (ServiceResponse, error)
		GetTable(c context.Context, arg util.DataFilter) ([]ServiceResponse, error)
		CountRecords(c context.Context, arg util.DataFilter) (int, error)
		Read(c context.Context) ([]ServiceResponse, error)
		Create(c context.Context, arg ServiceCreate) (ServiceResponse, error)
		Update(c context.Context, id string, arg ServiceUpdate) (ServiceResponse, error)
		Delete(c context.Context, ids []string) ([]string, []string, error)
	}

	repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) ServiceRepository {
	return &repository{
		db: db,
	}
}

func (q *repository) GetDataBy(c context.Context, field string, value string) (ServiceResponse, error) {
	var r ServiceResponse
	var enc_id string

	query := `
	SELECT * FROM ` + table + `
	WHERE ` + field + ` = $1
	LIMIT 1`

	row := q.db.QueryRowContext(c, query, value)
	err := row.Scan(
		&enc_id,
		&r.Name,
		&r.Slug,
		&r.Sort,
		&r.IsActive,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return r, errors.New("record not found")
		}
		return r, err
	}

	r.HashedID, _ = util.Encrypt(enc_id, "f")

	return r, err
}

func (q *repository) GetTable(c context.Context, arg util.DataFilter) ([]ServiceResponse, error) {
	offset := (arg.Page - 1) * arg.PageSize
	items := []ServiceResponse{}
	args := make([]interface{}, 0)
	args = append(args, arg.PageSize, offset)

	query := `
		SELECT 
		` + table + `.id, 
		` + table + `.name, 
		` + table + `.slug, 
		` + table + `.sort, 
		` + table + `.is_active, 
		COUNT(` + subtable + `.id) AS sub_service, 
		` + table + `.created_at, 
		` + table + `.updated_at
		FROM ` + table + `
		LEFT JOIN ` + subtable + ` ON ` + subtable + `.service_id = ` + table + `.id
		GROUP BY ` + table + `.id
		`

	if arg.Search != "" {
		query = query + ` WHERE lower(` + table + `.name) LIKE CONCAT('%%',$3::text,'%%')`
		args = append(args, arg.Search)
	}

	query = query + `
		ORDER BY ` + arg.SortBy + ` ` + arg.SortOrder + `
		LIMIT $1 OFFSET $2`

	rows, err := q.db.QueryContext(c, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var r ServiceResponse
		var enc_id string

		if err := rows.Scan(
			&enc_id,
			&r.Name,
			&r.Slug,
			&r.Sort,
			&r.IsActive,
			&r.SubService,
			&r.CreatedAt,
			&r.UpdatedAt,
		); err != nil {
			return nil, err
		}

		r.HashedID, _ = util.Encrypt(enc_id, "f")

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

func (q *repository) CountRecords(c context.Context, arg util.DataFilter) (int, error) {
	args := make([]interface{}, 0)
	query := `SELECT COUNT(*) FROM ` + table + ``

	if arg.Search != "" {
		query += ` WHERE lower(` + table + `.name) LIKE CONCAT('%%',$1::text,'%%')`
		args = append(args, arg.Search)
	}

	var totalRecords int
	if err := q.db.QueryRowContext(c, query, args...).Scan(&totalRecords); err != nil {
		return 0, err
	}

	return totalRecords, nil
}

func (q *repository) Read(c context.Context) ([]ServiceResponse, error) {
	items := []ServiceResponse{}

	query := `
		SELECT 
			id, name, slug, sort, is_active, created_at, updated_at
		FROM ` + table + `
		ORDER BY id ASC
		`

	rows, err := q.db.QueryContext(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r ServiceResponse
		var enc_id string

		if err := rows.Scan(
			&enc_id,
			&r.Name,
			&r.Slug,
			&r.Sort,
			&r.IsActive,
			&r.CreatedAt,
			&r.UpdatedAt,
		); err != nil {
			return nil, err
		}

		r.HashedID, _ = util.Encrypt(enc_id, "f")

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

func (q *repository) Create(c context.Context, arg ServiceCreate) (ServiceResponse, error) {
	var r ServiceResponse
	var enc_id string

	query := `
	INSERT INTO ` + table + ` (
		name, slug, sort, is_active
	) VALUES ($1, $2, $3, $4)
	RETURNING *`

	row := q.db.QueryRowContext(c, query, arg.Name, arg.Slug, arg.Sort, arg.IsActive)

	err := row.Scan(
		&enc_id,
		&r.Name,
		&r.Slug,
		&r.Sort,
		&r.IsActive,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	r.HashedID, _ = util.Encrypt(enc_id, "f")

	return r, err
}

func (q *repository) Update(c context.Context, id string, arg ServiceUpdate) (ServiceResponse, error) {
	var r ServiceResponse

	decryptedID, _ := util.Decrypt(id, "f")

	query := `
	UPDATE ` + table + `
	SET
		name = COALESCE($2, name),
		slug = COALESCE($3, slug),
		sort = COALESCE($4, sort),
		is_active = COALESCE($5, is_active)
	WHERE id = $1
	RETURNING *`

	row := q.db.QueryRowContext(c, query,
		decryptedID,
		arg.Name,
		arg.Slug,
		arg.Sort,
		arg.IsActive,
	)

	err := row.Scan(
		&r.HashedID,
		&r.Name,
		&r.Slug,
		&r.Sort,
		&r.IsActive,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	r.HashedID = id

	return r, err
}

func (q *repository) Delete(c context.Context, ids []string) ([]string, []string, error) {
	var successIDs, failedIDs []string

	for _, id := range ids {
		decryptedID, _ := util.Decrypt(id, "f")
		query := `DELETE FROM ` + table + ` WHERE id = $1`
		_, err := q.db.ExecContext(c, query, decryptedID)
		if err != nil {
			failedIDs = append(failedIDs, id)
		} else {
			successIDs = append(successIDs, id)
		}
	}

	return successIDs, failedIDs, nil
}
