package types
import "time"

type Account struct {
	ID       int   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	UserID   int   `json:"user_id"`
}

type AccountSecurity struct {
	ID                      *int      `json:"id"`
	UserID                  int       `json:"user_id"`
	Attempts                int       `json:"attempts"`
	LastAttempt             time.Time `json:"last_attempt"`
	LastTimePasswordChanged time.Time `json:"last_time_password_changed"`
}