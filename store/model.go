package store

type User struct {
	ID       int64
	Email    string `sql:",unique,notnull"`
	Password string `sql:",notnull"`

	Baskets []*Basket `pg:"many2many:share"`
}

type Basket struct {
	ID   int64
	Name string `sql:",notnull"`

	Users []*User `pg:"many2many:share"`
}

type Share struct {
	UserID   int64 `sql:",pk,notnull"`
	User     *User
	BasketID int64 `sql:",pk,notnull"`
}

type Product struct {
	ID         int64
	Name       string `sql:",notnull"`
	ImageURL   string `sql:",notnull"`
	CategoryID int64  `sql:",notnull"`
}

type Goods struct {
	ID        int64
	BasketID  int64 `sql:",notnull"`
	Product   *Product
	ProductID int64  `sql:",notnull"`
	Completed bool   `sql:",notnull"`
	Quantity  int64  `sql:",notnull"`
	Price     int64  `sql:",notnull"`
	Unit      string `sql:",notnull"`
}
