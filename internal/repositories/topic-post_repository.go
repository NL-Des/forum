package repositories

import (
	"database/sql"
	"forum/internal/domain"
)

type topicPostRepository struct {
	db *sql.DB
}

func NewTopicPostRepository(db *sql.DB) domain.TopicPostRepository {
	return &topicPostRepository{db: db}
}

func (r *topicPostRepository) GetAllTopics() ([]domain.Topic, error) {
	rows, err := r.db.Query(`
		SELECT id, title, content, created_at, updated_at, user_id 
		FROM topics 
		ORDER BY created_at DESC`)
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

func (r *topicPostRepository) InsertTopic(title, content string, userID int, categories_id []int) error {
	result, err := r.db.Exec(`
		INSERT INTO topics (title, content, created_at, updated_at, user_id) 
		VALUES (?, ?, datetime('now'), datetime('now'), ?)`,
		title, content, userID)
	if err == nil {
		topicID, err := result.LastInsertId()
		if err != nil {
			return err
		}
		for _, category_id := range categories_id {
			_, err = r.db.Exec(`
				INSERT INTO topic_categories (topic_id, category_id) 
				VALUES (?, ?)`,
				topicID, category_id)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func (r *topicPostRepository) GetTopicByID(id int) (*domain.Topic, error) {
	var t domain.Topic
	err := r.db.QueryRow(`
		SELECT 
			topics.id, 
			topics.title, 
			topics.content, 
			topics.created_at, 
			topics.updated_at, 
			users.username
		FROM topics
		JOIN users ON topics.user_id = users.id
		WHERE topics.id = ?
		ORDER BY topics.created_at ASC`, id).
		Scan(&t.ID, &t.Title, &t.Content, &t.CreatedAt, &t.UpdatedAt, &t.UserName)
	return &t, err
}

func (r *topicPostRepository) GetPostsByTopicID(topicID int) ([]domain.Post, error) {
	rows, err := r.db.Query(`
		SELECT 
			posts.id, 
			posts.content, 
			posts.created_at, 
			posts.updated_at,
			users.username 
		FROM posts 
		JOIN users ON posts.user_id = users.id 
		WHERE posts.topic_id = ? 
		ORDER BY posts.created_at ASC`, topicID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		var p domain.Post
		err := rows.Scan(&p.ID, &p.Content, &p.CreatedAt, &p.UpdatedAt, &p.UserName)
		if err != nil {
			return nil, err
		}
		p.Likes = 0
		p.Dislikes = 0
		posts = append(posts, p)
	}
	return posts, nil
}

func (r *topicPostRepository) InsertPost(topicID int, content string, UserName int) error {
	_, err := r.db.Exec(`
		INSERT INTO posts (content, created_at, updated_at, topic_id, user_id) 
		VALUES (?, datetime('now'), datetime('now'), ?, ?)`,
		content, topicID, UserName)
	return err
}

func (r *topicPostRepository) GetTopicsByUserId(UserId int) ([]domain.Topic, error) {
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

func (r *topicPostRepository) GetTopicsByCategories(CategorieName string) ([]domain.Topic, error) {
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

func (r *topicPostRepository) GetTopicsByCategoriesAndUserId(CategorieName string, UserId int) ([]domain.Topic, error) {
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
