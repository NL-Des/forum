package domain

type User struct {
	ID       int64
	Username string
	Password string // (doit être "hashé")
	Email    string
}

type UserService interface {
	Register(username, email, password string) error
	Authenticate(email, password string) (*User, error)
	TokenLogIn(Token, email string) error
	Home(Token string) (*User, error)
	Logout(Token string) error
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByUsername(Logusername string) (*User, error)
	InsertToken(Token, email string) error
	GetUserByToken(Token string) (*User, error)
	DeleteTokenLog(Token string) error
}
