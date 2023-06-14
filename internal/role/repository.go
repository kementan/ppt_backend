package role

import (
	"context"
	"database/sql"
)

type (
	RoleRepository interface {
		Create(ctx context.Context, arg RoleCreate) (RoleResponse, error)
		Read(ctx context.Context) ([]RoleResponse, error)
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
	query := `
	INSERT INTO ppt_roles (
		name, created_at
	) VALUES ($1, $2) 
	RETURNING name, created_at`

	row := q.db.QueryRowContext(ctx, query, arg.Name, arg.CreatedAt)

	var i RoleResponse

	err := row.Scan(
		&i.Name,
		&i.CreatedAt,
	)

	return i, err
}

func (q *repository) Read(ctx context.Context) ([]RoleResponse, error) {
	query := `
	SELECT id, name, created_at 
	FROM ppt_roles ORDER BY id`

	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []RoleResponse{}

	for rows.Next() {
		var r RoleResponse
		if err := rows.Scan(
			&r.Name,
			&r.CreatedAt,
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
