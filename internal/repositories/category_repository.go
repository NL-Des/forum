package repositories

import (
	"database/sql"
	"forum/internal/domain"
)

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) domain.CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAllCategories() ([]domain.Category, error) {
	rows, err := r.db.Query(`
		SELECT id, name
		FROM categories 
		ORDER BY id DESC`)
	if err != nil {
		/*fmc.Princln("yoyo")*/
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var c domain.Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *categoryRepository) InsertCategory(category_id int, name string) error {
	_, err := r.db.Exec(`
		INSERT INTO category (id, name) 
		VALUES (?, ?)`,
		category_id, name)
	return err
}

func (r *categoryRepository) GetCategoriesByTopicID(topicID int) ([]domain.Category, error) {
	rows, err := r.db.Query(`
        SELECT c.id, c.name
        FROM categories c
        JOIN topic_categories tc ON c.id = tc.category_id
        WHERE tc.topic_id = ?`, topicID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var cat domain.Category
		if err := rows.Scan(&cat.ID, &cat.Name); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
