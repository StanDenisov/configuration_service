package router

import (
	"bitbucket.org/warlorddevteam/conf/controller"
	"github.com/labstack/echo"
)

//Router - default logic router
func Router(e *echo.Echo) {
	e.GET("/:service_name", controller.GetConf)
}
