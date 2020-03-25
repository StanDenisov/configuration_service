package router

import (
	"github.com/WarlordDev/configuration_service/controller"

	"github.com/labstack/echo"
)

//Router - default logic router
func Router(e *echo.Echo) {
	e.GET("/:service_name", controller.GetConf)
}
