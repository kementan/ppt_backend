package user

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/redis/go-redis/v9"
	"gitlab.com/xsysproject/ppt_backend/helper"
)

type (
	UserRepository interface {
		GetDataBy(ctx context.Context, field string, value string) (UserResponse, error)
		CheckESData(ctx context.Context, field string, value string) (bool, error)
		Create(ctx context.Context, arg UserCreate) (UserResponse, error)
		Dummy(ctx context.Context) error
		Read(ctx context.Context) ([]UserResponse, error)
		Update(ctx context.Context, id string, arg UserUpdate) (UserResponse, error)
		Delete(ctx context.Context, id string) error
		DeleteESRecord(index, doc_id string) error

		// Login(ctx context.Context, id string, arg UserUpdate) (UserResponse, error)
		// CheckData(ctx context.Context, key string, value string) (bool, error)
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

func (q *repository) GetDataBy(ctx context.Context, field string, value string) (UserResponse, error) {
	var r UserResponse
	var enc_id string

	query := `
	SELECT * FROM ` + table + `
	WHERE ` + field + ` = $1
	LIMIT 1`

	row := q.db.QueryRowContext(ctx, query, value)
	err := row.Scan(
		&enc_id,
		&r.Name,
		&r.NIK,
		&r.NIP,
		&r.Email,
		&r.Password,
		&r.Phone,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	encID, _ := helper.Encrypt(enc_id)
	r.HashedID = encID

	return r, err
}

func (q *repository) CheckESData(ctx context.Context, field string, value string) (bool, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"size":    500000,
		"_source": []string{field},
	}

	queryJSON, err := json.Marshal(query)
	if err != nil {
		return true, fmt.Errorf("error converting query to JSON: %v", err)
	}

	res, err := q.edb.Search(
		q.edb.Search.WithContext(context.Background()),
		q.edb.Search.WithIndex(table),
		q.edb.Search.WithBody(bytes.NewReader(queryJSON)),
	)
	if err != nil {
		return true, fmt.Errorf("error executing search request: %v", err)
	}
	defer res.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return true, fmt.Errorf("error parsing search response: %v", err)
	}

	hits, ok := data["hits"].(map[string]interface{})
	if !ok {
		return true, fmt.Errorf("invalid hits format in search response 2")
	}

	hitsArray, ok := hits["hits"].([]interface{})
	if !ok {
		return true, fmt.Errorf("invalid hits format in search response 3")
	}

	for _, hit := range hitsArray {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			continue
		}

		source, ok := hitMap["_source"]
		if !ok {
			continue
		}

		sourceMap, ok := source.(map[string]interface{})
		if !ok {
			continue
		}

		sourceValue, ok := sourceMap[field].(string)
		if !ok {
			continue
		}

		if sourceValue == "" {
			err = q.DeleteESRecord(table, hitMap["_id"].(string))
			if err != nil {
				return true, fmt.Errorf("%s", err)
			}
		}

		decRes, _ := helper.Decrypt(sourceValue)

		if decRes == value {
			return true, fmt.Errorf("data %s with value %s already exists", field, value)
		}
	}

	return false, nil
}

func (q *repository) Dummy(ctx context.Context) error {
	numRecords := 1000000
	start := time.Now()
	for i := 1; i <= numRecords; i++ {
		docID := fmt.Sprintf("document-%d", i)

		roleID := "value1-" + strconv.Itoa(i)
		name := "value2-" + strconv.Itoa(i)
		email := "value3-" + strconv.Itoa(i)
		nip := "value4-" + strconv.Itoa(i)
		nik := "value5-" + strconv.Itoa(i)
		phone := "value6-" + strconv.Itoa(i)

		user := UserDummy{
			RoleID: roleID,
			Name:   name,
			Email:  email,
			NIP:    nip,
			NIK:    nik,
			Phone:  phone,
		}

		docBytes, err := json.Marshal(user)
		if err != nil {
			continue
		}

		req := esapi.IndexRequest{
			Index:      table,
			DocumentID: docID,
			Body:       bytes.NewReader(docBytes),
		}

		res, err := req.Do(context.Background(), q.edb)
		if err != nil {
			continue
		}
		defer res.Body.Close()

		if res.IsError() {
			continue
		}

		log.Printf("Indexed document %s", docID)

		time.Sleep(1 * time.Millisecond)
	}

	// Calculate the elapsed time
	elapsed := time.Since(start)
	log.Printf("Indexed %d documents in %s", numRecords, elapsed)

	return nil
}

func (q *repository) Create(ctx context.Context, arg UserCreate) (UserResponse, error) {
	var r UserResponse
	var enc_id, enc_role_id string

	query := `
	INSERT INTO ` + table + ` (
		role_id, name, nik, nip, email, password, phone
	) VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING *`

	row := q.db.QueryRowContext(ctx, query, arg.RoleID, arg.Name, arg.NIK, arg.NIP, arg.Email, arg.Password, arg.Phone)

	err := row.Scan(
		&enc_id,
		&enc_role_id,
		&r.Name,
		&r.NIK,
		&r.NIP,
		&r.Email,
		&r.Password,
		&r.Phone,
		&r.CreatedAt,
		&r.UpdatedAt,
	)

	r.HashedID, _ = helper.Encrypt(enc_id)
	r.RoleID, _ = helper.Encrypt(enc_role_id)

	if err == nil {
		_, err2 := q.edb.Index(
			"ppt_users",
			strings.NewReader(fmt.Sprintf(`{"role_id": "%s", "name": "%s", "email": "%s", "nik": "%s", "nip": "%s", "phone": "%s"}`, enc_role_id, r.Name, r.Email, r.NIK, r.NIP, r.Phone)),
			q.edb.Index.WithDocumentID(enc_id),
		)

		if err2 != nil {
			err = helper.MergeErrors(err, err2)
		}
	}

	return r, err
}

func (q *repository) Read(ctx context.Context) ([]UserResponse, error) {
	query := `
	SELECT id, name, created_at, updated_at
	FROM ` + table + ` ORDER BY id`

	rows, err := q.db.QueryContext(ctx, query)
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

func (q *repository) Update(ctx context.Context, id string, arg UserUpdate) (UserResponse, error) {
	var r UserResponse

	decryptedID, _ := helper.Decrypt(id)

	query := `
	UPDATE ` + table + `
	SET
		name = COALESCE($2, name)
	WHERE id = $1
	RETURNING *`

	row := q.db.QueryRowContext(ctx, query,
		decryptedID,
		arg.Name,
	)

	err := row.Scan(
		&r.RoleID,
		&r.Name,
		&r.NIK,
		&r.NIP,
		&r.Email,
		&r.Password,
		&r.Phone,
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

	if err == nil {
		err = q.DeleteESRecord(table, decryptedID)
	}

	return err
}

func (q *repository) DeleteESRecord(index, doc_id string) error {
	req := esapi.DeleteRequest{
		Index:      index,
		DocumentID: doc_id,
	}

	res, err := req.Do(context.Background(), q.edb)
	if err != nil {
		return fmt.Errorf("something's wrong in elastic %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("something's wrong in elastic %s", res.String())
	}

	return nil
}
