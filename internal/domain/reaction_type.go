package domain

import "time"

type Reaction struct {
	ID         int64
	Value      int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	TargetType string
	TargetID   int64
	UserID     int64
}

type ReactionRepository interface {
	CreateOrUpdate(r *Reaction) error
	Delete(userID int64, targetType string, targetID int64) error
	GetByTarget(targetType string, targetID int64) ([]Reaction, error)
	CountByTarget(targetType string, targetID int64) (likes int, dislikes int, err error)
}

type ReactionService interface {
	React(userID int64, targetType string, targetID int64, value int) error
	RemoveReaction(userID int64, targetType string, targetID int64) error
	GetReactions(targetType string, targetID int64) ([]Reaction, error)
	GetReactionCounts(targetType string, targetID int64) (likes int, dislikes int, err error)
}
