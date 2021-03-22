package app

import (
	"handsongo/pkg/config"
	"handsongo/pkg/reporters"
	"handsongo/pkg/router"
	"handsongo/pkg/server"
	"io"
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

	return router.NewRouter(lgr)
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
