package service

import (
	"github.com/itimofeev/auchan/store"
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

func (s *Service) GetUserByEmail(email string) (user *store.User, err error) {
	return s.store.GetUserByEmail(email)
}

func (s *Service) GetUserByID(id int64) (user *store.User, err error) {
	return s.store.GetUserByID(id)
}

func (s *Service) CreateBasket(user *store.User, basketName string) (basket *store.Basket, err error) {
	return s.store.CreateBasket(user, basketName)
}

func (s *Service) GetUserBaskets(user *store.User) (baskets []*store.Basket, err error) {
	return s.store.GetUserBaskets(user)
}
