package main

import (
	"log"
	"os"

	"github.com/fekuna/go-post-article/config"
	"github.com/fekuna/go-post-article/internal/server"
	"github.com/fekuna/go-post-article/pkg/db/mysql"
	"github.com/fekuna/go-post-article/pkg/logger"
	"github.com/fekuna/go-post-article/pkg/utils"
)

func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)

	// mysqlDB
	mysqlDB, err := mysql.NewMysqlDB(cfg)
	if err != nil {
		appLogger.Fatalf("Mysql init: %s", err)
	} else {
		appLogger.Infof("Mysql connected, Status: %#v", mysqlDB.Stats())
	}
	defer mysqlDB.Close()

	s := server.NewServer(cfg, appLogger, mysqlDB)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
