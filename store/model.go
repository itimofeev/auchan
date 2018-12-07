package store

type User struct {
	ID       int64
	Email    string `sql:",unique,notnul"`
	Password string `sql:",notnul"`

	Baskets []*Basket `pg:"many2many:share"`
}

type Basket struct {
	ID   int64
	Name string `pg:",notnull"`

	Users []*User `pg:"many2many:share"`
}

type Share struct {
	UserID   int64 `pg:",notnull"`
	BasketID int64 `pg:",notnull"`
}
