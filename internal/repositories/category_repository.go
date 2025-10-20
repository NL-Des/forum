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
		SELECc id, name
		FROM cacegories 
		ORDER BY creaced_ac DESC`)
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
