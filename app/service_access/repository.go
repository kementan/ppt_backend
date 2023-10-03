package service_access

import (
	"context"
	"database/sql"

	"github.com/gigaflex-co/ppt_backend/util"
)

type (
	ServiceAccessRepository interface {
		Create(c context.Context, arg ServiceAccessCreate) (ServiceAccessResponse, error)
		Read(c context.Context) ([]ServiceAccessResponse, error)
		Update(c context.Context, id string, arg ServiceAccessUpdate) (ServiceAccessResponse, error)
		Delete(c context.Context, id string) error
	}

	repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) ServiceAccessRepository {
	return &repository{
		db: db,
	}
}

func (q *repository) Create(c context.Context, arg ServiceAccessCreate) (ServiceAccessResponse, error) {
	var r ServiceAccessResponse
	var enc_id, enc_service_id, enc_role_id string

	query := `
	INSERT INTO ` + table + ` (
		role_id, service_id
	) VALUES ($1, $2)
	RETURNING *`

	row := q.db.QueryRowContext(c, query, arg.RoleID, arg.ServiceID)

	err := row.Scan(
		&enc_id,
		&enc_service_id,
		&enc_role_id,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	r.HashedID, _ = util.Encrypt(enc_id, "f")
	r.HashedServiceID, _ = util.Encrypt(enc_service_id, "f")
	r.HashedRoleID, _ = util.Encrypt(enc_role_id, "f")

	return r, err
}

func (q *repository) Read(c context.Context) ([]ServiceAccessResponse, error) {
	query := `
	SELECT *
	FROM ` + table + ` ORDER BY id`

	rows, err := q.db.QueryContext(c, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []ServiceAccessResponse{}

	for rows.Next() {
		var r ServiceAccessResponse
		var enc_id, enc_role_id, enc_service_id string

		if err := rows.Scan(
			&enc_id,
			&enc_role_id,
			&enc_service_id,
			&r.CreatedAt,
			&r.UpdatedAt,
		); err != nil {
			return nil, err
		}

		r.HashedID, _ = util.Encrypt(enc_id, "f")
		r.HashedRoleID, _ = util.Encrypt(enc_role_id, "f")
		r.HashedServiceID, _ = util.Encrypt(enc_service_id, "f")

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

func (q *repository) Update(c context.Context, id string, arg ServiceAccessUpdate) (ServiceAccessResponse, error) {
	var r ServiceAccessResponse
	var enc_id, enc_role_id, enc_service_id string

	decryptedID, _ := util.Decrypt(id, "f")

	query := `
	UPDATE ` + table + `
	SET
		role_id = COALESCE($2, role_id)
		service_id = COALESCE($3, service_id)
	WHERE id = $1
	RETURNING *`

	row := q.db.QueryRowContext(c, query,
		decryptedID,
		arg.RoleID,
		arg.ServiceID,
	)

	err := row.Scan(
		&enc_id,
		&enc_role_id,
		&enc_service_id,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	r.HashedID, _ = util.Encrypt(enc_id, "f")
	r.HashedRoleID, _ = util.Encrypt(enc_role_id, "f")
	r.HashedServiceID, _ = util.Encrypt(enc_service_id, "f")

	return r, err
}

func (q *repository) Delete(c context.Context, id string) error {
	decryptedID, _ := util.Decrypt(id, "f")

	query := `
	DELETE FROM ` + table + ` WHERE id = $1`

	_, err := q.db.ExecContext(c, query, decryptedID)
	return err
}
