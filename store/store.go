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
	db.Exec(`ALTER TABLE goods
  ADD CONSTRAINT goods_unique UNIQUE (user_id, basket_id, goods_id)`)

	store.CreateUser("user1@gmail.com", "123")
	store.CreateUser("user2@gmail.com", "123")

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

func (s *Store) SearchProducts(name *string) (products []*Product, err error) {
	query := s.db.Model(&products)
	if name != nil {
		query = query.Where("name like ?", *name+"%")
	}

	return products, query.Limit(20).Select()
}

func (s *Store) GetGoodsForBasket(basket *Basket) (goods []*Goods, err error) {
	return goods, s.db.Model(&goods).Relation("Product").Relation("User").Where("basket_id = ?", basket.ID).Select()
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

func (s *Store) UpdateGoodsInBasket(user *User, basket *Basket, productId, quantity int64) (goods *Goods, err error) {
	goods = &Goods{
		ProductID: productId,
		BasketID:  basket.ID,
		UserID:    user.ID,
		User:      user,
	}

	return goods, s.db.RunInTransaction(func(tx *pg.Tx) error {
		product := &Product{ID: productId}
		err = s.db.Model(product).Where("id = ?", productId).Select()
		if err != nil {
			return err
		}

		goods.Product = product

		if quantity <= 0 {
			_, err := s.db.Model(goods).Where("basket_id = ? AND product_id = ? AND user_id = ?", basket.ID, productId, user.ID).Delete()
			return err
		}

		exists := true
		err = s.db.Model(goods).Where("basket_id = ? AND product_id = ? AND user_id = ?", basket.ID, productId, user.ID).Select()
		if err == pg.ErrNoRows {
			err = nil
			exists = false
		}
		if err != nil {
			return err
		}

		goods.Quantity = quantity
		goods.Price = int64(product.Price * 100)
		goods.UserID = user.ID
		goods.User = user

		if exists {
			return s.db.Update(goods)
		}

		return s.db.Insert(goods)
	})
}
