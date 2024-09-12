package domain

type Player struct {
	Height        string      `json:"height"`
	Joined        interface{} `json:"joined"`
	ContractUntil interface{} `json:"contractUntil"`
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	Image         string      `json:"image"`
	ShirtNumber   string      `json:"shirtNumber"`
	Age           int         `json:"age"`
	TeamId        int         `json:"teamId"`
}
