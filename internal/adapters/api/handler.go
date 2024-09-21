package api

import (
	"encoding/json"
	"net/http"

	"GoGateway/internal/app"
	"GoGateway/util"
)

type Handler struct {
	AuthService app.AuthService
	Logger      util.Logger
}

func NewHandler(authService app.AuthService, logger util.Logger) *Handler {
	return &Handler{
		AuthService: authService,
		Logger:      logger,
	}
}

// GetUser handles GET /user?id=123
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Missing user ID")
		return
	}

	user, err := h.AuthService.GetUserByID(userID)
	if err != nil {
		util.RespondWithError(w, err.GetStatusCode(), err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, user)
}

// Authenticate handles POST /authenticate
func (h *Handler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	token, err := h.AuthService.Authenticate(creds.Username, creds.Password)
	if err != nil {
		util.RespondWithError(w, err.GetStatusCode(), err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}
