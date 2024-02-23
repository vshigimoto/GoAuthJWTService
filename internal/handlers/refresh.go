package handlers

import (
	"context"
	"encoding/json"
	"github.com/vshigimoto/GoAuthJWTService/internal/entity"
	"net/http"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
	UserGuid     string `json:"user_guid"`
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

	err := h.services.CompareHashAndRefresh(context.Background(), req.RefreshToken, req.UserGuid)
	if err != nil {
		h.l.Infof("Error compare refresh token request: %v", err)
		http.Error(w, "Invalid refresh_token", http.StatusBadRequest)
		return
	}

	err = h.services.Delete(context.Background(), req.UserGuid)
	if err != nil {
		h.l.Infof("Error delete data from mongoDB: %v", err)
		http.Error(w, "Error delete data", http.StatusInternalServerError)
		return
	}

	accessToken, err := h.services.GenerateAccessToken(context.Background(), req.UserGuid)
	if err != nil {
		h.l.Error("Failed to generate access token")
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := h.services.GenerateRefreshToken(context.Background(), req.UserGuid)
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
