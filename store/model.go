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

type Product struct {
	ID         int64
	Name       string `pg:",notnull"`
	ImageURL   string `pg:",notnull"`
	CategoryID int64  `pg:",notnull"`
}

type Goods struct {
	ID        int64
	Product   Product `pg:",notnull"`
	Completed bool    `pg:",notnull"`
	Quantity  int     `pg:",notnull"`
	Price     int     `pg:",notnull"`
	Unit      string  `pg:",notnull"`
}
