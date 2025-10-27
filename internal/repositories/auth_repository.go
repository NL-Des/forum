package repositories

import (
	"database/sql"
	"fmt"
	"forum/internal/domain"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) domain.AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) UserExisting(username string) bool {
	var id int
	err := r.db.QueryRow(`
		SELECT id
		FROM users 
		WHERE username = ?`, username).Scan(&id)

	if err == sql.ErrNoRows {
		// Aucun utilisateur trouvé
		return false
	} else if err != nil {
		// Erreur réelle (connexion, SQL, etc.)
		fmt.Println("Erreur SQL:", err)
		return false
	}

	// Si on arrive ici, un utilisateur existe
	return true
}

func (r *authRepository) RegisterAuth(username, email string) error {
	if r.db == nil {
		return fmt.Errorf("database not initialized")
	}

	var (
		res sql.Result
		err error
	)

	if email == "" {
		res, err = r.db.Exec("INSERT INTO users (username) VALUES (?)", username)
	} else {
		res, err = r.db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username, email)
	}

	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	_, err = res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	return nil
}

func (r *authRepository) LoginAuth(username string) error {
	row := r.db.QueryRow("SELECT id, username FROM users WHERE username = ?", username)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username)
	if err != nil {
		return err
	}
	return nil
}

func (r *authRepository) LoginAuthByUsername(Token, username string) error {
	_, err := r.db.Exec("UPDATE users SET token = ? WHERE username = ? ;", Token, username)
	return err
}
