package snippets

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/snippets", h.getSnippetByDifficulty).Methods("GET")
}

func (h *Handler) getSnippetByDifficulty(w http.ResponseWriter, r *http.Request) {
	// Get the difficulty query parameter
	difficulty := r.URL.Query().Get("difficulty")
	if difficulty == "" {
		http.Error(w, "difficulty query parameter is required", http.StatusBadRequest)
		return
	}

	// Validate difficulty value
	if difficulty != "easy" && difficulty != "medium" && difficulty != "hard" {
		http.Error(w, "difficulty must be 'easy', 'medium', or 'hard'", http.StatusBadRequest)
		return
	}

	// Query the database
	query := `SELECT snippet_id, code_language, difficulty_level, snippet_text 
              FROM code_snippets 
              WHERE difficulty_level = ? 
              ORDER BY RAND() 
              LIMIT 1`
	row := h.db.QueryRow(query, difficulty)

	// Scan the result into a struct
	var snippet Snippet
	err := row.Scan(&snippet.ID, &snippet.Language, &snippet.Difficulty, &snippet.Text)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "no snippets found for the specified difficulty", http.StatusNotFound)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// Return the snippet as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snippet)
}

type Snippet struct {
	ID         int    `json:"id"`
	Language   string `json:"language"`
	Difficulty string `json:"difficulty"`
	Text       string `json:"text"`
}
