package domain

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
