package leaderboard

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ktarafder/devtype-backend/service/auth" // Import your auth package for JWT handling
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/game/finish", h.updateLeaderboard).Methods("POST")
}

type GameFinishPayload struct {
	Score float64 `json:"score"`
}

func (h *Handler) updateLeaderboard(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the JWT in the Authorization header
	userID, err := auth.GetUserIDFromJWT(r)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Parse the JSON payload
	var payload GameFinishPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the score
	if payload.Score <= 0 {
		http.Error(w, "Score must be greater than 0", http.StatusBadRequest)
		return
	}

	// Update the user's total score
	tx, err := h.db.Begin() // Use a transaction
	if err != nil {
		http.Error(w, "Failed to begin transaction", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec(`UPDATE users SET total_score = total_score + ? WHERE id = ?`, payload.Score, userID)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to update user score", http.StatusInternalServerError)
		return
	}

	// Update the leaderboard
	_, err = tx.Exec(`
		WITH ranked_users AS (
			SELECT 
				id AS user_id, 
				total_score, 
				ROW_NUMBER() OVER (ORDER BY total_score DESC) AS rank
			FROM users
			WHERE total_score > 0
			LIMIT 3
		)
		UPDATE leaderboard AS lb
		JOIN ranked_users AS ru ON lb.rank = ru.rank
		SET lb.user_id = ru.user_id, 
			lb.total_score = ru.total_score, 
			lb.updated_at = CURRENT_TIMESTAMP;
	`)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to update leaderboard", http.StatusInternalServerError)
		return
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Leaderboard updated successfully",
	})
}
