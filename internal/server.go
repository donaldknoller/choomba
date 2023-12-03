package internal

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	_ "modernc.org/sqlite"
)

func Ayo() {
	fmt.Println("ayo")
}

type ServerContext struct {
	e  *echo.Echo
	db *sqlx.DB
}

func NewServer() (*ServerContext, error) {

	db, err := sqlx.Open("sqlite", ConnectionName())
	if err != nil {
		return nil, err
	}
	if dbErr := db.Ping(); dbErr != nil {
		return nil, dbErr
	}

	for _, migration := range Migrations() {
		_, err = db.ExecContext(
			context.Background(),
			migration,
		)
		if err != nil {
			return nil, err
		}
	}

	log.Debug().Msg("ayylmao")
	return &ServerContext{
		e:  nil,
		db: db,
	}, nil
}
