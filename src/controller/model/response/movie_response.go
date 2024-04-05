package response

type MovieResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	UrlImg      string `json:"url_img"`
}
