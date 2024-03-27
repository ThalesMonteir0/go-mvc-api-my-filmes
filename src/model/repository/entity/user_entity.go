package entity

type UserEntity struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
