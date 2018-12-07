package store

type User struct {
	ID       int64
	Email    string `sql:",unique"`
	Password string
}
