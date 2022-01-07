package beers

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/sackslab/bender/cmd/bender/currencyrates"
	"github.com/sackslab/bender/internal/currencylayer"
	"gorm.io/gorm"
)

type service struct {
	db         *gorm.DB
	currLayerC *currencylayer.Client
}

func NewService(db *gorm.DB, currLayerC *currencylayer.Client) *service {
	svc := &service{db: db}
	db.AutoMigrate(&Beer{})
	return svc
}

func (svc *service) CreateBeer(ctx context.Context, in CreateBeer) (*Beer, error) {
	var beer Beer
	if err := copier.Copy(&beer, &in); err != nil {
		return nil, err
	}
	beer.Country = &in.Country.Code
	beer.Currency = &in.Currency.Name

	if result := svc.db.WithContext(ctx).Create(&beer); result.Error != nil {
		return nil, result.Error
	}

	return &beer, nil
}

// TODO: add pagination support
func (svc *service) ListBeers(ctx context.Context) ([]Beer, error) {
	var beers []Beer

	if result := svc.db.Find(&beers); result.Error != nil {
		return nil, result.Error
	}

	return beers, nil
}

func (svc *service) GetBeer(ctx context.Context, pk int64) (*Beer, error) {
	var beer Beer
	if result := svc.db.WithContext(ctx).First(&beer, pk); result.Error != nil {
		return nil, result.Error
	}

	return &beer, nil
}

func (svc *service) GetBoxPrice(ctx context.Context, pk int64, opts BoxPriceQuery) (*BoxPrice, error) {
	beer, err := svc.GetBeer(ctx, pk)
	if err != nil {
		return nil, err
	}

	totalAmount := float64(opts.Quantity) * *beer.Price
	if beer.Currency != &opts.Currency.Name {
		rates, err := svc.currLayerC.Latest()
		if err != nil {
			return nil, err
		}

		totalAmount, err = currencyrates.GetConversion(rates, totalAmount, *beer.Currency, opts.Currency.Name)
		if err != nil {
			return nil, err
		}
	}

	return &BoxPrice{
		TotalAmount: totalAmount,
		Currency:    opts.Currency.Name,
	}, nil
}
