package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/itimofeev/auchan/server/models"
	"github.com/itimofeev/auchan/store"
	"github.com/itimofeev/auchan/util"
	"time"
)

const secret = "some secret"

func NewService(store *store.Store) *Service {
	return &Service{
		store: store,
	}
}

type Service struct {
	store *store.Store
}

func (s *Service) LoginUser(email, password string) (string, error) {
	user, err := s.store.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	// TODO check pass
	return authTokenForUserID(user.ID), nil
}

func authTokenForUserID(userID int64) string {
	claims := &Claims{
		UserID: userID,
	}

	return signAndStringifyToken(newJWTToken(claims, time.Hour*100), secret)
}

func newJWTToken(claims *Claims, ttl time.Duration) *jwt.Token {
	currentTime := time.Now()
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: currentTime.Add(ttl).Unix(),
		IssuedAt:  currentTime.Unix(),
		Issuer:    "Auchan",
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func signAndStringifyToken(token *jwt.Token, secret string) string {
	tokenString, err := token.SignedString([]byte(secret))
	util.CheckErr(err, "token.SignedString")

	return tokenString
}

func (s *Service) GetUserByEmail(email string) (user *store.User, err error) {
	return s.store.GetUserByEmail(email)
}

func (s *Service) GetUserByID(id int64) (user *store.User, err error) {
	return s.store.GetUserByID(id)
}

func (s *Service) GetByTokenEnsure(token string) (*models.User, error) {
	claims, err := parseJWT(token)

	if err != nil {
		if _, ok := err.(*jwt.ValidationError); ok {
			return nil, util.NewUnauthorized("auth.token.validation.error", "bad authenticate token")
		}
		return nil, err
	}
	u, err := s.store.GetUserByID(claims.UserID)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, util.NewUnauthorized("error.users.not.found.by.id", "user not found by id")
	}
	return &models.User{
		ID:    u.ID,
		Email: u.Email,
	}, nil
}

type Claims struct {
	UserID int64 `json:"userId"`
	jwt.StandardClaims
}

func parseJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		b := []byte(secret)
		return b, nil
	})
	return claims, err
}
