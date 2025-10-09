package repositories

import (
	"database/sql"
	"forum/internal/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id int) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
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
