package di

import (
	"Localizacion/controllers"

	"go.uber.org/dig"
)

func LocalizationControllerInjector() *controllers.LocalizationController {
	return &controllers.LocalizationController{}
}

func BuildContainer() *dig.Container {
	Container := dig.New()
	Container.Provide(LocalizationControllerInjector)
	return Container
}
