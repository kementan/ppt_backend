package role

import (
	"context"
	"database/sql"

	"gitlab.com/xsysproject/ppt_backend/helper"
)

type (
	RoleRepository interface {
		GetDataBy(ctx context.Context, field string, value string) (RoleResponse, error)
		Create(ctx context.Context, arg RoleCreate) (RoleResponse, error)
		Read(ctx context.Context) ([]RoleResponse, error)
		Update(ctx context.Context, id string, arg RoleUpdate) (RoleResponse, error)
		Delete(ctx context.Context, id string) error
	}

	repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) RoleRepository {
	return &repository{
		db: db,
	}
}

func (q *repository) GetDataBy(ctx context.Context, field string, value string) (RoleResponse, error) {
	var r RoleResponse
	var enc_id string

	query := `
	SELECT * FROM ` + table + `
	WHERE ` + field + ` = $1
	LIMIT 1`

	row := q.db.QueryRowContext(ctx, query, value)
	err := row.Scan(
		&enc_id,
		&r.Name,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	encID, _ := helper.Encrypt(enc_id)
	r.HashedID = encID

	return r, err
}

func (q *repository) Create(ctx context.Context, arg RoleCreate) (RoleResponse, error) {
	var r RoleResponse
	var enc_id string

	query := `
	INSERT INTO ` + table + ` (
		name
	) VALUES ($1)
	RETURNING *`

	row := q.db.QueryRowContext(ctx, query, arg.Name)

	err := row.Scan(
		&enc_id,
		&r.Name,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	r.HashedID, _ = helper.Encrypt(enc_id)

	return r, err
}

func (q *repository) Read(ctx context.Context) ([]RoleResponse, error) {
	query := `
	SELECT id, name, created_at, updated_at
	FROM ` + table + ` ORDER BY id`

	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []RoleResponse{}

	for rows.Next() {
		var r RoleResponse
		var enc_id string

		if err := rows.Scan(
			&enc_id,
			&r.Name,
			&r.CreatedAt,
			&r.UpdatedAt,
		); err != nil {
			return nil, err
		}

		encryptedID, _ := helper.Encrypt(enc_id)
		r.HashedID = encryptedID

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

func (q *repository) Update(ctx context.Context, id string, arg RoleUpdate) (RoleResponse, error) {
	var r RoleResponse

	decryptedID, _ := helper.Decrypt(id)

	query := `
	UPDATE ` + table + `
	SET
		name = COALESCE($2, name)
	WHERE id = $1
	RETURNING id, name, created_at, updated_at`

	row := q.db.QueryRowContext(ctx, query,
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

func (q *repository) Delete(ctx context.Context, id string) error {
	decryptedID, _ := helper.Decrypt(id)

	query := `
	DELETE FROM ` + table + ` WHERE id = $1`

	_, err := q.db.ExecContext(ctx, query, decryptedID)
	return err
}
