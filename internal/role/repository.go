package role

import (
	"context"
	"database/sql"

	"gitlab.com/xsysproject/ppt_backend/helper"
)

type (
	RoleRepository interface {
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

func (q *repository) Create(ctx context.Context, arg RoleCreate) (RoleResponse, error) {
	var i RoleResponse

	query := `
	INSERT INTO ` + table + ` (
		name, created_at
	) VALUES ($1, $2) 
	RETURNING name, created_at`

	row := q.db.QueryRowContext(ctx, query, arg.Name, arg.CreatedAt)

	err := row.Scan(
		&i.Name,
		&i.CreatedAt,
	)

	return i, err
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
		r.ID = encryptedID

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
		&r.ID,
		&r.Name,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	r.ID = id

	return r, err
}

func (q *repository) Delete(ctx context.Context, id string) error {
	decryptedID, _ := helper.Decrypt(id)

	query := `
	DELETE FROM ` + table + ` WHERE id = $1`
	_, err := q.db.ExecContext(ctx, query, decryptedID)
	return err
}
