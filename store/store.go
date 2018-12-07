package store

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/itimofeev/auchan/util"
	"time"
)

func NewStore(connectURL string) *Store {
	opts, err := pg.ParseURL(connectURL)
	util.CheckErr(err, "pg.ParseURL")
	db := pg.Connect(opts)
	util.CheckErr(createSchema(db))

	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}
		util.Log.Printf("%s %s", time.Since(event.StartTime), query)
	})

	return &Store{
		db: db,
	}
}

type Store struct {
	db *pg.DB
}

func createSchema(db *pg.DB) error {
	err := db.CreateTable((*User)(nil), &orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateUser(email, password string) (*User, error) {
	user := &User{
		Email:    email,
		Password: password,
	}
	return user, s.db.Insert(user)
}

func (s *Store) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	err := s.db.Model(user).
		Where("email = ?", email).
		Select()
	return user, err
}

func (s *Store) GetUserByID(id int64) (*User, error) {
	user := &User{ID: id}
	err := s.db.Select(user)
	return user, err
}
