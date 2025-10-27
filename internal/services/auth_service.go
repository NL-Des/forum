package services

import (
	"fmt"
	"forum/internal/domain"
)

type authService struct {
	repo domain.AuthRepository
}

func NewAuthService(repo domain.AuthRepository) domain.AuthService {
	return &authService{repo: repo}
}

func (s *authService) GitHub(username, email string) error {
	Exist := s.repo.UserExisting(username)
	if Exist {
		err := s.repo.LoginAuth(username)
		if err != nil {
			fmt.Println("erreur login ")
			return err
		}
		return nil
	}
	err := s.repo.RegisterAuth(username, email)
	if err != nil {
		fmt.Println("erreur register")
		return err
	}
	return nil
}

func (s *authService) AuthToken(Token, username string) error {
	err := s.repo.LoginAuthByUsername(Token, username)
	if err != nil {
		return err
	}
	return nil
}
