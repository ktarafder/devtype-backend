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
	router.HandleFunc("/snippets", h.getSnippetsByDifficulty).Methods("GET")
}

func (h *Handler) getSnippetsByDifficulty(w http.ResponseWriter, r *http.Request) {
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

	// Query the database for 3 random snippets
	query := `SELECT snippet_id, code_language, difficulty_level, snippet_text 
              FROM code_snippets 
              WHERE difficulty_level = ? 
              ORDER BY RAND() 
              LIMIT 5`
	rows, err := h.db.Query(query, difficulty)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Collect the results into a slice
	var snippets []Snippet
	for rows.Next() {
		var snippet Snippet
		err := rows.Scan(&snippet.ID, &snippet.Language, &snippet.Difficulty, &snippet.Text)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		snippets = append(snippets, snippet)
	}

	// Return the snippets as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snippets)
}

type Snippet struct {
	ID         int    `json:"id"`
	Language   string `json:"language"`
	Difficulty string `json:"difficulty"`
	Text       string `json:"text"`
}
