package services

import (
	"errors"
	"fmt"
	"forum/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(username, email, password string) error {
	existing, err := s.repo.GetByEmail(email)
	if err == nil && existing != nil {
		return errors.New("❌ email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := domain.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.repo.Create(&user)
}

func (s *userService) Authenticate(email, password string) (*domain.User, error) {
	user, err := s.repo.GetByEmail(email)
	fmt.Println(user.Email)
	fmt.Println(user.Password)
	if err != nil {
		fmt.Println("❌ invalid email")
		return nil, errors.New("❌ invalid email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("❌ invalid password")
		return nil, errors.New("❌ invalid password")
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
