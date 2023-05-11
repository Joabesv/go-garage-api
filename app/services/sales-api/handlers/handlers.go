package handlers

import (
	"net/http"
	"os"

	"github.com/ardanlabs/service/business/web/v1/testgrp"
	"github.com/ardanlabs/service/foundation/web"
	"go.uber.org/zap"
)

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown)

	app.Handle(http.MethodGet, "/status", testgrp.Status)

	return app
}
