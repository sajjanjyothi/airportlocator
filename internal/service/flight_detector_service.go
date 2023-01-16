// service service defenition for this http server
package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sajjanjyothi/volume_assignment/api"
	"github.com/sajjanjyothi/volume_assignment/internal/flightpath"
	"go.uber.org/zap"
)

// flightService flihgt service implementation for API
type flightService struct {
	flightPathDetector flightpath.FlightPath
}

func NewService(filghtPathDetector flightpath.FlightPath) api.ServerInterface {
	return &flightService{
		flightPathDetector: filghtPathDetector,
	}
}

// GetSourceDestinationAirports Get source and desctination airports from itinerary
// (GET /calculate)
func (s *flightService) GetSourceDestinationAirports(ctx echo.Context) error {
	var flightData [][]string

	if err := ctx.Bind(&flightData); err != nil {
		zap.L().Error("Request parsing failed", zap.Error(err))
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	paths, err := s.flightPathDetector.GetSourceDestinationAirports(flightData)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "internal server error")
	}

	return ctx.JSON(http.StatusOK, paths)
}
