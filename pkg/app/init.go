package app

import (
	"handsongo/pkg/router"
	"handsongo/pkg/server"
	"handsongo/pkg/config"
	"net/http"
)


func initHTTPServer(configFile string) {
	cfg := config.NewConfig(configFile)
	rt := initRouter(cfg)

	server.NewServer(cfg, rt).Start()
}



func initRouter(cfg config.Config) http.Handler {

	return router.Router()
}
