package meta

// Data represents the metadata of the game.
type Data struct {
	Name    string   `json:"name"`
	Version string   `json:"version"`
	Authors []string `json:"authors"`
	Link    string   `json:"link"`

	Description string   `json:"description"`
	License     string   `json:"license"`
	Tags        []string `json:"tags"`

	Icon        string   `json:"icon"`
	Screenshots []string `json:"screenshots"`

	ReleaseDate string `json:"release_date"`

	SupportEmail string `json:"support_email"`
	Website      string `json:"website"`
}
