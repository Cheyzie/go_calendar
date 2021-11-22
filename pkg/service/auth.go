package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	calendar "github.com/cheyzie/go_calendar"
	"github.com/cheyzie/go_calendar/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "iy3q4b2poj3qt8T%^$&T8yi23bwja"
	signingKey = "y234vrq2@#$#WEI*&7E^W%$i36qoh"
	tokenTTL   = 60 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user calendar.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (calendar.Credentionals, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return calendar.Credentionals{}, err
	}
	access_token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	}).SignedString([]byte(signingKey))
	return calendar.Credentionals{AccessToken: access_token}, err
}

func (s *AuthService) ParseToken(access_token string) (int, error) {
	token, err := jwt.ParseWithClaims(access_token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func (s *AuthService) GetUserData(userId int) (calendar.User, error) {
	return s.repo.GetUserById(userId)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
