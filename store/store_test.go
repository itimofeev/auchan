package store

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	s := NewStore("postgresql://postgres@db:5432/postgres?sslmode=disable")

	user, err := s.CreateUser("user1@gmail.com", "123")
	require.NoError(t, err)

	loaded, err := s.GetUserByEmail("adsfasdf@mail.ru")
	require.NoError(t, err)

	require.Equal(t, user.Email, loaded.Email)
	require.Equal(t, user.Password, loaded.Password)
}
