package main

import (
	"github.com/donaldknoller/choomba/internal"
	"github.com/rs/zerolog/log"
	"math/rand"
	"time"
)

func main() {
	internal.InitConfig()
	s, err := internal.NewServer()
	if err != nil {
		log.Panic().Err(err).Msg("oh no")
	}
	id := randId()
	err = s.CreateChoomba(id)
	if err != nil {
		log.Panic().Err(err).Msg("oh no")
	}
	c, err := s.GetChoomba(id)
	if err != nil {
		log.Panic().Err(err).Msg("oh no")
	}
	log.Info().Interface("c", c).Msg("choomba")
}

func randId() string {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 5)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
