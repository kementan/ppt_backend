package user

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/redis/go-redis/v9"
)

type (
	UserRepository interface {
		InitCreate(c context.Context) (bool, error)
		CleanUser(c context.Context) (bool, error)
		PassByUEmail(c context.Context, value string) (UserLoginResponse, error)
		GetDataBy(c context.Context, field string, value string) (UserResponse, error)
		GetTable(c context.Context, arg util.DataFilter) ([]UserResponse, error)
		CountRecords(c context.Context, arg util.DataFilter) (int, error)
		IsNullPassword(c context.Context, field string, value string) (UserResponse, error)
		IsComplete(c context.Context, email string) (UserResponse, error)
		IsVerified(c context.Context, email string) (UserResponse, error)
		DoVerify(c context.Context, email string) (UserResponse, error)
		GetConfig(c context.Context) (SMTPData, error)
		GetDefaultRole(c context.Context) (int, error)
		Create(c context.Context, arg UserCreate) (UserResponse, error)
		Read(c context.Context) ([]UserResponse, error)
		Update(c context.Context, field, id string, arg UserUpdate, status string) (UserResponse, error)
		Delete(c context.Context, ids []string) ([]string, []string, error)
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
	var enc_id, enc_role_id, dec_name string
	var pob, dob, dec_nik sql.NullString
	query := `
	SELECT 
	u.id AS u_id,
	u.name, 
	u.nik, 
	u.pob, 
	u.dob, 
	u.role_id, 
	r.name as role_name, 
	u.username, 
	u.password, 
	u.email,
	u.is_verified,
	u.is_complete,
	u.is_active
	FROM ` + table + ` AS u 
	LEFT JOIN ` + subtable + ` AS r ON r.id = u.role_id
	WHERE ` + field + ` = $1
	LIMIT 1`

	row := q.db.QueryRowContext(c, query, value)
	err := row.Scan(
		&enc_id,
		&dec_name,
		&dec_nik,
		&pob,
		&dob,
		&enc_role_id,
		&r.RoleName,
		&r.Username,
		&r.Password,
		&r.Email,
		&r.IsVerified,
		&r.IsComplete,
		&r.IsActive,
	)

	if dec_nik.Valid {
		r.NIK, _ = util.Decrypt(dec_nik.String, "f")
	}

	if pob.Valid {
		r.POB = pob.String
	}

	if dob.Valid {
		r.DOB = dob.String
	}

	r.HashedID, _ = util.Encrypt(enc_id, "f")
	r.Name, _ = util.Decrypt(dec_name, "f")
	r.RoleID, _ = util.Encrypt(enc_role_id, "f")

	return r, err
}

func (q *repository) IsNullPassword(c context.Context, field string, value string) (UserResponse, error) {
	var r UserResponse

	query := `
	SELECT username, email, is_verified FROM ` + table + `
	WHERE ` + field + ` = $1
	AND password IS NULL
	LIMIT 1`

	row := q.db.QueryRowContext(c, query, value)
	err := row.Scan(
		&r.Username,
		&r.Email,
		&r.IsVerified,
	)

	return r, err
}

func (q *repository) IsComplete(c context.Context, email string) (UserResponse, error) {
	var r UserResponse
	var dec_name string

	query := `
	SELECT name, is_complete FROM ` + table + `
	WHERE email = $1
	AND is_complete = $2
	LIMIT 1`

	row := q.db.QueryRowContext(c, query, email, true)
	err := row.Scan(
		&dec_name,
		&r.IsComplete,
	)

	r.Name, _ = util.Decrypt(dec_name, "f")

	return r, err
}

func (q *repository) IsVerified(c context.Context, email string) (UserResponse, error) {
	var r UserResponse
	var dec_name string

	query := `
	SELECT name, is_verified FROM ` + table + `
	WHERE email = $1
	AND is_verified = $2
	LIMIT 1`

	row := q.db.QueryRowContext(c, query, email, true)
	err := row.Scan(
		&dec_name,
		&r.IsVerified,
	)

	r.Name, _ = util.Decrypt(dec_name, "f")

	return r, err
}

func (q *repository) DoVerify(c context.Context, email string) (UserResponse, error) {
	var r UserResponse

	query := `
	UPDATE ` + table + `
	SET
		is_verified = COALESCE($2, is_verified)
	WHERE email = $1
	RETURNING is_verified`

	row := q.db.QueryRowContext(c, query,
		email,
		true,
	)

	err := row.Scan(
		&r.IsVerified,
	)

	return r, err
}

func (q *repository) GetConfig(c context.Context) (SMTPData, error) {
	var cfg ConfigData
	var smtp SMTPData
	query := `
	SELECT name, value
	FROM ppt_configurations ORDER BY id`

	rows, err := q.db.QueryContext(c, query)
	if err != nil {
		return smtp, err
	}

	defer rows.Close()

	items := []ConfigData{}

	for rows.Next() {
		if err := rows.Scan(
			&cfg.Name,
			&cfg.Value,
		); err != nil {
			return smtp, err
		}

		items = append(items, cfg)
	}

	if err := rows.Close(); err != nil {
		return smtp, err
	}

	if err := rows.Err(); err != nil {
		return smtp, err
	}

	for _, val := range items {
		switch val.Name {
		case "smtp_server":
			smtp.SMTPServer = val.Value
		case "smtp_port":
			port, err := strconv.Atoi(val.Value)
			if err != nil {
				return SMTPData{}, err
			}
			smtp.SMTPPort = port
		case "smtp_email":
			smtp.SMTPEmail = val.Value
		case "smtp_email_password":
			smtp.SMTPPassword = val.Value
		}
	}

	return smtp, nil
}

func (q *repository) GetDefaultRole(c context.Context) (int, error) {
	var id int

	query := `
	SELECT id FROM ppt_roles
	WHERE name = 'default'
	OR name = 'Default'
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
	var dec_name, dec_nik sql.NullString

	query := `
	SELECT nik, province_id, regency_id, subdistrict_id, urbanvillage_id, name, username, email, password, google_id, address FROM ` + table + `
	WHERE (username = $1 OR email = $1)
	LIMIT 1`

	row := q.db.QueryRowContext(c, query, value)
	err := row.Scan(
		&dec_nik,
		&r.ProvinceID,
		&r.RegencyID,
		&r.SubdistrictID,
		&r.UrbanvillageID,
		&dec_name,
		&r.Username,
		&r.Email,
		&r.Password,
		&r.GoogleID,
		&r.Address,
	)

	if err != nil {
		return r, err
	}

	if dec_nik.Valid {
		dec_nik.String, _ = util.Decrypt(dec_nik.String, "f")
	}
	r.NIK = dec_nik

	if dec_name.Valid {
		dec_name.String, _ = util.Decrypt(dec_name.String, "f")
	}
	r.Name = dec_name

	return r, err
}

func (q *repository) GetTable(c context.Context, arg util.DataFilter) ([]UserResponse, error) {
	offset := (arg.Page - 1) * arg.PageSize
	items := []UserResponse{}
	args := make([]interface{}, 0)
	args = append(args, arg.PageSize, offset)

	query := `
		SELECT 
		u.id, 
		u.name, 
		u.username, 
		u.email, 
		r.name AS role_name, 
		u.is_active, 
		u.is_complete, 
		u.is_verified, 
		u.created_at, 
		u.updated_at
		FROM ` + table + ` AS u
		LEFT JOIN ` + subtable + ` AS r ON r.id = u.role_id`

	if arg.Search != "" {
		query += ` WHERE lower(u.username) LIKE CONCAT('%%',$3::text,'%%')`
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
		var r UserResponse
		var enc_id, dec_name string

		if err := rows.Scan(
			&enc_id,
			&dec_name,
			&r.Username,
			&r.Email,
			&r.RoleName,
			&r.IsActive,
			&r.IsComplete,
			&r.IsVerified,
			&r.CreatedAt,
			&r.UpdatedAt,
		); err != nil {
			return nil, err
		}

		r.HashedID, _ = util.Encrypt(enc_id, "f")
		r.Name, _ = util.Decrypt(dec_name, "f")

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
			role_id, province_id, regency_id, subdistrict_id, urbanvillage_id, name, username, email, phone, password, nik, nip, is_active, is_complete, is_verified
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
		true,
		true,
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

	smtp_password, _ := util.Encrypt("b51ccde81e880682739e80de", "f")
	config := []map[string]any{
		{"name": "help_desk", "value": "6281133335533", "is_lock": true},
		{"name": "frontend_url", "value": "http://localhost:4200", "is_lock": true},
		{"name": "backend_url", "value": "http://localhost:8080", "is_lock": true},
		{"name": "app_title", "value": "Portal Pertanian Terintegrasi", "is_lock": true},
		{"name": "smtp_server", "value": "smtp.forwardemail.net", "is_lock": true},
		{"name": "smtp_port", "value": "587", "is_lock": true},
		{"name": "smtp_email", "value": "portal.pertanian@ghalyf.com", "is_lock": true},
		{"name": "smtp_email_password", "value": smtp_password, "is_lock": true},
		{"name": "dukcapil_treshold", "value": "100", "is_lock": true},
		{"name": "dukcapil_user_id", "value": "26082022160454BADAN_PENYULUHAN_SDM8370", "is_lock": true},
		{"name": "dukcapil_password", "value": "TH854Y", "is_lock": true},
		{"name": "dukcapil_ip_user", "value": "10.160.84.10", "is_lock": true},
		{"name": "api_token_sipdps_jawa_barat", "value": "FE9C98E3B93FE119DB275F93C761D", "is_lock": false},
		{"name": "api_token_sipdps_jawa_tengah", "value": "7D1D23FF6BCD27AD411A4EC624391", "is_lock": false},
		{"name": "api_token_perbenihan", "value": "PMAT-01H9KCS5K0Q15HZ3YWFTB7WPNE", "is_lock": false},
	}

	for _, v := range config {
		query4 := `
		INSERT INTO ppt_configurations (
			name, value, is_lock
		) VALUES ($1, $2, $3)`

		q.db.Exec(query4, v["name"], v["value"], v["is_lock"])
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

func (q *repository) Create(c context.Context, arg UserCreate) (UserResponse, error) {
	var r UserResponse
	var enc_role_id string

	query := `
	INSERT INTO ` + table + ` (
		role_id, name, username, email, password, google_id, is_active
	) VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING role_id, name, username, email`

	row := q.db.QueryRowContext(c, query, arg.RoleID, arg.Name, arg.Username, arg.Email, arg.Password, arg.GoogleID, arg.IsActive)

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

func (q *repository) Update(c context.Context, field, id string, arg UserUpdate, status string) (UserResponse, error) {
	var r UserResponse
	var row *sql.Row
	var err error

	if status == "bio" {
		query := `
		UPDATE ` + table + `
		SET
			role_id = COALESCE($2, role_id),
			name = COALESCE($3, name),
			pob = COALESCE($4, pob),
			dob = COALESCE($5, dob),
			nik = COALESCE($6, nik),
			gender = COALESCE($7, gender),
			phone = COALESCE($8, phone),
			password = COALESCE($9, password),
			is_active = COALESCE($10, is_active),
			is_complete = COALESCE($11, is_complete),
			is_verified = COALESCE($12, is_verified),
			username = COALESCE($13, username),
			email = COALESCE($14, email)
		WHERE ` + field + ` = $1
		RETURNING role_id, name`

		row = q.db.QueryRowContext(c, query,
			id,
			arg.RoleID,
			arg.Name,
			arg.POB,
			arg.DOB,
			arg.NIK,
			arg.Gender,
			arg.Phone,
			arg.Password,
			arg.IsActive,
			arg.IsComplete,
			arg.IsVerified,
			arg.Username,
			arg.Email,
		)
	} else if status == "address" {
		query := `
		UPDATE ` + table + `
		SET
			province_id = COALESCE($2, province_id),
			regency_id = COALESCE($3, province_id),
			subdistrict_id = COALESCE($4, subdistrict_id),
			urbanvillage_id = COALESCE($5, urbanvillage_id),
			address = COALESCE($6, address),
			latitude = COALESCE($7, latitude),
			longitude = COALESCE($8, longitude),
			is_complete = COALESCE($9, is_complete)
		WHERE ` + field + ` = $1
		RETURNING role_id, name`

		row = q.db.QueryRowContext(c, query,
			id,
			arg.ProvinceID,
			arg.RegencyID,
			arg.SubdistrictID,
			arg.UrbanvillageID,
			arg.Address,
			arg.Latitude,
			arg.Longitude,
			true,
		)
	}

	err = row.Scan(
		&r.RoleID,
		&r.Name,
	)

	return r, err
}

func (q *repository) UpdateSubSector(c context.Context, user_id int, subsectorIDs []string) error {
	query1 := `DELETE FROM ppt_sub_sector_accesses WHERE user_id = $1`
	_, err := q.db.ExecContext(c, query1, user_id)
	if err != nil {
		return err
	}

	query2 := "INSERT INTO ppt_sub_sector_accesses (user_id, sub_sector_id) VALUES "
	values := []interface{}{}

	placeholders := make([]string, len(subsectorIDs))
	for i := range subsectorIDs {
		placeholders[i] = fmt.Sprintf("($%d, $%d)", 2*i+1, 2*i+2)
		values = append(values, user_id, subsectorIDs[i])
	}

	query2 += strings.Join(placeholders, ",")
	_, err = q.db.Exec(query2, values...)
	return err
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
