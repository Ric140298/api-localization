package controllers

import (
	"api-localization/logger"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type LocalizationController struct {
}

func (controller *LocalizationController) CalculateRoute(c echo.Context) error {
	latitude1 := c.QueryParam("lat1")
	longitude1 := c.QueryParam("long")
	latitude2 := c.QueryParam("lat2")
	longitude2 := c.QueryParam("long2")
	url := fmt.Sprintf("https://api.tomtom.com/routing/1/calculateRoute/%s%s%s%s/json?computeBestOrder=true&sectionType=traffic&routeType=fastest&avoid=unpavedRoads&travelMode=car&key=mPAGwhEHVNv5yWJlTHCsbHceQm4pFTY2", latitude1+"%2C", longitude1+"%3A", latitude2+"%2C", longitude2)
	var temporalStorage interface{}
	client := http.Client{Timeout: time.Second * 30}
	req, err := http.NewRequest("GET", url, nil)
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
