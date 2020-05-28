package main

import (
	"api-localization/logger"
	customMiddleware "api-localization/middleware/logger"
	"api-localization/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/tomtom"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	addr         = "Melbourne VIC"
	lat, lng     = 20.532714, -100.440404
	radius       = 50
	zoom         = 18
	addrFR       = "Champs de Mars Paris"
	latFR, lngFR = 48.854395, 2.304770
)

var server *echo.Echo

func init() {
	logger.NewLogger("api-localization")
	server = echo.New()
	server.HideBanner = true
}
func main() {
	for _, r := range routes.Routes {
		server.Add(r.Method, r.Pattern, r.HandlerFunc).Name = r.Name
	}
	server.Use(

		middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet},
			AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
			MaxAge:       7200,
		}),
		customMiddleware.Logger,
	)
	server.Logger.Fatal(server.Start(":" + "4004"))

}

func ExampleGeocoder() {

	fmt.Println("TomTom")
	try(tomtom.Geocoder(os.Getenv("TOMTOM_API_KEY")))
}

func try(geocoder geo.Geocoder) {
	location, _ := geocoder.Geocode(addr)
	if location != nil {
		fmt.Printf("%s location is (%.6f, %.6f)\n", addr, location.Lat, location.Lng)
	} else {
		fmt.Println("got <nil> location")
	}
	address, _ := geocoder.ReverseGeocode(lat, lng)
	if address != nil {
		fmt.Printf("Address of (%.6f,%.6f) is %s\n", lat, lng, address.FormattedAddress)
		fmt.Printf("Detailed address: %#v\n", address)
	} else {
		fmt.Println("got <nil> address")
	}
	fmt.Print("\n")

}
