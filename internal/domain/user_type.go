package domain

type User struct {
	ID       int64
	Username string
	Password string // (doit être "hashé")
	Email    string
}

/*
type UserRepository interface {
	Create(user *User) error
	GetByID(id int) (*User, error)
	GetByEmail(email string) (*User, error)
}

type UserService interface {
	Register(username, email, password string) error
	Authenticate(email, password string) (*User, error)
}
*/
