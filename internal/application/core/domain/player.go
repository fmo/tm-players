package domain

type Player struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	TeamId              int    `json:"teamId"`
	Age                 int    `json:"age"`
	Position            string `json:"position"`
	MarketValue         int    `json:"marketValue"`
	MarketValueCurrency string `json:"marketValueCurrency"`
}
