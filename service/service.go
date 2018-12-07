package service

import "github.com/itimofeev/auchan/store"

func NewService(store *store.Store) *Service {
	return &Service{
		store: store,
	}
}

type Service struct {
	store *store.Store
}

func (s *Service) CheckPassword(email, password string) (string, error) {
	_, err := s.store.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	return "hi, there!", nil
}
