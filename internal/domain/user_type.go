package domain

type User struct {
	ID       int64
	Username string
	Password string // (doit être "hashé")
	Email    string
}

type UserRepository interface {
	Create(user *User) error
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetByUsername(Logusername string) (*User, error)
	InsertToken(Token, email string) error
	GetUserByToken(Token string) (*User, error)
	DeleteTokenLog(Token string) error
}

type UserService interface {
	Register(username, email, password string) error
	Authenticate(email, password string) (*User, error)
	TokenLogIn(Token, email string) error
	Home(Token string) (*User, error)
	Logout(Token string) error
}
