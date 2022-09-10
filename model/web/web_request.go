package web

type WebRequest struct {
	OriginalUrl string `json:"original_url" validate:"url"`
}
