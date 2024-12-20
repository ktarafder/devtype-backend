package types

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
	DeleteUser(userID string) error
}
type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"-"`
	CreatedAt string `json:"created_at"`
	TotalScore float64 `json:"total_score"`
}
type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=108"`
}

type LoginUserPayload struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type TypingSession struct {
	SessionID       int     `json:"session_id"`
	OverallAccuracy float64 `json:"overall_accuracy"`
	OverallSpeed    float64 `json:"overall_speed"`
	UserID          int     `json:"user_id"`
	SnippetID       int     `json:"snippet_id"`
}

type TypingSessionPayload struct {
	OverallAccuracy float64 `json:"overall_accuracy" validate:"required"`
	OverallSpeed    float64 `json:"overall_speed" validate:"required"`
	UserID          int     `json:"user_id"`
	SnippetID       int     `json:"snippet_id" validate:"required"`
}

type TypingSessionStore interface {
	CreateTypingSession(session TypingSession) error
	GetTypingSessionsByUserID(userID int) ([]TypingSession, error)
}

type Feedback struct {
    ImprovementArea string `json:"improvement_area"`
    FeedbackText    string `json:"feedback_text"`
    UserID   int    `json:"user_id"`
}

type FeedbackStore interface {
    CreateFeedback(feedback Feedback) error
}
