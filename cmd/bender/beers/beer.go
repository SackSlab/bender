package beers

import "github.com/sackslab/bender/internal/fxgorm/models"

type Beer struct {
	models.BaseModel

	Name     *string  `gorm:"not null"`
	Brewery  *string  `gorm:"not null"`
	Country  *string  `gorm:"not null"`
	Price    *float64 `gorm:"not null"`
	Currency *string  `gorm:"not null"`
}

func (beer *Beer) TableName() string {
	return "beers"
}
