package beers

import "github.com/sackslab/bender/internal/validators"

type CreateBeer struct {
	Name     string           `json:"Name" binding:"required"`
	Brewery  string           `json:"Brewery" binding:"required"`
	Country  validators.I3166 `json:"Country" binding:"required"`
	Price    float64          `json:"Price" binding:"required,min=1"`
	Currency validators.I4217 `json:"Currency" binding:"required"`
}

type BoxPriceQuery struct {
	Quantity int64
	Currency validators.I4217
}

type BoxPrice struct {
	TotalAmount float64 `json:"TotalAmount"`
	Currency    string  `json:"Currency"`
}
