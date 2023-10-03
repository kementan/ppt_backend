package util

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gigaflex-co/ppt_backend/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var (
	appConfig, _  = config.LoadConfig("./.")
	revokedTokens = make(map[string]struct{})
)

func CreateToken(c *gin.Context, rdb *redis.Client, email string) (string, error) {
	encEmail, _ := Encrypt(email, "s")
	_, err := rdb.Get(context.Background(), encEmail).Result()
	if err != redis.Nil {
		return "", errors.New("isloggedin")
	}

	expirationTime := time.Now().Add(appConfig.TokenAccessDuration)

	claims := &Claims{
		Email: encEmail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(appConfig.SecretKey))
	if err != nil {
		return "", err
	}

	fullToken := tokenString
	_, err = rdb.Set(context.Background(), encEmail, fullToken, appConfig.TokenAccessDuration).Result()
	if err != nil {
		return "", fmt.Errorf("error while storing session identifier in redis: %v", err)
	}

	c.Header("Authorization", "Bearer "+tokenString)

	return fullToken, nil
}

func AuthMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			JERR(c, http.StatusUnauthorized, errors.New("authorization header is required"))
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 {
			JERR(c, http.StatusUnauthorized, errors.New("invalid token format"))
			c.Abort()
			return
		}

		tokenString := tokenParts[1]
		email := tokenParts[0]

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(appConfig.SecretKey), nil
		})

		if err != nil || !token.Valid {
			JERR(c, http.StatusUnauthorized, errors.New("invalid token"))
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			JERR(c, http.StatusUnauthorized, errors.New("invalid token claims"))
			c.Abort()
			return
		}

		encEmail, _ := Encrypt(email, "s")
		exists, err := rdb.Exists(context.Background(), encEmail).Result()
		if err != nil {
			JERR(c, http.StatusUnauthorized, errors.New("failed to check session identifier in Redis"))
			c.Abort()
			return
		}

		if exists > 0 {
			JERR(c, http.StatusForbidden, errors.New("user already logged in"))
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func Refresh(c *gin.Context) (string, error) {
	claims, ok := c.Get("claims")
	if !ok {
		return "", errors.New("failed to retrieve claims")
	}

	refreshClaims, ok := claims.(*Claims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	// uncomment this line if you want to give rule minimal time left 30 seconds token can be refreshed
	// expirationTime := refreshClaims.ExpiresAt.Time
	// if time.Until(expirationTime) > 30*time.Second {
	// 	return "", errors.New("token cannot be refreshed")
	// }

	newExpirationTime := time.Now().Add(appConfig.TokenRefreshDuration)
	refreshClaims.ExpiresAt = jwt.NewNumericDate(newExpirationTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	tokenString, err := token.SignedString([]byte(appConfig.SecretKey))
	if err != nil {
		return "", errors.New("internal server error")
	}

	c.Header("Authorization", "Bearer "+tokenString)

	return tokenString, nil
}

func RevokeToken(c *gin.Context, rdb *redis.Client) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is required")
	}

	tokenString := authHeader[len("Bearer "):]

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(appConfig.SecretKey), nil
	})

	if err != nil {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	_, err = rdb.Del(context.Background(), claims.Email).Result()
	if err != nil {
		return claims.Email, errors.New("error while deleting session identifier in redis")
	}

	revokedTokens[tokenString] = struct{}{}

	return claims.Email, nil
}
