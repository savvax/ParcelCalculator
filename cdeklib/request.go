package cdeklib

type Request struct {
	Type         int      `json:"type"`
	Date         string   `json:"date"`
	Currency     int      `json:"currency"`
	Lang         string   `json:"lang"`
	FromLocation Location `json:"from_location"`
	ToLocation   Location `json:"to_location"`
	Packages     []Size   `json:"packages"`
}
