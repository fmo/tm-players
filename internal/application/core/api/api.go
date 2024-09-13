package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fmo/tm-players/internal/application/core/domain"
	"github.com/fmo/tm-players/internal/ports"
	log "github.com/sirupsen/logrus"
	"time"
)

type Application struct {
	playerData ports.PlayerData
	database   ports.DBPort
	cache      ports.CachePort
}

func NewApplication(playerData ports.PlayerData, database ports.DBPort, cache ports.CachePort) *Application {
	return &Application{
		playerData: playerData,
		database:   database,
		cache:      cache,
	}
}

func (a Application) SavePlayer(ctx context.Context, season, teamId int) error {
	cacheKey := fmt.Sprintf("tm:players:squad:%d", teamId)
	squadData, err := a.cache.Get(ctx, cacheKey)
	var players []domain.Player

	if err == nil {
		err = json.Unmarshal([]byte(squadData), &players)
		if err == nil {
			log.Info("Found squad for team %s in Redis returning response", teamId)
		}
	} else {
		log.Info("Getting player data from rapidapi")
		players = a.playerData.GetPlayers(ctx, season, teamId)
		jsonSquad, err := json.Marshal(players)
		if err == nil {
			a.cache.Set(ctx, cacheKey, jsonSquad, 240*time.Hour)
		}
	}

	for _, player := range players {
		err := a.database.Save(ctx, &player)
		if err != nil {
			return err
		}
	}

	return nil
}
