package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vshigimoto/GoAuthJWTService/internal/entity"
)

func (h *Handlers) GetTokens(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.l.Error("Not allowed method")
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	userGUID := r.URL.Query().Get("user_guid")
	if userGUID == "" {
		h.l.Error("Empty GUID")
		http.Error(w, "User GUID is required", http.StatusBadRequest)
		return
	}

	accessToken, err := h.services.GenerateAccessToken(context.Background(), userGUID)
	if err != nil {
		h.l.Error("Failed to generate access token")
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := h.services.GenerateRefreshToken(context.Background(), userGUID)
	if err != nil {
		h.l.Error("Failed to generate refresh token")
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	response := entity.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	h.l.Info("Success create access token")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
