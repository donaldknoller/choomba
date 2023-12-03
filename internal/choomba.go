package internal

type Choomba struct {
	URL       string `db:"url" json:"url"`
	MediaType string `db:"mediatype" json:"mediaType"`
	Tags      string `db:"tags" json:"tags"`
	AuxLinks  string `db:"auxlinks" json:"auxLinks"`
	Summary   string `db:"summary" json:"summary"`
	SavedAt   int64  `db:"savedAt" json:"savedAt"`
}
