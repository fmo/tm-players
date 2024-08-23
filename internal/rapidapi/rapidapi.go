package rapidapi

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
)

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

type PlayersApi struct {
	logger *logrus.Logger
}

func NewPlayersApi(l *logrus.Logger) PlayersApi {
	return PlayersApi{
		logger: l,
	}
}

func (p PlayersApi) GetPlayers(season, teamId int) []Player {
	url := fmt.Sprintf("https://transfermarkt-db.p.rapidapi.com/v1/clubs/squad?season_id=%d&locale=UK&club_id=%d",
		season,
		teamId,
	)

	response := RapidRequest(url)

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
	}).Info("rapid api response summary with player names")

	return playerResponse.Players
}
