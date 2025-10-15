package database

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Topic struct {
	ID      int
	Title   string
	Content string
}

type Datas struct {
	Topics     []Topic
	IsLoggedIn bool
}

type Thread struct {
	Topic Topic
	Posts []Post
}

type Post struct {
	ID        int
	Content   string
	CreatedAt string
}

func GetAllTopics() ([]Topic, error) {
	DB := OpenDB("forum.db")
	rows, err := DB.Query("SELECT id, title, content FROM topics ORDER BY id DESC")
	if err != nil {
		fmt.Println("yoyo")
		return nil, err
	}
	defer rows.Close()

	var topics []Topic
	for rows.Next() {
		var t Topic
		err := rows.Scan(&t.ID, &t.Title, &t.Content)
		if err != nil {
			return nil, err
		}
		topics = append(topics, t)
	}
	DB.Close()
	return topics, nil
}

func InsertTopic(title, content string, userID int) error {
	DB := OpenDB("forum.db")
	_, err := DB.Exec("INSERT INTO topics (title, content, category_id, user_id) VALUES (?, ?, ?, ?)", title, content, 1, userID)
	DB.Close()
	return err
}

func GetTopicByID(id int) (Topic, error) {
	var t Topic
	DB := OpenDB("forum.db")
	err := DB.QueryRow("SELECT id, title, content FROM topics WHERE id = ?", id).Scan(&t.ID, &t.Title, &t.Content)
	DB.Close()
	return t, err
}

func GetPostsByTopicID(topicID int) ([]Post, error) {
	DB := OpenDB("forum.db")
	rows, err := DB.Query("SELECT id, content, created_at FROM messages WHERE topic_id = ? ORDER BY created_at ASC", topicID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.Content, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	DB.Close()
	return posts, nil
}

func InsertPost(topicID int, content string, userID int) error {
	DB := OpenDB("forum.db")
	_, err := DB.Exec("INSERT INTO messages (content, topic_id, user_id, created_at) VALUES (?, ?, ?, datetime('now'))", content, topicID, userID)
	DB.Close()
	return err
}
