package types

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
}
type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"-"`
	CreatedAt string `json:"created_at"`
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
	UserID          int     `json:"user_id" validate:"required"`
	SnippetID       int     `json:"snippet_id" validate:"required"`
}

type TypingSessionStore interface {
	CreateTypingSession(session TypingSession) error
}
