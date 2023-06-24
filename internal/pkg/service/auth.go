package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/JamshedJ/goHR/internal/log"
	"github.com/golang-jwt/jwt/v5"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

func (a *App) GenerateToken(ctx context.Context, u models.User) (string, error) {
	if !u.Validate() {
		return "", models.ErrBadRequest
	}
	u.Password = generatePasswordHash(u.Password)
	err := a.db.AuthenticateUser(ctx, &u)
	if err != nil {
		if err == models.ErrNoRows {
			return "", models.ErrUnauthorized
		}
		log.Error.Println("GenerateToken a.db.AuthenticateUser", err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		User: u,
	})

	return token.SignedString([]byte(signingKey))
}

func (a *App) ParseToken(jwtString string) (user models.User, err error) {
	var claims models.JWTClaims
	token, err := jwt.ParseWithClaims(jwtString, &claims, func(token *jwt.Token) (interface{}, error) {
		if sm, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || sm.Name != "HS256" {
			log.Debug.Println("app ParseToken jwt.ParseWithClaims incorrect signing method", sm.Name)
			return nil, models.ErrUnauthorized
		}
		return []byte(signingKey), nil
	})
	if err != nil || !token.Valid {
		return user, models.ErrUnauthorized
	}
	return claims.User, nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
