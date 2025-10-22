package repositories

import (
	"database/sql"
	"forum/internal/domain"
)

type filterRepository struct {
	db *sql.DB
}

func NewFilterRepository(db *sql.DB) domain.FilterRepository {
	return &filterRepository{db: db}
}

func (r *filterRepository) GetTopicsByUserId(UserId int) ([]domain.Topic, error) {
	rows, err := r.db.Query(`
		SELECT id, title, content, created_at, updated_at, user_id 
		FROM topics 
		WHERE user_id = ? 
		ORDER BY created_at DESC`, UserId)
	if err != nil {
		/*fmt.Println("yoyo")*/
		return nil, err
	}
	defer rows.Close()

	var topics []domain.Topic
	for rows.Next() {
		var t domain.Topic
		err := rows.Scan(&t.ID, &t.Title, &t.Content, &t.CreatedAt, &t.UpdatedAt, &t.UserName)
		if err != nil {
			return nil, err
		}
		// Likes/Dislikes initialisés à 0
		t.Likes = 0
		t.Dislikes = 0
		topics = append(topics, t)
	}
	return topics, nil
}

func (r *filterRepository) GetTopicsByCategories(CategorieName string) ([]domain.Topic, error) {
	rows, err := r.db.Query(`
		SELECT t.id, t.title, t.content, t.created_at, t.updated_at, u.username
		FROM topics t
		JOIN users u ON t.user_id = u.id
		JOIN topic_categories tc ON t.id = tc.topic_id
		JOIN categories c ON tc.category_id = c.id
		WHERE c.name = ?
		ORDER BY t.created_at DESC;`, CategorieName)
	if err != nil {
		/*fmt.Println("yoyo")*/
		return nil, err
	}
	defer rows.Close()

	var topics []domain.Topic
	for rows.Next() {
		var t domain.Topic
		err := rows.Scan(&t.ID, &t.Title, &t.Content, &t.CreatedAt, &t.UpdatedAt, &t.UserName)
		if err != nil {
			return nil, err
		}
		// Likes/Dislikes initialisés à 0
		t.Likes = 0
		t.Dislikes = 0
		topics = append(topics, t)
	}
	return topics, nil
}

func (r *filterRepository) GetTopicsByCategoriesAndUserId(CategorieName string, UserId int) ([]domain.Topic, error) {
	rows, err := r.db.Query(`
		SELECT t.id, t.title, t.content, t.created_at, t.updated_at, u.username
		FROM topics t
		JOIN users u ON t.user_id = u.id
		JOIN topic_categories tc ON t.id = tc.topic_id
		JOIN categories c ON tc.category_id = c.id
		WHERE c.name = ? 
		  AND t.user_id = ?
		ORDER BY t.created_at DESC;`, CategorieName, UserId)
	if err != nil {
		/*fmt.Println("yoyo")*/
		return nil, err
	}
	defer rows.Close()

	var topics []domain.Topic
	for rows.Next() {
		var t domain.Topic
		err := rows.Scan(&t.ID, &t.Title, &t.Content, &t.CreatedAt, &t.UpdatedAt, &t.UserName)
		if err != nil {
			return nil, err
		}
		// Likes/Dislikes initialisés à 0
		t.Likes = 0
		t.Dislikes = 0
		topics = append(topics, t)
	}
	return topics, nil

}
