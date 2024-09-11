package ports

import "context"

type APIPorts interface {
	SavePlayer(ctx context.Context, season, teamId int) error
}
