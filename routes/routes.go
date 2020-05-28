package routes

import (
	"Localizacion/controllers"
	"Localizacion/di"

	"github.com/labstack/echo"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc echo.HandlerFunc
}

// ServerRoutes :
type ServerRoutes []Route

var Routes ServerRoutes

const serviceName = "/localization"

func init() {
	controllersProvider := di.BuildContainer()
	controllersProvider.Invoke(func(localizationController *controllers.LocalizationController) {
		Routes = ServerRoutes{
			Route{Method: "GET", Name: "GetRoute", Pattern: serviceName + "/calcule", HandlerFunc: localizationController.CalculateRoute},
		}
	})

}
