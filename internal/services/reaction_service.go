package services

import (
	"forum/internal/domain"
)

type reactionService struct {
	repo domain.ReactionRepository
}

func NewReactionService(repo domain.ReactionRepository) domain.ReactionService {
	return &reactionService{repo: repo}
}

func (s *reactionService) React(userID int64, targetType string, targetID int64, value int) error {
	reaction := &domain.Reaction{
		Value:      value,
		TargetType: targetType,
		TargetID:   targetID,
		UserID:     userID,
	}
	return s.repo.CreateOrUpdate(reaction)
}

func (s *reactionService) RemoveReaction(userID int64, targetType string, targetID int64) error {
	return s.repo.Delete(userID, targetType, targetID)
}

func (s *reactionService) GetReactions(targetType string, targetID int64) ([]domain.Reaction, error) {
	return s.repo.GetByTarget(targetType, targetID)
}

func (s *reactionService) GetReactionCounts(targetType string, targetID int64) (int, int, error) {
	return s.repo.CountByTarget(targetType, targetID)
}
