package services

import (
	"fmt"
	"forum/internal/domain"
)

type topicPostService struct {
	repo domain.TopicPostRepository
}

func NewTopicPostService(repo domain.TopicPostRepository) domain.TopicPostService {
	return &topicPostService{repo: repo}
}

func (s *topicPostService) GetAllTopics() ([]domain.Topic, error) {
	return s.repo.GetAllTopics()
}

func (s *topicPostService) CreateTopic(title, content string, userID int, categories_id []int) error {

	if title == "" {
		return fmt.Errorf("❌ the 'Title' field can't be empty")
	}
	// limiter la taille du contenu ??
	return s.repo.InsertTopic(title, content, userID, categories_id)
}

func (s *topicPostService) GetThreadByID(id int) (*domain.Thread, error) {
	topic, err := s.repo.GetTopicByID(id)
	if err != nil {
		return nil, err
	}

	posts, err := s.repo.GetPostsByTopicID(id)
	if err != nil {
		return nil, err
	}

	return &domain.Thread{
		Topic: *topic,
		Posts: posts,
	}, nil
}

func (s *topicPostService) AddPost(topicID int, content string, userID int) error {
	if content == "" {
		return fmt.Errorf("❌ the message can't be empty")
	}
	return s.repo.InsertPost(topicID, content, userID)
}

func (s *topicPostService) FilterTopic(UserId int) ([]domain.Topic, error) {
	topics, err := s.repo.GetTopicsByUserId(UserId)
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func (s *topicPostService) FilterByCategorie(CategorieName string) ([]domain.Topic, error) {
	topics, err := s.repo.GetTopicsByCategories(CategorieName)
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func (s *topicPostService) FilterByCategorieAndUserId(CategorieName string, UserId int) ([]domain.Topic, error) {
	topics, err := s.repo.GetTopicsByCategoriesAndUserId(CategorieName, UserId)
	if err != nil {
		return nil, err
	}
	return topics, nil
}
