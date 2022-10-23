package handler

import (
	"fmt"
	"net/http"
	"testing"
	"wagerservice/cmd/entity-server/db/model"
	"wagerservice/cmd/entity-server/registry"
	"wagerservice/cmd/entity-server/service/wager"
	"wagerservice/config"
	"wagerservice/internal/pkg/logger"
	"wagerservice/test/testdb"
	"wagerservice/test/testgin"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type wagerHandlerTestSuite struct {
	suite.Suite
	wagerHandlerFactory func() (*WagerHandler, func())
}

func (suite *wagerHandlerTestSuite) SetupTest() {
	suite.wagerHandlerFactory = func() (*WagerHandler, func()) {
		db, release := testdb.GetDB()
		// Init config
		config := config.InitFromFile("./config/config.toml")
		// Init logger
		logger, _ := logger.Init(config)
		// Init Provider
		provider := &registry.Provider{
			DB:     db,
			Config: config,
			Logger: logger,
		}
		// Init service
		wagerService := wager.NewWagerService(provider.DB)
		// Init Handler
		return NewWagerHandler(
			wagerService,
			*provider,
		), release
	}
}

func (suite *wagerHandlerTestSuite) TestWagerHandler_List() {
	t := suite.T()
	wagerHdl, release := suite.wagerHandlerFactory()
	defer release()

	t.Run("get list wager success", func(t *testing.T) {
		ginCtx, _, recorder := testgin.GetTestContext()
		ginCtx.Request = testgin.MustMakeRequest("GET", "/wager?page=1&limit=3&order_by=ID&order=ASC", nil)
		wagerHdl.List(ginCtx)
		result := recorder.Result()
		defer result.Body.Close()

		if result.StatusCode != http.StatusOK {
			t.Errorf("want http status: %d, got: %d", http.StatusOK, result.StatusCode)
		}
		body := testgin.ExtractBody(result.Body)

		var total int64
		wagerHdl.provider.DB.Model(&model.Wager{}).Count(&total)

		assert.JSONEq(
			t,
			testgin.JSONStr(`
			{
				"data":{
				   "data":[
					  {
						 "id":1,
						 "total_wager_value":1500,
						 "odds":10,
						 "selling_percentage":80,
						 "selling_price":1200.55,
						 "current_selling_price":1200.55,
						 "percentage_sold":null,
						 "amount_sold":null,
						 "placed_at":"2022-10-22T11:30:41.685+07:00"
					  },
					  {
						 "id":2,
						 "total_wager_value":2500,
						 "odds":20,
						 "selling_percentage":90,
						 "selling_price":2300.7,
						 "current_selling_price":2300.7,
						 "percentage_sold":null,
						 "amount_sold":null,
						 "placed_at":"2022-10-22T11:31:42.435+07:00"
					  },
					  {
						 "id":3,
						 "total_wager_value":1250,
						 "odds":15,
						 "selling_percentage":60,
						 "selling_price":800.05,
						 "current_selling_price":800.05,
						 "percentage_sold":null,
						 "amount_sold":null,
						 "placed_at":"2022-10-22T11:33:07.287+07:00"
					  }
				   ],
				   "total": `+fmt.Sprintf("%d", total)+`
				},
				"error":null,
				"success":true
			 }
			`),
			body,
		)
	})
}

func (suite *wagerHandlerTestSuite) TestWagerHandler_Create() {
	t := suite.T()
	wagerHdl, release := suite.wagerHandlerFactory()
	defer release()

	t.Run("validation failure: when missing param", func(t *testing.T) {
		ginCtx, _, recorder := testgin.GetTestContext()
		ginCtx.Request = testgin.MustMakeRequest("POST", "/wagers", map[string]interface{}{
			// "total_wager_value":  1150,
			"odds":               10,
			"selling_percentage": 70,
			"selling_price":      810.3,
		})
		wagerHdl.Create(ginCtx)
		result := recorder.Result()
		defer result.Body.Close()

		if result.StatusCode != http.StatusBadRequest {
			t.Errorf("want http status: %d, got: %d", http.StatusBadRequest, result.StatusCode)
		}
	})

	t.Run("validation failure: invalid selling_price param", func(t *testing.T) {
		ginCtx, _, recorder := testgin.GetTestContext()
		ginCtx.Request = testgin.MustMakeRequest("POST", "/wagers", map[string]interface{}{
			"total_wager_value":  1150,
			"odds":               10,
			"selling_percentage": 70,
			"selling_price":      500.3, // must be greater than total_wager_value * (selling_percentage / 100)
		})
		wagerHdl.Create(ginCtx)
		result := recorder.Result()
		defer result.Body.Close()

		if result.StatusCode != http.StatusBadRequest {
			t.Errorf("want http status: %d, got: %d", http.StatusBadRequest, result.StatusCode)
		}
	})

	t.Run("create sucess", func(t *testing.T) {
		ginCtx, _, recorder := testgin.GetTestContext()
		wagerReq := map[string]interface{}{
			"total_wager_value":  1150,
			"odds":               10,
			"selling_percentage": 70,
			"selling_price":      850.3,
		}
		ginCtx.Request = testgin.MustMakeRequest("POST", "/wagers", wagerReq)

		wagerHdl.Create(ginCtx)
		result := recorder.Result()
		defer result.Body.Close()

		if result.StatusCode != http.StatusCreated {
			t.Errorf("want http status: %d, got: %d", http.StatusCreated, result.StatusCode)
		}
	})
}

func (suite *wagerHandlerTestSuite) TestWagerHandler_Buy() {
	t := suite.T()
	wagerHdl, release := suite.wagerHandlerFactory()
	defer release()

	t.Run("validation failure: when missing param", func(t *testing.T) {
		ginCtx, _, recorder := testgin.GetTestContext()
		ginCtx.Request = testgin.MustMakeRequest("POST", "/wagers/buy/4", nil)
		ginCtx.Params = []gin.Param{{Key: "wager_id", Value: "4"}}

		wagerHdl.Buy(ginCtx)
		result := recorder.Result()
		defer result.Body.Close()

		if result.StatusCode != http.StatusBadRequest {
			t.Errorf("want http status: %d, got: %d", http.StatusBadRequest, result.StatusCode)
		}
	})

	t.Run("validation failure: invalid wagerID", func(t *testing.T) {
		ginCtx, _, recorder := testgin.GetTestContext()
		ginCtx.Request = testgin.MustMakeRequest("POST", "/wagers/buy/100", map[string]interface{}{
			"buying_price": 20.5,
		})
		ginCtx.Params = []gin.Param{{Key: "wager_id", Value: "100"}}
		wagerHdl.Buy(ginCtx)
		result := recorder.Result()
		defer result.Body.Close()

		if result.StatusCode != http.StatusBadRequest {
			t.Errorf("want http status: %d, got: %d", http.StatusBadRequest, result.StatusCode)
		}
	})

	t.Run("validation failure: invalid buying_price", func(t *testing.T) {
		ginCtx, _, recorder := testgin.GetTestContext()
		ginCtx.Request = testgin.MustMakeRequest("POST", "/wagers/buy/4", map[string]interface{}{
			"buying_price": 900.5, //must be lesser or equal to current_selling_price of the wager_id
		})
		ginCtx.Params = []gin.Param{{Key: "wager_id", Value: "4"}}
		wagerHdl.Buy(ginCtx)
		result := recorder.Result()
		defer result.Body.Close()

		if result.StatusCode != http.StatusBadRequest {
			t.Errorf("want http status: %d, got: %d", http.StatusBadRequest, result.StatusCode)
		}
	})

	t.Run("buy wager success", func(t *testing.T) {
		ginCtx, _, recorder := testgin.GetTestContext()
		buyingPrice := 800.15
		ginCtx.Request = testgin.MustMakeRequest("POST", "/wagers/buy/4", map[string]interface{}{
			"buying_price": buyingPrice,
		})
		ginCtx.Params = []gin.Param{{Key: "wager_id", Value: "4"}}
		wagerHdl.Buy(ginCtx)
		result := recorder.Result()
		defer result.Body.Close()

		if result.StatusCode != http.StatusCreated {
			t.Errorf("want http status: %d, got: %d", http.StatusCreated, result.StatusCode)
		}
		// Check in the DB
		modelWager := &model.Wager{Model: gorm.Model{ID: 4}}
		if err := modelWager.Get(wagerHdl.provider.DB); err != nil {
			t.Errorf("want getting wager object but get error")
		}
		if modelWager.CurrentSellingPrice != float32(buyingPrice) {
			t.Errorf("want CurrentSellingPrice: %v, got: %v", buyingPrice, modelWager.CurrentSellingPrice)
		}
		// TODO (phuong) : check more fields here
	})
}

func TestWagerHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(wagerHandlerTestSuite))
}
