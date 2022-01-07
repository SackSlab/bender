package beers

type CreateBeer struct {
	Name     string  `json:"Name" binding:"required"`
	Brewery  string  `json:"Brewery" binding:"required"`
	Country  string  `json:"Country" binding:"required"`
	Price    float64 `json:"Price" binding:"required,min=1"`
	Currency string  `json:"Currency" binding:"required"`
}
