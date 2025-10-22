package domain

import "time"

type Topic struct {
	ID         int
	Title      string
	Content    string
	UserName   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Categories []Category
	Likes      int
	Dislikes   int
}

type Thread struct {
	Topic      Topic
	Posts      []Post
	Categories []Category
}

type Post struct {
	ID        int
	Content   string
	UserName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Likes     int
	Dislikes  int
}

type TopicPostRepository interface {
	GetAllTopics() ([]Topic, error)
	InsertTopic(title, content string, userID int, categories_id []int) error
	GetTopicByID(id int) (*Topic, error)
	GetPostsByTopicID(topicID int) ([]Post, error)
	InsertPost(topicID int, content string, userID int) error
	GetTopicsByUserId(UserId int) ([]Topic, error)
	GetTopicsByCategories(CategorieName string) ([]Topic, error)
	GetTopicsByCategoriesAndUserId(CategorieName string, UserId int) ([]Topic, error)
}

type TopicPostService interface {
	GetAllTopics() ([]Topic, error)
	CreateTopic(title, content string, userID int, categories_id []int) error
	GetThreadByID(id int) (*Thread, error)
	AddPost(topicID int, content string, userID int) error
	FilterTopic(UserId int) ([]Topic, error)
	FilterByCategorie(CategorieName string) ([]Topic, error)
	FilterByCategorieAndUserId(CategorieName string, UserId int) ([]Topic, error)
}
