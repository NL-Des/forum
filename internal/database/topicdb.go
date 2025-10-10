package database

import (
	_ "github.com/mattn/go-sqlite3"
)

type Topic struct {
	ID      int
	Title   string
	Content string
}

func GetAllTopics() ([]Topic, error) {
	DB := OpenDB("forum.db")
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
