package typeform

type User struct {
	Alias    string `json:"alias"`
	Email    string `json:"email"`
	Language string `json:"language"`
	UserID   string `json:"user_id"`
}
