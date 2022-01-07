package currencylayer

type RatesResponse struct {
	Success bool               `json:"success"`
	Source  string             `json:"source"`
	Quotes  map[string]float64 `json:"quotes"`
}
