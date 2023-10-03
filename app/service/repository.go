package service

import (
	"context"
	"database/sql"

	"github.com/gigaflex-co/ppt_backend/util"
)

type (
	ServiceRepository interface {
		GetDataBy(c context.Context, field string, value string) (ServiceResponse, error)

		Create(c context.Context, arg ServiceCreate) (ServiceResponse, error)
		Read(c context.Context) ([]ServiceResponse, error)
		Update(c context.Context, id string, arg ServiceUpdate) (ServiceResponse, error)
		Delete(c context.Context, id string) error
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
		&r.Image,
		&r.IsActive,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	r.HashedID, _ = util.Encrypt(enc_id, "f")

	return r, err
}

func (q *repository) Create(c context.Context, arg ServiceCreate) (ServiceResponse, error) {
	var r ServiceResponse
	var enc_id string

	query := `
	INSERT INTO ` + table + ` (
		name, image, is_active
	) VALUES ($1, $2, $3)
	RETURNING *`

	row := q.db.QueryRowContext(c, query, arg.Name, arg.Image, arg.IsActive)

	err := row.Scan(
		&enc_id,
		&r.Name,
		&r.Image,
		&r.IsActive,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	r.HashedID, _ = util.Encrypt(enc_id, "f")

	return r, err
}

func (q *repository) Read(c context.Context) ([]ServiceResponse, error) {
	query := `
	SELECT *
	FROM ` + table + ` ORDER BY id`

	rows, err := q.db.QueryContext(c, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []ServiceResponse{}

	for rows.Next() {
		var r ServiceResponse
		var enc_id string

		if err := rows.Scan(
			&enc_id,
			&r.Name,
			&r.Image,
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

func (q *repository) Update(c context.Context, id string, arg ServiceUpdate) (ServiceResponse, error) {
	var r ServiceResponse

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
		&r.Image,
		&r.IsActive,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	r.HashedID = id

	return r, err
}

func (q *repository) Delete(c context.Context, id string) error {
	decryptedID, _ := util.Decrypt(id, "f")

	query := `
	DELETE FROM ` + table + ` WHERE id = $1`

	_, err := q.db.ExecContext(c, query, decryptedID)
	return err
}
