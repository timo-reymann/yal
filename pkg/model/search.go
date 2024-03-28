package model

// SearchEngine information
type SearchEngine struct {
	// Title to display in UI for search entry
	Title string `json:"title"`
	// UrlPrefix is the url prepended to the search term
	UrlPrefix string `json:"urlPrefix"`
}
