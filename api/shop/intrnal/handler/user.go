package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type User struct {
	ID          json.RawMessage `json:"id"`
	UserName    string          `json:"username"`
	Email       string          `json:"email"`
	Role        string          `json:"role"`
	Status      string          `json:"status"`
	LastLoginAt time.Time       `json:"last_login_at"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

func GetMeMock(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:          json.RawMessage(`"0cefb56b-55c7-41ba-97c7-c30ee4918dc2"`),
		UserName:    "tamaco489",
		Email:       "tamaco489@gmail.com",
		Role:        "user",
		Status:      "active",
		LastLoginAt: time.Date(2025, 6, 11, 0, 0, 0, 0, time.UTC),
		CreatedAt:   time.Date(2025, 6, 11, 0, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2025, 6, 11, 0, 0, 0, 0, time.UTC),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode response", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
