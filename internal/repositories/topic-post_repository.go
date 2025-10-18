package repositories

import (
	"database/sql"
	"fmt"
	"forum/internal/domain"
)

type TopicPostRepository struct {
	db *sql.DB
}

func NewTopicPostRepository(db *sql.DB) *TopicPostRepository {
	return &TopicPostRepository{db: db}
}

func (r *TopicPostRepository) GetAllTopics() ([]domain.Topic, error) {
	rows, err := r.db.Query("SELECT id, title, content FROM topics ORDER BY id DESC")
	if err != nil {
		fmt.Println("yoyo")
		return nil, err
	}
	defer rows.Close()

	var topics []domain.Topic
	for rows.Next() {
		var t domain.Topic
		err := rows.Scan(&t.ID, &t.Title, &t.Content)
		if err != nil {
			return nil, err
		}
		topics = append(topics, t)
	}
	return topics, nil
}

func (r *TopicPostRepository) InsertTopic(title, content string, userID int) error {
	_, err := r.db.Exec("INSERT INTO topics (title, content, category_id, user_id) VALUES (?, ?, ?, ?)", title, content, 1, userID)
	return err
}

func (r *TopicPostRepository) GetTopicByID(id int) (*domain.Topic, error) {
	var t domain.Topic
	err := r.db.QueryRow("SELECT id, title, content FROM topics WHERE id = ?", id).Scan(&t.ID, &t.Title, &t.Content)
	return &t, err
}

func (r *TopicPostRepository) GetPostsByTopicID(topicID int) ([]domain.Post, error) {
	rows, err := r.db.Query("SELECT id, content, created_at FROM messages WHERE topic_id = ? ORDER BY created_at ASC", topicID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		var p domain.Post
		err := rows.Scan(&p.ID, &p.Content, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func (r *TopicPostRepository) InsertPost(topicID int, content string, userID int) error {
	_, err := r.db.Exec("INSERT INTO messages (content, topic_id, user_id, created_at) VALUES (?, ?, ?, datetime('now'))", content, topicID, userID)
	return err
}
