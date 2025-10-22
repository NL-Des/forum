package domain

type FilterRepository interface {
	GetTopicsByUserId(UserId int) ([]Topic, error)
	GetTopicsByCategories(CategorieName string) ([]Topic, error)
	GetTopicsByCategoriesAndUserId(CategorieName string, UserId int) ([]Topic, error)
	GetLikedTopicsByUser(userID int64) ([]Topic, error)
}

type FilterService interface {
	FilterTopic(UserId int) ([]Topic, error)
	FilterByCategorie(CategorieName string) ([]Topic, error)
	FilterByCategorieAndUserId(CategorieName string, UserId int) ([]Topic, error)
	GetLikedTopics(userID int64) ([]Topic, error)
}
