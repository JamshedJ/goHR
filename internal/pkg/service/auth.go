package service

import (
	"context"
	"time"

	"github.com/JamshedJ/goHR/internal/configs"
	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

//var (
//	salt       = configs.App.Salt
//)

func (s *Service) GenerateToken(ctx context.Context, u models.User) (string, error) {

	if err := u.Validate(); err != nil {
		log.Warning.Println("service GenerateToken", err)
		return "", err
	}
	u.Password = models.GeneratePasswordHash(u.Password)
	err := s.db.AuthenticateUser(ctx, &u)
	if err != nil {
		if err == models.ErrNoRows {
			log.Warning.Println("service GenerateToken s.db.AuthenticateUser", err)
			return "", models.ErrUnauthorized
		}
		log.Error.Println("service GenerateToken s.db.AuthenticateUser", err)
		return "", err
	}

	tokenTTL := time.Duration(configs.App.TokenTTLHours) * time.Hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		User: u,
	})

	return token.SignedString([]byte(configs.App.SignKey))
}

func (s *Service) ParseToken(jwtString string) (user models.User, err error) {
	var claims models.JWTClaims
	token, err := jwt.ParseWithClaims(jwtString, &claims, func(token *jwt.Token) (interface{}, error) {
		if sm, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || sm.Name != "HS256" {
			log.Debug.Println("service ParseToken jwt.ParseWithClaims incorrect signing method", sm.Name)
			return nil, models.ErrUnauthorized
		}
		return []byte(configs.App.SignKey), nil
	})
	if err != nil || !token.Valid {
		return user, models.ErrUnauthorized
	}
	return claims.User, nil
}
