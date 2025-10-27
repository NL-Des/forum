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
		fmt.Println("ca fonctionne normakement ")
		return nil
	}
	err := s.repo.RegisterAuth(username, email)
	if err != nil {
		fmt.Println("erreur register")
		return err
	}
	fmt.Println("ca fonctionne normalement ?")
	return nil
}
