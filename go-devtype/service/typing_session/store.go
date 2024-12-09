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

func (s *Store) GetTypingSessionsByUserID(userID int) ([]types.TypingSession, error) {
	rows, err := s.db.Query(`
		SELECT overall_accuracy, overall_speed 
		FROM typing_session 
		WHERE user_id = ?`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []types.TypingSession
	for rows.Next() {
		var session types.TypingSession
		err := rows.Scan(&session.OverallAccuracy, &session.OverallSpeed)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}
