package beers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type controller struct {
	logger *zap.Logger
	svc    *service
}

func NewController(logger *zap.Logger, svc *service) *controller {
	ctrl := new(controller)
	ctrl.logger = logger
	ctrl.svc = svc

	return ctrl
}

func (ctrl *controller) CreateBeer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var form CreateBeer
		if err := c.ShouldBindJSON(&form); err != nil {
			c.Error(err)
			return
		}
	}
}

func (ctrl *controller) ListBeers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (ctrl *controller) GetBeer() gin.HandlerFunc {
	return func(*gin.Context) {}
}
func (ctrl *controller) GetBoxPrice() gin.HandlerFunc {
	return func(*gin.Context) {}
}

func (ctrl *controller) ConfigureRouter(r *gin.Engine) {
	ctrl.logger.With(zap.Any("service", "BeersController")).Info("Setting up beers routes")
	rg := r.Group("beers")

	rg.POST("/", ctrl.CreateBeer())
	rg.GET("/", ctrl.ListBeers())
	rg.GET("/:id", ctrl.GetBeer())
	rg.GET("/:id/boxprice", ctrl.GetBoxPrice())
}
