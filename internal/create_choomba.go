package internal

func (s *ServerContext) CreateChoomba(url string) error {
	choomba := Choomba{
		URL:       url,
		MediaType: "image/jpeg",
		Tags:      "cyberpunk,art",
		AuxLinks:  "http://link1.com,http://link2.com",
		Summary:   "A depiction of a cyberpunk cityscape",
		SavedAt:   1672525200, // Example UNIX timestamp
	}

	tx := s.db.MustBegin()
	_, err := tx.NamedExec(
		`INSERT INTO choomba (url, mediatype, tags, auxlinks, summary, savedAt) VALUES (:url, :mediatype, :tags, :auxlinks, :summary, :savedAt)`,
		&choomba,
	)
	if err != nil {
		return err
	}
	return tx.Commit()
}
