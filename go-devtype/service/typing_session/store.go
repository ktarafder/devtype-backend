package typing_session

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

func (s *Store) CreateTypingSession(session types.TypingSession) error {
	// SQL statement to insert a new typing session
	query := `
		INSERT INTO typing_session (overall_accuracy, overall_speed, user_id, snippet_id)
		VALUES (?, ?, ?, ?)
	`

	// Execute the query with the session data
	_, err := s.db.Exec(query, session.OverallAccuracy, session.OverallSpeed, session.UserID, session.SnippetID)
	if err != nil {
		return err
	}
	return nil
}
