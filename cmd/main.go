package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sajjanjyothi/volume_assignment/api"
	"github.com/sajjanjyothi/volume_assignment/internal/flightpath"
	"github.com/sajjanjyothi/volume_assignment/internal/service"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	zap.ReplaceGlobals(logger)

	//New echo server instance
	e := echo.New()

	//New flight path detector
	airportDetector := flightpath.NewAirportDetector()

	//New API service implementation
	s := service.NewService(airportDetector)

	//register handlers with echo
	api.RegisterHandlers(e, s)

	//API viewer
	e.File("/", "./api/api_viewer.html")
	e.File("airportcalculator.yaml", "./api/airportcalculator.yaml")
	logger.Info("Webserver starting")
	logger.Fatal("echo server start failed", zap.Any("echo-server-error", e.Start(":8080")))
}
