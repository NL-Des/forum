package repositories

import (
	"database/sql"
	"forum/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

/*
constructeur: on crée une interface indirecte (UserRep)  (qui sera envoyée au service)
à partir du contenu concret de la BdD (userRep) qui ne doit pas être exporté directement
*/
func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

// créer nouvel utilisateur:
func (r *userRepository) Create(user *domain.User) error {
	res, err := r.db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	// Récupérer l’ID auto-généré
	id, err := res.LastInsertId()
	if err == nil {
		user.ID = id
	}
	return nil
}

// récupérer un utilisateur de la BdD par son ID:
func (r *userRepository) GetByID(id int) (*domain.User, error) {
	row := r.db.QueryRow("SELECT id, username, email, password FROM users WHERE id = ?", id)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// récupérer un utilisateur de la BdD par son email
func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	row := r.db.QueryRow("SELECT id, username, email, password FROM users WHERE email = ?", email)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
