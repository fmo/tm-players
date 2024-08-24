package maps

import (
	pb "github.com/fmo/football-proto/golang/player"
	"github.com/fmo/tm-players/internal/rapidapi"
	"github.com/sirupsen/logrus"
)

type MapPlayersObj struct {
	logger *logrus.Logger
}

func NewMapPlayers(l *logrus.Logger) MapPlayersObj {
	return MapPlayersObj{
		logger: l,
	}
}

func (m MapPlayersObj) MapPlayers(players []rapidapi.Player, returnPlayer *[]*pb.Player, teamId int) {
	if len(players) == 0 {
		m.logger.Info("No players to map")
		return
	}

	m.logger.WithFields(logrus.Fields{
		"playerCountToMap": len(players),
	}).Info("Mapping starting...")

	for _, p := range players {
		player := &pb.Player{
			Name:                p.Name,
			TransfermarktId:     p.ID,
			ShirtNumber:         p.ShirtNumber,
			MarketValue:         int32(p.MarketValue.Value),
			MarketValueCurrency: p.MarketValue.Currency,
			TeamId:              int32(teamId),
			Position:            p.Positions.First.Name,
		}

		*returnPlayer = append(*returnPlayer, player)
	}
}
