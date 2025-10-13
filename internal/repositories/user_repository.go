package repositories

import (
	"database/sql"
	"forum/internal/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id int) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	GetByUsername(Logusername string) (*domain.User, error)
	InsertToken(Token, email string) error
	GetUserByToken(Token string) (*domain.User, error)
	DeleteTokenLog(Token string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		user.Username, user.Email, user.Password)
	return err
}

func (r *userRepository) GetByID(id int) (*domain.User, error) {
	row := r.db.QueryRow("SELECT id, username, email, password FROM users WHERE id = ?", id)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	row := r.db.QueryRow("SELECT id, username, email, password FROM users WHERE email = ?", email)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *userRepository) GetByUsername(Logusername string) (*domain.User, error) {
	row := r.db.QueryRow("SELECT id, username, email, password FROM users WHERE username = ?", Logusername)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) InsertToken(Token, email string) error {
	_, err := r.db.Exec("UPDATE users SET token = ? WHERE email = ? ;", Token, email)
	return err
}

func (r *userRepository) GetUserByToken(Token string) (*domain.User, error) {
	row := r.db.QueryRow("SELECT username FROM users WHERE Token = ?", Token)
	user := &domain.User{}
	err := row.Scan(&user.Username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) DeleteTokenLog(Token string) error {
	_, err := r.db.Exec("UPDATE users SET token = NULL WHERE Token = ?", Token)
	if err != nil {
		return nil
	}
	return err
}
