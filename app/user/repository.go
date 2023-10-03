package user

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/redis/go-redis/v9"
)

type (
	UserRepository interface {
		InitCreate(c context.Context) (bool, error)
		CleanUser(c context.Context) (bool, error)
		CheckData(c context.Context) (bool, error)
		PassByUEmail(c context.Context, value string) (UserLoginResponse, error)
		GetDataBy(c context.Context, field string, value string) (UserResponse, error)
		GetDefaultRole(c context.Context) (int, error)
		Create(c context.Context, arg UserCreate) (UserResponse, error)
		Read(c context.Context) ([]UserResponse, error)
		Update(c context.Context, id string, arg UserUpdate) (UserResponse, error)
		Delete(c context.Context, id string) error
	}

	repository struct {
		db  *sql.DB
		rdb *redis.Client
		edb *elasticsearch.Client
	}
)

func NewRepository(db *sql.DB, rdb *redis.Client, edb *elasticsearch.Client) UserRepository {
	return &repository{
		db:  db,
		rdb: rdb,
		edb: edb,
	}
}

func (q *repository) GetDataBy(c context.Context, field string, value string) (UserResponse, error) {
	var r UserResponse
	var enc_id, enc_role_id string

	query := `
	SELECT * FROM ` + table + `
	WHERE ` + field + ` = $1
	LIMIT 1`

	row := q.db.QueryRowContext(c, query, value)
	err := row.Scan(
		&enc_id,
		&enc_role_id,
		&r.Name,
		&r.NIK,
		&r.NIP,
		&r.Email,
		&r.Phone,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	r.HashedID, _ = util.Encrypt(enc_id, "f")
	r.RoleID, _ = util.Encrypt(enc_role_id, "f")

	return r, err
}

func (q *repository) GetDefaultRole(c context.Context) (int, error) {
	var id int

	query := `
	SELECT id FROM ppt_roles
	WHERE name = 'default'
	LIMIT 1`

	row := q.db.QueryRowContext(c, query)
	err := row.Scan(
		&id,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (q *repository) PassByUEmail(c context.Context, value string) (UserLoginResponse, error) {
	var r UserLoginResponse

	query := `
	SELECT username, email, password FROM ` + table + `
	WHERE (username = $1 OR email = $1)
	LIMIT 1`

	row := q.db.QueryRowContext(c, query, value)
	err := row.Scan(
		&r.Username,
		&r.Email,
		&r.Password,
	)

	return r, err
}

func (q *repository) InitCreate(c context.Context) (bool, error) {
	queries := []string{
		"DELETE FROM ppt_configurations",
		"DELETE FROM ppt_service_accesses",
		"DELETE FROM ppt_sub_menu_accesses",
		"DELETE FROM ppt_users",
		"DELETE FROM ppt_roles",
		"DELETE FROM ppt_sub_menus",
		"DELETE FROM ppt_menus",
	}

	for _, val := range queries {
		_, err := q.db.Exec(val)
		if err != nil {
			return false, fmt.Errorf("error executing pq delete request: %v", err)
		}
	}

	name, _ := util.Encrypt("Super Admin", "f")
	username := "superadmin"
	email := "super@admin.com"
	phone, _ := util.Encrypt("081133335533", "f")
	password, _ := util.HashPassword("password")
	nik, _ := util.Encrypt("3316110802720002", "f")
	nip, _ := util.Encrypt("1234123412341234", "f")

	query1 := `
		WITH new_role AS (
			INSERT INTO ppt_roles (name, is_create, is_read, is_update, is_delete, is_public)
			VALUES ('superadmin', true, true, true, true, false)
			RETURNING id
		)
		INSERT INTO ` + table + ` (
			role_id, province_id, regency_id, subdistrict_id, urbanvillage_id, name, username, email, phone, password, nik, nip, is_active
		) VALUES ((SELECT id FROM new_role), 
		'11',
		'11.01',
		'11.01.01',
		'11.01.01.2001',
		'` + name + `', 
		'` + username + `', 
		'` + email + `', 
		'` + phone + `', 
		'` + password + `', 
		'` + nik + `', 
		'` + nip + `', 
		true)
		RETURNING id, role_id`

	query2 := `
			INSERT INTO ppt_roles (name, is_create, is_read, is_update, is_delete, is_public)
			VALUES ('default', false, true, false, false, true)`
	q.db.Exec(query2)

	var user_id, role_id int

	err := q.db.QueryRowContext(c, query1).Scan(&user_id, &role_id)
	if err != nil {
		return false, err
	}

	menu := []map[string]string{
		{"name": "Dashboard", "slug": "dashboard", "sort": "1"},
		{"name": "User", "slug": "", "sort": "2"},
		{"name": "Menu", "slug": "", "sort": "3"},
		{"name": "Layanan", "slug": "", "sort": "4"},
		{"name": "Laporan", "slug": "", "sort": "5"},
	}

	sub_menu := []map[string]string{
		{"m": "User", "name": "Data Role", "slug": "admin/user", "sort": "1"},
		{"m": "User", "name": "Data User", "slug": "admin/role", "sort": "2"},
		{"m": "Menu", "name": "Data Menu", "slug": "admin/menu", "sort": "1"},
		{"m": "Menu", "name": "Akses Menu", "slug": "admin/akses-menu", "sort": "2"},
		{"m": "Layanan", "name": "Data Layanan", "slug": "admin/layanan", "sort": "1"},
		{"m": "Layanan", "name": "Akses Layanan", "slug": "admin/akses-layanan", "sort": "2"},
		{"m": "Laporan", "name": "Data Laporan", "slug": "admin/laporan", "sort": "1"},
	}

	roleid := strconv.Itoa(role_id)
	for _, v := range menu {
		var menu_id int
		query3 := `
		INSERT INTO ppt_menus (name, slug, is_active, sort)
		VALUES ($1, $2, $3, $4)
		RETURNING id
		`
		err = q.db.QueryRowContext(c, query3, v["name"], v["slug"], true, v["sort"]).Scan(&menu_id)
		if err != nil {
			return false, err
		}

		for _, sv := range sub_menu {
			if v["slug"] == "" && v["name"] == sv["m"] {
				menuid := strconv.Itoa(menu_id)
				query4 := `
				WITH new_sub_menu AS (
					INSERT INTO ppt_sub_menus (menu_id, name, slug, is_active, sort
					) VALUES ($1, $2, $3, $4, $5)
					RETURNING id
				)
				INSERT INTO ppt_sub_menu_accesses (
					role_id, sub_menu_id
				) VALUES ($6, (SELECT id FROM new_sub_menu))
				`

				q.db.Exec(
					query4,
					menuid,
					sv["name"],
					sv["slug"],
					true,
					sv["sort"],
					roleid,
				)
				if err != nil {
					return false, err
				}
			}
		}
	}

	smtp_password, _ := util.Encrypt("456Info123", "f")
	config := []map[string]string{
		{"name": "app_url", "value": "http://localhost"},
		{"name": "app_title", "value": "Portal Pertanian Terintegrasi"},
		{"name": "smtp_server", "value": "mail.ghalyf.com"},
		{"name": "smtp_port", "value": "465"},
		{"name": "smtp_email", "value": "info@ghalyf.com"},
		{"name": "smtp_email_password", "value": smtp_password},
	}

	for _, v := range config {
		query4 := `
		INSERT INTO ppt_configurations (
			name, value
		) VALUES ('` + v["name"] + `', '` + v["value"] + `')`

		q.db.Exec(query4)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (q *repository) CleanUser(c context.Context) (bool, error) {
	query := `
	DELETE FROM ppt_users
	WHERE province_id IS NULL
	AND regency_id IS NULL
	AND subdistrict_id IS NULL
	AND urbanvillage_id IS NULL
	AND address IS NULL
	AND is_active = false`

	_, err := q.db.Exec(query)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (q *repository) CheckData(c context.Context) (bool, error) {
	return true, nil
}

func (q *repository) Create(c context.Context, arg UserCreate) (UserResponse, error) {
	var r UserResponse
	var enc_role_id string

	query := `
	INSERT INTO ` + table + ` (
		role_id, name, username, email, password, is_active
	) VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING role_id, name, username, email`

	row := q.db.QueryRowContext(c, query, arg.RoleID, arg.Name, arg.Username, arg.Email, arg.Password, arg.IsActive)

	err := row.Scan(
		&enc_role_id,
		&r.Name,
		&r.Username,
		&r.Email,
	)

	r.RoleID, _ = util.Encrypt(enc_role_id, "f")

	return r, err
}

func (q *repository) Read(c context.Context) ([]UserResponse, error) {
	query := `
	SELECT id, name, created_at, updated_at
	FROM ` + table + ` ORDER BY id`

	rows, err := q.db.QueryContext(c, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []UserResponse{}

	for rows.Next() {
		var r UserResponse
		var enc_id string

		if err := rows.Scan(
			&enc_id,
			&r.Name,
		); err != nil {
			return nil, err
		}

		encryptedID, _ := util.Encrypt(enc_id, "f")
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

func (q *repository) Update(c context.Context, id string, arg UserUpdate) (UserResponse, error) {
	var r UserResponse

	decryptedID, _ := util.Decrypt(id, "f")

	query := `
	UPDATE ` + table + `
	SET
		name = COALESCE($2, name)
	WHERE id = $1
	RETURNING *`

	row := q.db.QueryRowContext(c, query,
		decryptedID,
		arg.Name,
	)

	err := row.Scan(
		&r.RoleID,
		&r.Name,
		&r.NIK,
		&r.NIP,
		&r.Email,
		&r.Phone,
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
