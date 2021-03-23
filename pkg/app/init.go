package app

import (
	"handsongo/pkg/config"
	"handsongo/pkg/db"
	"handsongo/pkg/reporters"
	"handsongo/pkg/router"
	"handsongo/pkg/server"
	"handsongo/pkg/user"
	"io"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"
)

func initHTTPServer(configFile string) {
	cfg := config.NewConfig(configFile)
	lgr := initLogger(cfg)

	rt := initRouter(cfg, lgr)

	server.NewServer(cfg, lgr, rt).Start()
}

func initRouter(cfg config.Config, lgr *zap.Logger) http.Handler {
	userRepo := initRepository(cfg)
	userService := initService(userRepo,cfg)
	return router.NewRouter(lgr,userService)
}

func initService( userRepo user.UserRepository, cfg config.Config) (user.Service) {

	userService := user.NewUserService(userRepo)

	return userService
}


func initRepository(cfg config.Config) (user.UserRepository) {
	dbConfig := cfg.GetDBConfig()
	dbHandler := db.NewDBHandler(dbConfig)

	db, err := dbHandler.GetDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	return user.NewUserRepository(db)
}

func initLogger(cfg config.Config) *zap.Logger {
	return reporters.NewLogger(
		cfg.GetLogConfig().GetLevel(),
		getWriters(cfg.GetLogFileConfig())...,
	)
}

func getWriters(cfg config.LogFileConfig) []io.Writer {
	return []io.Writer{
		os.Stdout,
		reporters.NewExternalLogFile(cfg),
	}
}
