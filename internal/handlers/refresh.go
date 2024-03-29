package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vshigimoto/GoAuthJWTService/internal/entity"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
	UserGUID     string `json:"user_guid"`
}

// RefreshToken handles the refresh token operation, generating new access and refresh tokens.
// @Summary      Refresh token
// @Description  Generates new access and refresh tokens based on the provided refresh token and user GUID.
// @Accept       json
// @Produce      json
// @Param        requestBody  body      RefreshTokenRequest  true  "Refresh Token Request"
// @Success      200          {object}  entity.TokenResponse  "A new pair of access and refresh tokens"
// @Failure      400          {object}  string               "Error decoding request or Invalid refresh_token"
// @Failure      405          {object}  string               "Only POST method is allowed"
// @Failure      500          {object}  string               "Failed to generate tokens"
// @Failure      500          {object}  string               "Error delete data"
// @Router       /token/refresh [post]
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

	err := h.services.CompareHashAndRefresh(r.Context(), req.RefreshToken, req.UserGUID)
	if err != nil {
		h.l.Infof("Error compare refresh token request: %v", err)
		http.Error(w, "Invalid refresh_token", http.StatusBadRequest)
		return
	}

	err = h.services.Delete(r.Context(), req.UserGUID)
	if err != nil {
		h.l.Infof("Error delete data from mongoDB: %v", err)
		http.Error(w, "Error delete data", http.StatusInternalServerError)
		return
	}

	accessToken, err := h.services.GenerateAccessToken(r.Context(), req.UserGUID)
	if err != nil {
		h.l.Error("Failed to generate access token")
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := h.services.GenerateRefreshToken(r.Context(), req.UserGUID)
	if err != nil {
		h.l.Error("Failed to generate refresh token")
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	response := entity.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	h.l.Info("Success create access and refresh token pair")
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		h.l.Error("Failed to encode")
		return
	}
}
