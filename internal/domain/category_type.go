package domain

type Category struct {
	ID   int
	Name string
}

type CategoryRepository interface {
	GetAllCategories() ([]Category, error)
	InsertCategory(category_id int, name string) error
}

type CategoryService interface {
	GetAllCategories() ([]Category, error)
	CreateCategory(category_id int, name string) error
}
