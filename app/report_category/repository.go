package report_category

import (
	"context"
	"database/sql"

	"github.com/gigaflex-co/ppt_backend/util"
)

type (
	ReportCategoryRepository interface {
		GetDataBy(c context.Context, field string, value string) (ReportCategoryResponse, error)
		GetTable(c context.Context, arg util.DataFilter) ([]ReportCategoryResponse, error)
		CountRecords(c context.Context, arg util.DataFilter) (int, error)
		Create(c context.Context, arg ReportCategoryCreate) (ReportCategoryResponse, error)
		Read(c context.Context) ([]ReportCategoryResponse, error)
		Update(c context.Context, id string, arg ReportCategoryUpdate) (ReportCategoryResponse, error)
		Delete(c context.Context, ids []string) ([]string, []string, error)
	}

	repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) ReportCategoryRepository {
	return &repository{
		db: db,
	}
}

func (q *repository) GetDataBy(c context.Context, field string, value string) (ReportCategoryResponse, error) {
	var r ReportCategoryResponse
	var enc_id string

	query := `
	SELECT id, name, created_at, updated_at FROM ` + table + `
	WHERE ` + field + ` = $1
	LIMIT 1`

	row := q.db.QueryRowContext(c, query, value)
	err := row.Scan(
		&enc_id,
		&r.Name,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	r.HashedID, _ = util.Encrypt(enc_id, "f")

	return r, err
}

func (q *repository) GetTable(c context.Context, arg util.DataFilter) ([]ReportCategoryResponse, error) {
	offset := (arg.Page - 1) * arg.PageSize
	items := []ReportCategoryResponse{}
	args := make([]interface{}, 0)
	args = append(args, arg.PageSize, offset)

	query := `
		SELECT id, name, created_at, updated_at
		FROM ` + table + `
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
		var r ReportCategoryResponse
		var enc_id string

		if err := rows.Scan(
			&enc_id,
			&r.Name,
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

func (q *repository) Create(c context.Context, arg ReportCategoryCreate) (ReportCategoryResponse, error) {
	var r ReportCategoryResponse
	var enc_id string

	query := `
	INSERT INTO ` + table + ` (
		name
	) VALUES ($1)
	RETURNING id, name, created_at, updated_at`

	row := q.db.QueryRowContext(c, query, arg.Name)

	err := row.Scan(
		&enc_id,
		&r.Name,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	r.HashedID, _ = util.Encrypt(enc_id, "f")

	return r, err
}

func (q *repository) Read(c context.Context) ([]ReportCategoryResponse, error) {
	query := `
	SELECT id, name, created_at, updated_at
	FROM ` + table + ` ORDER BY id`

	rows, err := q.db.QueryContext(c, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []ReportCategoryResponse{}

	for rows.Next() {
		var r ReportCategoryResponse
		var enc_id string

		if err := rows.Scan(
			&enc_id,
			&r.Name,
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

func (q *repository) Update(c context.Context, id string, arg ReportCategoryUpdate) (ReportCategoryResponse, error) {
	var r ReportCategoryResponse

	decryptedID, _ := util.Decrypt(id, "f")

	query := `
	UPDATE ` + table + `
	SET
		name = COALESCE($2, name)
	WHERE id = $1
	RETURNING id, name, created_at, updated_at`

	row := q.db.QueryRowContext(c, query,
		decryptedID,
		arg.Name,
	)

	err := row.Scan(
		&r.HashedID,
		&r.Name,
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
