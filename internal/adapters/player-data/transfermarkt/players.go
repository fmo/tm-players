package transfermarkt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fmo/tm-players/internal/application/core/domain"
	"github.com/fmo/tm-players/internal/ports"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type Adapter struct {
	rapidApiKey string
	cache       ports.CachePort
}

func NewAdapter(rapidApiKey string, cache ports.CachePort) (*Adapter, error) {
	return &Adapter{
		rapidApiKey: rapidApiKey,
		cache:       cache,
	}, nil
}

type Player struct {
	Height        string        `json:"height"`
	Foot          string        `json:"foot"`
	Injury        interface{}   `json:"injury"`
	Suspension    interface{}   `json:"suspension"`
	Joined        interface{}   `json:"joined"`
	ContractUntil interface{}   `json:"contractUntil"`
	Captain       bool          `json:"captain"`
	LastClub      interface{}   `json:"lastClub"`
	IsLoan        interface{}   `json:"isLoan"`
	WasLoan       interface{}   `json:"wasLoan"`
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	Image         string        `json:"image"`
	ImageLarge    interface{}   `json:"imageLarge"`
	ImageSource   string        `json:"imageSource"`
	ShirtNumber   string        `json:"shirtNumber"`
	Age           int           `json:"age"`
	DateOfBirth   int64         `json:"dateOfBirth"`
	HeroImage     string        `json:"heroImage"`
	IsGoalkeeper  bool          `json:"isGoalkeeper"`
	Positions     Positions     `json:"positions"`
	Nationalities []Nationality `json:"nationalities"`
	MarketValue   MarketValue   `json:"marketValue"`
	TeamId        int           `json:"teamId"`
}

type Positions struct {
	First  Position  `json:"first"`
	Second *Position `json:"second"`
	Third  *Position `json:"third"`
}

type Position struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	Group     string `json:"group"`
}

type Nationality struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type MarketValue struct {
	Value       int         `json:"value"`
	Currency    string      `json:"currency"`
	Progression interface{} `json:"progression"`
}

type Data struct {
	Players []Player `json:"data"`
}

func (a Adapter) GetPlayers(ctx context.Context, season, teamId int) []domain.Player {
	cacheKey := fmt.Sprintf("tm:players:squad:%d", teamId)
	squadData, err := a.cache.Get(ctx, cacheKey)

	var players []Player

	if err == nil && squadData != "" {
		err = json.Unmarshal([]byte(squadData), &players)
		if err == nil {
			log.Info("Found squad for team %s in Redis returning response", teamId)
		} else {
			log.Error(err)
		}
	} else {
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

		req.Header.Add("X-RapidAPI-Key", a.rapidApiKey)

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

		players = playerResponse.Players

		jsonSquad, err := json.Marshal(players)
		if err == nil {
			a.cache.Set(ctx, cacheKey, jsonSquad, 240*time.Hour)
		}
	}

	var domainPlayer []domain.Player

	for _, p := range players {
		p := domain.Player{
			TeamId:   teamId,
			Name:     p.Name,
			ID:       p.ID,
			Age:      p.Age,
			Position: p.Positions.First.Name,
		}
		domainPlayer = append(domainPlayer, p)
	}

	playerNames := make([]string, 0, 3)
	for i, p := range players {
		if i >= 3 {
			break
		}
		playerNames = append(playerNames, p.Name)
	}

	log.WithFields(logrus.Fields{
		"firstThreeNames": playerNames,
	}).Info("Rapid api response summary with player names")

	return domainPlayer
}
