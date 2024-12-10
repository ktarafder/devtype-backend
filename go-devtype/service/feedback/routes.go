package feedback

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ktarafder/devtype-backend/service/auth" 
	"github.com/ktarafder/devtype-backend/types"
	"github.com/ktarafder/devtype-backend/utils"
)

type Handler struct {
    store types.FeedbackStore
}

func NewHandler(store types.FeedbackStore) *Handler {
    return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/feedback", h.handleFeedback).Methods("POST")
}

func (h *Handler) handleFeedback(w http.ResponseWriter, r *http.Request) {
    // Extract UserID from the JWT
    userID, err := auth.GetUserIDFromJWT(r)
    if err != nil {
        utils.WriteError(w, http.StatusUnauthorized, err)
        return
    }

    // Parse the JSON payload
    var payload types.Feedback
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    // Attach the UserID to the feedback
    payload.UserID = userID

    // Store the feedback in the database
    err = h.store.CreateFeedback(payload)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    // Return success response
    utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Feedback saved successfully"})
}
