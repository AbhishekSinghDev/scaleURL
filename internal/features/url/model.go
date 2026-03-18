package url

type Url struct {
	Id int64 `json:"id"`
	OriginalUrl string `json:"original_url"`
	ShortCode string `json:"short_code"`
	ExpiresAt string `json:"expires_at"`
	CreatedAt string `json:"created_at"`
}

type CreateURLParams struct {
	OriginalUrl string `json:"original_url"`
}
