package ports

import (
	"context"
	"github.com/fmo/tm-players/internal/application/core/domain"
)

type PlayerData interface {
	GetPlayers(ctx context.Context, season, teamId int) []domain.Player
}
