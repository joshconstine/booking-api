package response

type DocumentResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
