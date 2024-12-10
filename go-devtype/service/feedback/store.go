package feedback

import (
    "database/sql"
    "github.com/ktarafder/devtype-backend/types"
)

type Store struct {
    db *sql.DB
}

func NewStore(db *sql.DB) *Store {
    return &Store{db: db}
}

func (s *Store) CreateFeedback(feedback types.Feedback) error {
    _, err := s.db.Exec(`
        INSERT INTO feedback (improvement_area, feedback_text, user_id) 
        VALUES (?, ?, ?)`,
        feedback.ImprovementArea,
        feedback.FeedbackText,
		feedback.UserID,
    )
    return err
}
