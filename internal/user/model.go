package user

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

var table = "ppt_users"

type (
	UserLoginRequest struct {
		NIK      string `json:"nik" binding:"required,min=16,max=16"`
		Password string `json:"password" binding:"required,min=8"`
	}

	UserLoginReponse struct {
		SessionID             uuid.UUID    `json:"session_id"`
		AccessToken           string       `json:"access_token"`
		AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
		RefreshToken          string       `json:"refresh_token"`
		RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
		User                  UserResponse `json:"user"`
	}

	UserCreateRequest struct {
		RoleID   string `json:"role_id" binding:"required,min=1"`
		Name     string `json:"name" binding:"required,min=1"`
		NIK      string `json:"nik" binding:"required,min=1"`
		NIP      string `json:"nip" binding:"required,min=1"`
		Email    string `json:"email" binding:"required,min=1"`
		Password string `json:"password" binding:"required,min=1"`
		Phone    string `json:"phone" binding:"required,min=1"`
	}

	UserUpdateRequest struct {
		ID       string `json:"id" binding:"required,min=1"`
		RoleID   string `json:"role_id" binding:"required,min=1"`
		Name     string `json:"name" binding:"required,min=1"`
		NIK      string `json:"nik" binding:"required,min=1"`
		NIP      string `json:"nip" binding:"required,min=1"`
		Email    string `json:"email" binding:"required,min=1"`
		Password string `json:"password" binding:"required,min=1"`
		Phone    string `json:"phone" binding:"required,min=1"`
	}

	UserDeleteRequest struct {
		ID string `json:"id" binding:"required,min=1"`
	}

	UserCreate struct {
		RoleID    string    `json:"role_id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		NIK       string    `json:"nik"`
		NIP       string    `json:"nip"`
		Phone     string    `json:"phone"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	SearchResponse struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
		} `json:"hits"`
	}

	UserIndex struct {
		Email string `json:"email"`
		NIK   string `json:"nik"`
		NIP   string `json:"nip"`
		Phone string `json:"phone"`
	}

	UserDummy struct {
		Email string `json:"email"`
		NIK   string `json:"nik"`
		NIP   string `json:"nip"`
		Phone string `json:"phone"`
	}

	UserUpdate struct {
		RoleID    sql.NullString `json:"role_id"`
		Name      sql.NullString `json:"name"`
		Email     sql.NullString `json:"email"`
		Password  sql.NullString `json:"password"`
		NIK       sql.NullString `json:"nik"`
		NIP       sql.NullString `json:"nip"`
		CreatedAt sql.NullString `json:"created_at"`
		UpdatedAt sql.NullString `json:"updated_at"`
	}

	UserResponse struct {
		HashedID  string    `json:"id"`
		RoleID    string    `json:"role_id"`
		Name      string    `json:"name"`
		NIK       string    `json:"nik"`
		NIP       string    `json:"nip"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Phone     string    `json:"phone"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
