package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/tamaco489/supabase_sample/api/shop/intrnal/repository"
)

type User struct {
	ID          string    `json:"id"`
	UserName    string    `json:"username"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	Status      string    `json:"status"`
	LastLoginAt time.Time `json:"last_login_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetMe(w http.ResponseWriter, r *http.Request) {

	db := repository.GetPool()

	// NOTE: 本来はJWT等から取得するべき
	uid := "0cefb56b-55c7-41ba-97c7-c30ee4918dc2"

	rows, err := db.Query(r.Context(), "SELECT * FROM users WHERE id = $1", uid)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to query users", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	user := User{}
	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.UserName,
			&user.Email,
			&user.Role,
			&user.Status,
			&user.LastLoginAt,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			slog.ErrorContext(r.Context(), "failed to scan user", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode response", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
