package types

// News API
type NewsContent struct {
	Articles []article `json:"articles"`
}
type article struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url"`
}
