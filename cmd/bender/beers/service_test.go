package beers

import (
	"context"
	"errors"
	"testing"

	"github.com/sackslab/bender/internal/currencylayer"
	"github.com/sackslab/bender/internal/validators"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// change this for your test configuration
// TODO: load from environments
const dsn = "host=127.0.0.1 user=postgres password=postgres dbname=bender port=5432 sslmode=disable"

var opts = currencylayer.Options{
	HostURL: "http://api.currencylayer.com",
	// replace with your api key
	ApiKey: "0417d6f26c8212b21109a33120b7a1e2",
}

func makeServiceMock() (*service, error) {
	db, err := gorm.Open(pg.New(pg.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	currClient := currencylayer.NewClient(opts)
	return NewService(db, currClient), nil
}

func TestCreateBeer(t *testing.T) {
	svc, err := makeServiceMock()
	if err != nil {
		t.Fatalf("cannot setup database, err: %s", err)
	}

	testCases := []struct {
		desc          string
		beer          CreateBeer
		isSuccessCase bool
	}{
		{
			desc: "Should create beer",
			beer: CreateBeer{
				Brewery:  "Palo Santo",
				Name:     "La chuchi",
				Country:  validators.I3166{Code: "PRY"},
				Price:    12000.00,
				Currency: validators.I4217{Name: "PYG"},
			},
			isSuccessCase: true,
		},
		{
			desc:          "Should be fail on create beer",
			beer:          CreateBeer{},
			isSuccessCase: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			createdBeer, err := svc.CreateBeer(context.Background(), tC.beer)
			if tC.isSuccessCase {
				if err != nil {
					t.Errorf("an unexpected error ocurred when trying to create the registry %s", err)
				}

				_, err := svc.GetBeer(context.Background(), int64(createdBeer.ID))
				if errors.Is(err, gorm.ErrRecordNotFound) {
					t.Errorf("Expected beer with id '%d'", createdBeer.ID)
				}
			}

			if !tC.isSuccessCase {
				if err == nil {
					t.Error("expected error for unsuccess case")
				}
			}
		})
	}
}

func TestGetBeer(t *testing.T) {
	svc, err := makeServiceMock()
	if err != nil {
		t.Fatalf("cannot setup database, err: %s", err)
	}

	beers, err := svc.ListBeers(context.Background())
	if err != nil {
		t.Fatalf("unexpected error encountered, err: %s", err)
	}
	if len(beers) < 1 {
		t.Fatalf("at least one record is required to run the test")
	}

	testCases := []struct {
		desc          string
		beerID        uint
		isSuccessCase bool
	}{
		{
			desc:          "Should get correctly one registry",
			beerID:        beers[0].ID,
			isSuccessCase: true,
		},
		{
			desc:          "Should returns gorm.ErrRecordNotFound error",
			beerID:        123456789978431,
			isSuccessCase: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := svc.GetBeer(context.Background(), int64(tC.beerID))
			if err != nil && tC.isSuccessCase {
				t.Errorf("an unexpected error ocurred when trying to fetch the registry with id '%d', err %s", tC.beerID, err)
			}

			if !errors.Is(err, gorm.ErrRecordNotFound) && !tC.isSuccessCase {
				t.Errorf("an unexpected error ocurred when trying to fetch the registry with id '%d', err %s", tC.beerID, err)
			}
		})
	}
}
