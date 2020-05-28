package controllers

import (
	"Localizacion/logger"
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type LocalizationController struct {
}

func (controller *LocalizationController) CalculateRoute(c echo.Context) error {
	var temporalStorage interface{}
	client := http.Client{Timeout: time.Second * 30}
	req, err := http.NewRequest("GET", "https://api.tomtom.com/routing/1/calculateRoute/20.532230%2C-100.441125%3A20.532140%2C-100.441865/json?computeBestOrder=true&sectionType=traffic&routeType=fastest&avoid=unpavedRoads&travelMode=car&key=mPAGwhEHVNv5yWJlTHCsbHceQm4pFTY2", nil)
	if err != nil {
		logger.Error("LocalizationController", "CalculateRoute", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())

	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("LocalizationController", "CalculateRoute", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = json.NewDecoder(resp.Body).Decode(&temporalStorage)
	if err != nil {
		logger.Error("LocalizationController", "CalculateRoute", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, temporalStorage)

}
