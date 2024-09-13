package api

import (
	"context"
	"github.com/fmo/tm-players/internal/ports"
	log "github.com/sirupsen/logrus"
)

type Application struct {
	playerData ports.PlayerData
	database   ports.DBPort
}

func NewApplication(playerData ports.PlayerData, database ports.DBPort) *Application {
	return &Application{
		playerData: playerData,
		database:   database,
	}
}

func (a Application) SavePlayer(ctx context.Context, season, teamId int) error {
	log.Info("Getting player data from rapidapi")
	players := a.playerData.GetPlayers(ctx, season, teamId)

	for _, player := range players {
		err := a.database.Save(ctx, &player)
		if err != nil {
			return err
		}
	}

	return nil
}
