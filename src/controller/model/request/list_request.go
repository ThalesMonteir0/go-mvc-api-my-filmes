package request

type ListRequest struct {
	Title       string `json:"name"`
	Description string `json:"description"`
	UserID      int    `json:"user_id"`
}
