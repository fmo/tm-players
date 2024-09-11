package ports

import (
	"context"
	"github.com/fmo/tm-players/internal/application/core/domain"
)

type DBPort interface {
	Save(ctx context.Context, player *domain.Player) error
}
