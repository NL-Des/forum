package database

import (
	_ "github.com/mattn/go-sqlite3" // ou ton driver
)

type Topic struct {
	ID      int
	Title   string
	Content string
}

func GetAllTopics() ([]Topic, error) {
	rows, err := DB.Query("SELECT id, title, content FROM topics ORDER BY id DESC")
	if err != nil {
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
	return topics, nil
}

func InsertTopic(title, content string, userID int) error {
	_, err := DB.Exec("INSERT INTO topics (title, content, category_id, user_id) VALUES (?, ?, ?, ?)", title, content, 1, userID)
	return err
}

func GetTopicByID(id int) (Topic, error) {
	var t Topic
	err := DB.QueryRow("SELECT id, title, content FROM topics WHERE id = ?", id).Scan(&t.ID, &t.Title, &t.Content)
	return t, err
}
