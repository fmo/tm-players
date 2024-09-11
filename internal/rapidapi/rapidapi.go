package rapidapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fmo/tm-players/internal/application/core/domain"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Data struct {
	Players []domain.Player `json:"data"`
}

type PlayersApi struct {
	logger *logrus.Logger
}

func NewPlayersApi(l *logrus.Logger) PlayersApi {
	return PlayersApi{
		logger: l,
	}
}

func (p PlayersApi) GetPlayers(season, teamId int) []domain.Player {
	url := fmt.Sprintf("https://transfermarkt-db.p.rapidapi.com/v1/clubs/squad?season_id=%d&locale=UK&club_id=%d",
		season,
		teamId,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Add("X-RapidAPI-Key", os.Getenv("RAPID_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var playerResponse Data
	if err := json.Unmarshal(response, &playerResponse); err != nil {
		log.Fatalf("error unmarshalling json: %v\n", err)
	}

	playerNames := make([]string, 0, 3)
	for i, p := range playerResponse.Players {
		if i >= 3 {
			break
		}
		playerNames = append(playerNames, p.Name)
	}

	p.logger.WithFields(logrus.Fields{
		"firstThreeNames": playerNames,
	}).Info("Rapid api response summary with player names")

	return playerResponse.Players
}
