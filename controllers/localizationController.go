package controllers

import (
	"api-localization/logger"
	"api-localization/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type LocalizationController struct {
}

func (controller *LocalizationController) CalculateRoute(c echo.Context) error {

	latitude1 := c.QueryParam("lat1")
	longitude1 := c.QueryParam("long1")
	latitude2 := c.QueryParam("lat2")
	longitude2 := c.QueryParam("long2")
	url := fmt.Sprintf("https://api.tomtom.com/routing/1/calculateRoute/%s%s%s%s/json?computeBestOrder=true&sectionType=traffic&routeType=fastest&avoid=unpavedRoads&travelMode=car&key=mPAGwhEHVNv5yWJlTHCsbHceQm4pFTY2", latitude1+"%2C", longitude1+"%3A", latitude2+"%2C", longitude2)
	var temporalStorage map[string]interface{}
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
	fmt.Println(temporalStorage["routes"])
	return c.JSON(http.StatusOK, temporalStorage["routes"])

}

func (controller *LocalizationController) ShowCoordinatesforDangerZones(c echo.Context) error {
	storage := make(map[int][]float32)
	db := repositories.ReturnDangerLocations()
	rows, err := db.Query("SELECT * FROM locations")
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	for rows.Next() {
		var id int
		var lat float32
		var long float32
		err = rows.Scan(&id, &lat, &long)
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		storage[id] = []float32{lat, long}

	}
	return c.JSON(http.StatusOK, storage)
}

func (controller *LocalizationController) CalculateDistanceToFences(c echo.Context) error {
	point := c.QueryParam("point")
	objectID := c.QueryParam("object")
	apiKey := "mPAGwhEHVNv5yWJlTHCsbHceQm4pFTY2"
	projectID := "ea8ca9cc-367c-4208-893e-d6485a076e27"
	url := fmt.Sprintf("https://api.tomtom.com/geofencing/1/report/%s?key=%s&point=%s&object=%s", projectID, apiKey, point, objectID)
	var temporalStorage map[string]interface{}
	Storage := make(map[string]interface{})
	client := http.Client{Timeout: time.Second * 30}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error("LocalizationController", "CalculateDistanceToFences", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())

	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("LocalizationController", "CalculateDistanceToFences", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = json.NewDecoder(resp.Body).Decode(&temporalStorage)
	if err != nil {
		logger.Error("LocalizationController", "CalculateDistanceToFences", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	Storage["inside"] = temporalStorage["inside"]
	Storage["outside"] = temporalStorage["outside"]

	return c.JSON(http.StatusOK, Storage)

}
