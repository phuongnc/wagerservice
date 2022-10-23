package router

import (
	"net/http"
	"wagerservice/cmd/entity-server/registry"
	"wagerservice/cmd/entity-server/service/wager"

	"github.com/gin-gonic/gin"

	"wagerservice/cmd/entity-server/handler"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	envDevelopment = "development"
	envProduction  = "production"
	envStaging     = "staging"
	apiHealthPath  = "/healthz"
	swaggerPath    = "/swagger/*any"
)

func Handler(
	provider *registry.Provider,
) *gin.Engine {
	// Enable Release mode for production
	if provider.Config.Env == envProduction || provider.Config.Env == envStaging {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init routes
	router := gin.New()
	initRoutes(router, provider)
	return router
}

func initRoutes(router *gin.Engine, provider *registry.Provider) {
	router.GET(apiHealthPath, healthz)

	wagerHdl := handler.NewWagerHandler(
		wager.NewWagerService(provider.DB),
		*provider,
	)
	{
		wagerRouter := router.Group("wagers")
		wagerRouter.POST("", wagerHdl.Create)
		wagerRouter.POST("buy/:wager_id", wagerHdl.Buy)
		wagerRouter.GET("", wagerHdl.List)
	}
	router.GET(swaggerPath, ginSwagger.WrapHandler(swaggerfiles.Handler))
}

// healthz for checking service status
func healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
