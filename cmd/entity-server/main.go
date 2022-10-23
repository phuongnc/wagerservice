package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"wagerservice/cmd/entity-server/db"
	"wagerservice/cmd/entity-server/registry"
	"wagerservice/cmd/entity-server/router"
	"wagerservice/config"

	"wagerservice/internal/pkg/logger"
)

func main() {
	provider := &registry.Provider{}

	// Init configuration
	wd, _ := os.Getwd()
	configPath := filepath.Join(wd, "config/config.toml")
	provider.Config = config.InitFromFile(configPath)

	// Init logger
	logger, err := logger.Init(provider.Config)
	if err != nil {
		log.Fatal(err)
	}
	provider.Logger = logger

	// Init database connection
	DB, closeFunc, err := db.Init(provider.Config)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()
	provider.DB = DB

	// Init router
	handler := router.Handler(provider)

	// Init http server
	endPoint := fmt.Sprintf("0.0.0.0:%d", provider.Config.ServerPort)
	readTimeout := time.Minute
	writeTimeout := time.Minute
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        handler,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	provider.Logger.Info("[info] start http server listening %s", endPoint)
	_ = server.ListenAndServe()

}
