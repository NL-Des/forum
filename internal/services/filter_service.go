package services

import (
	"forum/internal/domain"
)

type filterService struct {
	repo domain.FilterRepository
}

func NewFilterService(repo domain.FilterRepository) domain.FilterService {
	return &filterService{repo: repo}
}

func (s *filterService) FilterTopic(UserId int) ([]domain.Topic, error) {
	topics, err := s.repo.GetTopicsByUserId(UserId)
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func (s *filterService) FilterByCategorie(CategorieName string) ([]domain.Topic, error) {
	topics, err := s.repo.GetTopicsByCategories(CategorieName)
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func (s *filterService) FilterByCategorieAndUserId(CategorieName string, UserId int) ([]domain.Topic, error) {
	topics, err := s.repo.GetTopicsByCategoriesAndUserId(CategorieName, UserId)
	if err != nil {
		return nil, err
	}
	return topics, nil
}
