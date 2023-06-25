package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/JamshedJ/goHR/internal/configs"
	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

var (
	salt       = configs.App.Salt
	signingKey = configs.App.SignKey
	tokenTTL   = time.Duration(configs.App.TokenTTLHours) * time.Hour
)

func (s *Service) GenerateToken(ctx context.Context, u models.User) (string, error) {
	if err := u.Validate(); err != nil {
		log.Warning.Println("service GenerateToken", err)
		return "", err
	}
	u.Password = generatePasswordHash(u.Password)
	err := s.db.AuthenticateUser(ctx, &u)
	if err != nil {
		if err == models.ErrNoRows {
			log.Warning.Println("service GenerateToken s.db.AuthenticateUser", err)
			return "", models.ErrUnauthorized
		}
		log.Error.Println("service GenerateToken s.db.AuthenticateUser", err)
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

func (s *Service) ParseToken(jwtString string) (user models.User, err error) {
	var claims models.JWTClaims
	token, err := jwt.ParseWithClaims(jwtString, &claims, func(token *jwt.Token) (interface{}, error) {
		if sm, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || sm.Name != "HS256" {
			log.Debug.Println("service ParseToken jwt.ParseWithClaims incorrect signing method", sm.Name)
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
