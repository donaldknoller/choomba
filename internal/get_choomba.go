package internal

import (
	"fmt"
)

// "http://example.com"
func (s *ServerContext) GetChoomba(id string) (*Choomba, error) {
	var c Choomba

	rows, err := s.db.Queryx("SELECT * FROM choomba WHERE url = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		scanErr := rows.StructScan(&c)
		if scanErr != nil {
			return nil, scanErr
		}
		fmt.Printf("%#v\n", c)
	}
	return &c, nil

}
