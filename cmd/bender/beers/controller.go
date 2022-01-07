package beers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sackslab/bender/internal/i4217"
	"github.com/sackslab/bender/internal/middlewares/apperror"
	"github.com/sackslab/bender/internal/validators"
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

		_, err := ctrl.svc.CreateBeer(c, form)
		if err != nil {
			c.Error(err)
			return
		}

		c.Status(http.StatusCreated)
	}
}

func (ctrl *controller) ListBeers() gin.HandlerFunc {
	return func(c *gin.Context) {
		beers, err := ctrl.svc.ListBeers(c)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, beers)
	}
}

func (ctrl *controller) GetBeer() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := getNumericParam(c, "id")
		if err != nil {
			c.Error(err)
			return
		}

		beer, err := ctrl.svc.GetBeer(c, int64(id))
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, beer)
	}
}

func (ctrl *controller) GetBoxPrice() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := getNumericParam(c, "id")
		if err != nil {
			c.Error(err)
			return
		}

		qtyStr := c.DefaultQuery("quantity", "6")
		qty, err := strconv.Atoi(qtyStr)
		if err != nil {
			c.Error(&apperror.AppError{
				Code:    http.StatusBadRequest,
				Message: "query param 'quantity' must be a numeric string",
			})
			return
		}

		currencyStr := c.Query("currency")
		currency, found := i4217.ByName(currencyStr)
		if !found {
			c.Error(&apperror.AppError{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("invalid currency code or type %s, expected ISO4217 values", currencyStr),
			})
			return
		}

		price, err := ctrl.svc.GetBoxPrice(c, int64(id), BoxPriceQuery{
			Currency: validators.I4217(currency),
			Quantity: int64(qty),
		})
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, price)
	}
}

func (ctrl *controller) ConfigureRouter(r *gin.Engine) {
	ctrl.logger.With(zap.Any("service", "BeersController")).Info("Setting up beers routes")
	rg := r.Group("beers")

	rg.POST("/", ctrl.CreateBeer())
	rg.GET("/", ctrl.ListBeers())
	rg.GET("/:id", ctrl.GetBeer())
	rg.GET("/:id/boxprice", ctrl.GetBoxPrice())
}

func getNumericParam(c *gin.Context, param string) (int, error) {
	strId := c.Param(param)
	id, err := strconv.Atoi(strId)
	if err != nil {
		return -1, &apperror.AppError{
			Code:    http.StatusBadRequest,
			Message: "param 'id' must be a numeric string",
		}
	}

	return id, nil
}
