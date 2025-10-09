package user

type User struct {
	ID       int64
	Username string
	Password string // (doit être "hashé")
	Email    string
}
