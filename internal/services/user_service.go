package services

import (
	"errors"
	"fmt"
	"forum/internal/domain"
	"forum/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(username, email, password string) error
	Authenticate(email, password string) (*domain.User, error)
	TokenLogIn(email, Token string) error
	Home(Token string) (*domain.User, error)
	Logout(Token string) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(username, email, password string) error {
	existing, _ := s.repo.GetByEmail(email)
	if existing != nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.repo.Create(user)
}

func (s *userService) Authenticate(email, password string) (*domain.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	fmt.Println("hashage de password")
	if err != nil {
		fmt.Println("mauvais hashage de password")

		return nil, errors.New("invalid email or password")

	}
	fmt.Println("hashage de password réussi")

	return user, nil
}

func (s *userService) TokenLogIn(Token, email string) error {
	err := s.repo.InsertToken(Token, email)
	if err != nil {
		return err
	}
	return nil
}
func (s *userService) Home(Token string) (*domain.User, error) {
	user, err := s.repo.GetUserByToken(Token)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Logout(Token string) error {
	err := s.repo.DeleteTokenLog(Token)
	if err != nil {
		return err
	}
	return nil
}
