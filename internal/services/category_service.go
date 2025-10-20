package services

import (
	"fmt"
	"forum/internal/domain"
)

type categoryService struct {
	repo domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) domain.CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetAllCategories() ([]domain.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *categoryService) CreateCategory(id int, name string) error {
	if name == "" {
		return fmt.Errorf("❌ the 'name' field can't be empty")
	}
	return s.repo.InsertCategory(id, name)
}
