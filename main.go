package main

import (
	"database/sql"

	"github.com/chonlatee11/gin-backend/api"
	db "github.com/chonlatee11/gin-backend/db/sqlc"
	"github.com/chonlatee11/gin-backend/util"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSOURCE)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	log.Info().Msgf("loaded config: %#v", config)

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
