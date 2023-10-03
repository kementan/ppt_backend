package user

import (
	"database/sql"
	"time"
)

var table = "ppt_users"

type (
	UserLoginRequest struct {
		Username string `json:"username" binding:"required,min=8,max=150"`
		Password string `json:"password" binding:"required,min=8,max=150"`
	}

	UserLoginResponse struct {
		Username string
		Email    string
		Password string
	}

	UserRegisterRequest struct {
		Name     string `json:"name" binding:"required,min=3,max=100"`
		Username string `json:"username" binding:"required,min=8,max=150"`
		Email    string `json:"email" binding:"required,min=8,max=150"`
		Password string `json:"password" binding:"required,min=8,max=150"`
	}

	UserRecaptchaRequest struct {
		Response string `json:"gresponse"`
	}

	UserCreateRequest struct {
		RoleID         string `json:"role_id" binding:"required,min=1,max=200"`
		ProvinceID     string `json:"province_id" binding:"required,min=1,max=200"`
		RegencyID      string `json:"regency_id" binding:"required,min=1,max=200"`
		SubdistrictID  string `json:"subdistrict_id" binding:"required,min=1,max=200"`
		UrbanvillageID string `json:"urbanvillage_id" binding:"required,min=1,max=200"`
		Name           string `json:"name" binding:"required,min=3,max=100"`
		Username       string `json:"username" binding:"required,min=8,max=150"`
		Email          string `json:"email" binding:"required,min=8,max=150"`
		Phone          string `json:"phone" binding:"required,min=9,max=15"`
		Password       string `json:"password" binding:"required,min=8,max=150"`
		NIK            string `json:"nik" binding:"required,min=16,max=16"`
		NIP            string `json:"nip" binding:"required,min=8,max=150"`
		DOB            string `json:"dob" binding:"required,min=8,max=150"`
		POB            string `json:"pob" binding:"required,min=8,max=150"`
		ImgUser        string `json:"img_user" binding:"required,min=1,max=200"`
		ImgID          string `json:"img_id"`
		Address        string `json:"address" binding:"required,min=8,max=150"`
	}

	UserUpdateRequest struct {
		ID             string `json:"id" binding:"required,min=1"`
		RoleID         string `json:"role_id" binding:"required,min=1,max=200"`
		ProvinceID     string `json:"province_id" binding:"required,min=1,max=200"`
		RegencyID      string `json:"regency_id" binding:"required,min=1,max=200"`
		SubdistrictID  string `json:"subdistrict_id" binding:"required,min=1,max=200"`
		UrbanvillageID string `json:"urbanvillage_id" binding:"required,min=1,max=200"`
		Name           string `json:"name" binding:"required,min=3,max=100"`
		Username       string `json:"username" binding:"required,min=8,max=150"`
		Email          string `json:"email" binding:"required,min=8,max=150"`
		Phone          string `json:"phone" binding:"required,min=9,max=15"`
		Password       string `json:"password" binding:"required,min=8,max=150"`
		NIK            string `json:"nik" binding:"required,min=16,max=16"`
		NIP            string `json:"nip" binding:"required,min=8,max=150"`
		DOB            string `json:"dob" binding:"required,min=8,max=150"`
		POB            string `json:"pob" binding:"required,min=8,max=150"`
		ImgUser        string `json:"img_user" binding:"required,min=1,max=200"`
		ImgID          string `json:"img_id"`
		Address        string `json:"address" binding:"required,min=8,max=150"`
	}

	UserDeleteRequest struct {
		ID string `json:"id" binding:"required,min=1"`
	}

	UserCreate struct {
		RoleID         int       `json:"role_id"`
		ProvinceID     string    `json:"province_id"`
		RegencyID      string    `json:"regency_id"`
		SubdistrictID  string    `json:"subdistrict_id"`
		UrbanvillageID string    `json:"urbanvillage_id"`
		Name           string    `json:"name"`
		Username       string    `json:"username"`
		Email          string    `json:"email"`
		Phone          string    `json:"phone"`
		Password       string    `json:"password"`
		NIK            string    `json:"nik"`
		NIP            string    `json:"nip"`
		DOB            string    `json:"dob"`
		POB            string    `json:"pob"`
		ImgUser        string    `json:"img_user"`
		ImgID          string    `json:"img_id"`
		Address        string    `json:"address"`
		IsActive       bool      `json:"is_active"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	UserUpdate struct {
		RoleID         sql.NullString `json:"role_id"`
		ProvinceID     sql.NullString `json:"province_id"`
		RegencyID      sql.NullString `json:"regency_id"`
		SubdistrictID  sql.NullString `json:"subdistrict_id"`
		UrbanvillageID sql.NullString `json:"urbanvillage_id"`
		Name           sql.NullString `json:"name"`
		Username       sql.NullString `json:"username"`
		Email          sql.NullString `json:"email"`
		Phone          sql.NullString `json:"phone"`
		Password       sql.NullString `json:"password"`
		NIK            sql.NullString `json:"nik"`
		NIP            sql.NullString `json:"nip"`
		DOB            sql.NullString `json:"dob"`
		POB            sql.NullString `json:"pob"`
		ImgUser        sql.NullString `json:"img_user"`
		ImgID          sql.NullString `json:"img_id"`
		Address        sql.NullString `json:"address"`
		CreatedAt      sql.NullString `json:"created_at"`
		UpdatedAt      sql.NullString `json:"updated_at"`
	}

	UserResponse struct {
		HashedID       string    `json:"id"`
		RoleID         string    `json:"role_id"`
		ProvinceID     string    `json:"province_id"`
		RegencyID      string    `json:"regency_id"`
		SubdistrictID  string    `json:"subdistrict_id"`
		UrbanvillageID string    `json:"urbanvillage_id"`
		Name           string    `json:"name"`
		Username       string    `json:"username"`
		Email          string    `json:"email"`
		Phone          string    `json:"phone"`
		NIK            string    `json:"nik"`
		NIP            string    `json:"nip"`
		DOB            string    `json:"dob"`
		POB            string    `json:"pob"`
		ImgUser        string    `json:"img_user"`
		ImgID          string    `json:"img_id"`
		Address        string    `json:"address"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	SearchResponse struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
		} `json:"hits"`
	}

	UserIndex struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
)
