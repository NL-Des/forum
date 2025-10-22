package repositories

import (
	"database/sql"
	"forum/internal/domain"
)

type reactionRepository struct {
	db *sql.DB
}

func NewReactionRepository(db *sql.DB) domain.ReactionRepository {
	return &reactionRepository{db: db}
}

func (r *reactionRepository) CreateOrUpdate(reaction *domain.Reaction) error {
	_, err := r.db.Exec(`
		INSERT INTO reactions (value, target_type, target_id, user_id)
		VALUES (?, ?, ?, ?)
		ON CONFLICT(target_type, target_id, user_id)
		DO UPDATE SET value=excluded.value, updated_at=CURRENT_TIMESTAMP
	`, reaction.Value, reaction.TargetType, reaction.TargetID, reaction.UserID)
	return err
}

func (r *reactionRepository) Delete(userID int64, targetType string, targetID int64) error {
	_, err := r.db.Exec(`
		DELETE FROM reactions WHERE user_id=? AND target_type=? AND target_id=?
	`, userID, targetType, targetID)
	return err
}

func (r *reactionRepository) GetByTarget(targetType string, targetID int64) ([]domain.Reaction, error) {
	rows, err := r.db.Query(`
		SELECT id, value, target_type, target_id, user_id, created_at, updated_at
		FROM reactions WHERE target_type=? AND target_id=?
	`, targetType, targetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reactions []domain.Reaction
	for rows.Next() {
		var rec domain.Reaction
		if err := rows.Scan(&rec.ID, &rec.Value, &rec.TargetType, &rec.TargetID, &rec.UserID, &rec.CreatedAt, &rec.UpdatedAt); err != nil {
			return nil, err
		}
		reactions = append(reactions, rec)
	}
	return reactions, nil
}

func (r *reactionRepository) CountByTarget(targetType string, targetID int64) (int, int, error) {
	row := r.db.QueryRow(`
		SELECT 
		COALESCE(SUM(CASE WHEN value=1 THEN 1 ELSE 0 END), 0) as likes,
		COALESCE(SUM(CASE WHEN value=-1 THEN 1 ELSE 0 END), 0) as dislikes
		FROM reactions WHERE target_type=? AND target_id=?
	`, targetType, targetID)

	var likes, dislikes int
	err := row.Scan(&likes, &dislikes)
	return likes, dislikes, err
}
