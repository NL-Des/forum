package domain

type Topic struct {
	ID        int
	Title     string
	Content   string
	CreatedAt string
	UpdatedAt string
	Likes     int
	Dislikes  int
}

type Thread struct {
	Topic Topic
	Posts []Post
}

type Post struct {
	ID        int
	Content   string
	CreatedAt string
	UpdatedAt string
	Likes     int
	Dislikes  int
}

type TopicPostRepository interface {
	GetAllTopics() ([]Topic, error)
	InsertTopic(title, content string, userID int) error
	GetTopicByID(id int) (*Topic, error)
	GetPostsByTopicID(topicID int) ([]Post, error)
	InsertPost(topicID int, content string, userID int) error
}

type TopicPostService interface {
	GetAllTopics() ([]Topic, error)
	CreateTopic(title, content string, userID int) error
	GetThreadByID(id int) (*Thread, error)
	AddPost(topicID int, content string, userID int) error
}
