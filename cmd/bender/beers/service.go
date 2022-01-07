package beers

type service struct{}

func NewService() *service {
	return &service{}
}

func (svc *service) CreateBeer() {
}

func (svc *service) ListBeers() {
}

func (svc *service) GetBeer() {
}

func (svc *service) GetBoxPrice() {
}
