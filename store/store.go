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

	store := &Store{
		db: db,
	}

	db.Exec(`ALTER TABLE shares
  ADD CONSTRAINT shares_unique UNIQUE (user_id, basket_id)`)

	store.CreateUser("user1@gmail.com", "123")
	store.CreateUser("user2@gmail.com", "123")

	store.db.Insert(&Product{
		Name:       "product1",
		CategoryID: 10,
		ImageURL:   "http://google.com",
	})

	store.db.Insert(&Product{
		Name:       "profit2",
		CategoryID: 20,
		ImageURL:   "http://google.com",
	})

	return store
}

type Store struct {
	db *pg.DB
}

func createSchema(db *pg.DB) error {
	orm.RegisterTable((*Share)(nil))

	for _, mdl := range []interface{}{
		(*User)(nil),
		(*Basket)(nil),
		(*Share)(nil),
		(*Product)(nil),
		(*Goods)(nil),
	} {
		err := db.CreateTable(mdl, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
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

func (s *Store) GetUserBaskets(user *User) (baskets []*Basket, err error) {
	_, err = s.db.Query(&baskets, `
	SELECT
		b.*
	FROM
		baskets b
		JOIN shares s on b.id = s.basket_id
	WHERE
		s.user_id = ?
`, user.ID)
	return baskets, err
}

func (s *Store) CreateBasket(user *User, basketName string) (*Basket, error) {
	basket := &Basket{
		Name:  basketName,
		Users: []*User{user},
	}

	return basket, s.db.RunInTransaction(func(tx *pg.Tx) error {
		if err := tx.Insert(basket); err != nil {
			return err
		}

		return tx.Insert(&Share{
			BasketID: basket.ID,
			UserID:   user.ID,
		})
	})
}

func (s *Store) SearchProducts(name string) (products []*Product, err error) {
	return products, s.db.Model(&products).Where("name like ?", name+"%").Select()
}

func (s *Store) GetGoodsForBasket(basket *Basket) (goods []*Goods, err error) {
	return goods, s.db.Model(&goods).Where("basket_id = ?", basket.ID).Select()
}

func (s *Store) GetSharesForBasket(basket *Basket) (shares []*Share, err error) {
	return shares, s.db.Model(&shares).Relation("User").Where("basket_id = ?", basket.ID).Select()
}

func (s *Store) AddUserToShare(basket *Basket, email string) (share *Share, err error) {
	user := &User{}
	err = s.db.Model(user).Where("email = ? ", email).Select()
	if err != nil {
		return nil, err
	}

	sh := &Share{
		UserID:   user.ID,
		User:     user,
		BasketID: basket.ID,
	}
	return sh, s.db.Insert(sh)
}
