package typing_session

import (
	"fmt"
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/ktarafder/devtype-backend/service/auth"
	"github.com/ktarafder/devtype-backend/types"
	"github.com/ktarafder/devtype-backend/utils"
)

type Handler struct {
	store types.TypingSessionStore
}

func NewHandler(store types.TypingSessionStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/typing-session", h.handleTypingSession).Methods("POST")
	router.HandleFunc("/typing-sessions", h.handleGetTypingSessions).Methods("GET")
}

func (h *Handler) handleTypingSession(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserIDFromJWT(r)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	// Parse JSON payload
	var payload types.TypingSessionPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", errors))
		return
	}

	// Create new typing session record
	err = h.store.CreateTypingSession(types.TypingSession{
		OverallAccuracy: payload.OverallAccuracy,
		OverallSpeed:    payload.OverallSpeed,
		UserID:          userID,              // From JWT
		SnippetID:       payload.SnippetID,   // From frontend payload
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Typing session created successfully"})
}

func (h *Handler) handleGetTypingSessions(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserIDFromJWT(r)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	sessions, err := h.store.GetTypingSessionsByUserID(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch typing sessions: %s", err))
		return
	}

	response := make([]map[string]float64, len(sessions))
	for i, session := range sessions {
		response[i] = map[string]float64{
			"overall_accuracy": session.OverallAccuracy,
			"overall_speed":    session.OverallSpeed,
		}
	}

	utils.WriteJSON(w, http.StatusOK, response)
}


