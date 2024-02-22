package handlers

import (
	"context"
	"encoding/json"
	"github.com/vshigimoto/GoAuthJWTService/internal/entity"
	"net/http"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *Handlers) RefreshToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.l.Infof("Not allowed method")
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RefreshTokenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.l.Infof("Error decoding refresh token request: %v", err)
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}

	userGuid, err := h.services.CompareHashAndRefresh(context.Background(), req.RefreshToken)
	if err != nil {
		h.l.Infof("Error compare refresh token request: %v", err)
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}

	accessToken, err := h.services.GenerateAccessToken(context.Background(), userGuid)
	if err != nil {
		h.l.Error("Failed to generate access token")
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := h.services.GenerateRefreshToken(context.Background(), userGuid)
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
